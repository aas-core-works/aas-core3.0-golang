package xmlization_test

import (
	"fmt"
	aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Assert that there is no de-serialization error when de-serializing
// from `source`.
func assertNoDeserializationError(
	t *testing.T,
	err error,
	source string,
) (ok bool) {
	ok = true
	if err != nil {
		ok = false
		t.Fatalf(
			"Expected no de-serialization error from %s, but got: %s",
			source, err.Error(),
		)
	}
	return
}

// Assert that there is no serialization error when serializing the instance
// originally coming from `source`.
func assertNoSerializationError(
	t *testing.T,
	err error,
	source string,
) (ok bool) {
	ok = true
	if err != nil {
		ok = false
		t.Fatalf(
			"Expected no serialization error when serializing "+
				"the instance obtained from %s, but got: %s",
			source, err.Error(),
		)
	}
	return
}

// NOTE (mristin, 2023-06-21):
// Currently, Go does not support self-closing tags,
// see: https://github.com/golang/go/issues/21399.
// We apply the following hack to make the tags self-closing even if they are not.
// This is unsafe in general, but works OK for the limited set of test data that we are
// here dealing with.
//
// The code has been taken from: https://github.com/golang/go/issues/21399#issuecomment-1342730174

func forceSelfClosingTags(text string) string {
	b := []byte(text)
	emptyTagIdxs := regexp.MustCompile(`<(\w+)></\w+>`).FindAllSubmatchIndex(b, -1)

	if len(emptyTagIdxs) == 0 {
		return text
	}

	var nb []byte

	for _, idx := range emptyTagIdxs {
		// Get everything in b up till the first of the submatch indexes (this is
		// the start of an "empty" <thing></thing> tag), then get the name of the tag
		// and put it in a self-closing tag.
		nb = append(b[0:idx[0]], fmt.Sprintf("<%s/>", b[idx[2]:idx[3]])...)

		// Finally, append everything *after* the submatch indexes
		nb = append(nb, b[len(b)-(len(b)-idx[1]):]...)
	}

	return string(nb)
}

// Assert that the serialization `other`, as XML document, equals the original
//
//	XML document `that` read from the `source`.
func assertSerializationEqualsDeserialization(
	t *testing.T,
	that string,
	other string,
	source string,
) (ok bool) {
	// Remove carriers to avoid problems between Windows, Posix and MacOS
	canonicalThat := strings.ReplaceAll(that, "\r", "")
	canonicalOther := strings.ReplaceAll(other, "\r", "")

	canonicalThat = strings.TrimSpace(canonicalThat)
	canonicalOther = strings.TrimSpace(canonicalOther)

	canonicalThat = forceSelfClosingTags(canonicalThat)
	canonicalOther = forceSelfClosingTags(canonicalOther)

	// NOTE (mristin, 2023-06-21):
	// The following hack is SUPER ugly and unsafe! However, it works. Given Go's
	// limited support for XML, we gave up on a safer approach :(. We tested
	// the following approaches before applying this hack:
	//  * A round-trip over `encoding/xml`. Failed due to
	//    https://github.com/golang/go/issues/13400.
	//  * Using `aqwari.net/xml/xmltree`. Failed as the special characters in the
	//    element content still has not been de-escaped or consistently escaped in
	//    a round trip.

	canonicalThat = strings.ReplaceAll(canonicalThat, "'", "&#39;")
	canonicalOther = strings.ReplaceAll(canonicalOther, "'", "&#39;")

	thatLines := strings.Split(canonicalThat, "\n")
	otherLines := strings.Split(canonicalOther, "\n")

	if canonicalThat != canonicalOther {
		b := new(strings.Builder)
		minLines := len(thatLines)
		if minLines > len(otherLines) {
			minLines = len(otherLines)
		}
		for i := 0; i < minLines; i++ {
			if thatLines[i] == otherLines[i] {
				b.WriteString(fmt.Sprintf("           %s\n", thatLines[i]))
			} else {
				b.WriteString(fmt.Sprintf("ORIGINAL   %s\n", thatLines[i]))
				b.WriteString(fmt.Sprintf("SERIALIZED %s\n", otherLines[i]))
				break
			}
		}

		t.Fatalf(
			"The canonicalized XML serialization of the de-serialized instance "+
				"from %s does not equal the canonicalized original XML document:\n"+
				"%s",
			source, b.String(),
		)
		ok = false
	}

	return
}

var causesForDeserializationFailure = [...]string{
	"TypeViolation",
	"RequiredViolation",
	"EnumViolation",
	"UnexpectedAdditionalProperty",
}

// Assert that there is a de-serialization error.
//
// If [aastesting.RecordMode] is set, the de-serialization error is re-recorded
// to `expectedPth`. Otherwise, the error is compared against the golden file
// `expectedPth`.
func assertIsDeserializationErrorAndEqualsExpectedOrRecord(
	t *testing.T,
	err error,
	source string,
	expectedPth string,
) (ok bool) {
	ok = true

	if err == nil {
		ok = false
		t.Fatalf("De-serialization error expected from %s, but got none", source)
		return
	}

	deseriaErr, is := err.(*aasxmlization.DeserializationError)
	if !is {
		ok = false
		t.Fatalf(
			"Expected a de-serialization error, "+
				"but got an error of type %T from %s: %v",
			err, source, err,
		)
		return
	}

	// Add a new line for POSIX systems.
	got := deseriaErr.Error() + "\n"

	if aastesting.RecordMode {
		parent := filepath.Dir(expectedPth)
		err := os.MkdirAll(parent, os.ModePerm)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to create the directory %s: %s", parent, err.Error(),
				),
			)
		}

		err = os.WriteFile(expectedPth, []byte(got), 0644)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to write to the file %s: %s", expectedPth, err.Error(),
				),
			)
		}
	} else {
		bb, err := os.ReadFile(expectedPth)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to read from file %s: %s", expectedPth, err.Error(),
				),
			)
		}

		expected := string(bb)

		// NOTE (mristin, 2023-06-07):
		// Git automatically strips and adds `\r`, so we have to remove it here
		// to obtain a canonical text.
		expected = strings.Replace(expected, "\r", "", -1)

		if expected != got {
			t.Fatalf(
				"What we got differs from the expected in %s. "+
					"We got:\n%s\nWe expected:\n%s",
				expectedPth, got, expected,
			)
			ok = false
		}
	}

	return
}
