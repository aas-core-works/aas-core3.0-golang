package jsonization_test

import (
	"encoding/json"
	"fmt"
	aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
	aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
	"os"
	"path/filepath"
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
			"Expected no de-serialization error from %s, "+
				"but got: %v",
			source, err,
		)
		return
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
				"the instance obtained from %s, "+
				"but got: %v",
			source, err,
		)
		return
	}
	return
}

// Assert that the serialization `other`, as JSON-able, equals the original
// JSON-able `that` read from `source`.
func assertSerializationEqualsDeserialization(
	t *testing.T,
	that interface{},
	other interface{},
	source string,
) (ok bool) {
	ok = true

	thatBytes, err := json.Marshal(that)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal that jsonable %v: %s", that, err.Error()))
	}
	thatText := string(thatBytes)

	otherBytes, err := json.Marshal(other)
	if err != nil {
		panic(
			fmt.Sprintf(
				"Failed to marshal other jsonable %v: %s", other, err.Error(),
			),
		)
	}
	otherText := string(otherBytes)

	if thatText != otherText {
		ok = false
		t.Fatalf(
			"The serialization of the de-serialized instance from %s does not equal "+
				"the original JSON-able:\nOriginal:\n%s\nSerialized:\n%s",
			source, thatText, otherText,
		)
		return
	}

	return
}

// Assert that there is a de-serialization error.
//
// If [aastesting.RecordMode] is set, the de-serialization error is re-recorded
// to `expectedPth`. Otherwise, the error is compared against the golden file
// `expectedPth`.
func assertDeserializationErrorEqualsExpectedOrRecord(
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

	var deseriaErr *aasjsonization.DeserializationError
	deseriaErr, ok = err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v from %s", err, source)
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
		_, err := os.Stat(expectedPth)
		if err != nil {
			ok = false
			t.Fatalf(
				"Failed to stat the file %s: %s; if the file does not exist, "+
					"you probably want to record the test data by "+
					"setting the environment variable %s",
				expectedPth, err.Error(), aastesting.RecordModeEnvironmentVariableName,
			)
			return
		}

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
			ok = false
			t.Fatalf(
				"What we got differs from the expected in %s. "+
					"We got:\n%s\nWe expected:\n%s",
				expectedPth, got, expected,
			)
			return
		}
	}

	return
}
