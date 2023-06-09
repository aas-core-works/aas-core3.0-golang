package verification_test

import (
	"encoding/json"
	"fmt"
	aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
	aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasverification "github.com/aas-core-works/aas-core3.0-golang/verification"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Assert that there are no verification errors in the `instance` de-serialized
// from `source`.
func assertNoVerificationErrors(
	t *testing.T,
	instance aastypes.IClass,
	source string,
) (ok bool) {
	errors := make([]*aasverification.VerificationError, 0)
	aasverification.Verify(
		instance,
		func(err *aasverification.VerificationError) (abort bool) {
			errors = append(errors, err)
			return
		},
	)

	ok = true
	if len(errors) > 0 {
		ok = false

		var sb strings.Builder

		sb.WriteString(
			fmt.Sprintf(
				"Expected no errors when verifying the instance de-serialized from "+
					"%s, but got %d error(s)\n",
				source, len(errors),
			),
		)

		for i, err := range errors {
			sb.WriteString(
				fmt.Sprintf(
					"Error %d:\n%s: %s\n",
					i+1,
					err.PathString(),
					err.Message,
				),
			)
		}

		jsonable, seriaErr := aasjsonization.ToJsonable(instance)
		if seriaErr != nil {
			panic(
				fmt.Sprintf(
					"Failed to serialize instance to JSON obtained from %s: %s",
					source, seriaErr.Error(),
				),
			)
		}
		jsonableBytes, err := json.Marshal(jsonable)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to marshal to JSON an instance serialized from %s: %s",
					source, err.Error(),
				),
			)
		}

		sb.WriteString("Instance:\n")
		sb.WriteString(string(jsonableBytes))

		t.Fatal(sb.String())
	}
	return
}

var causesForVerificationFailure = [...]string{
	"DateTimeStampUtcViolationOnFebruary29th",
	"MaxLengthViolation",
	"MinLengthViolation",
	"PatternViolation",
	"InvalidValueExample",
	"InvalidMinMaxExample",
	"SetViolation",
	"ConstraintViolation",
}

// Assert that either the verification errors match the recorded ones at `pth`, if
// [aastesting.RecordMode] is set, or re-record the verification errors at `pth`.
func assertEqualsExpectedOrRerecordVerificationErrors(
	t *testing.T,
	errors []*aasverification.VerificationError,
	source string,
	expectedPth string,
) (ok bool) {
	ok = true
	if len(errors) == 0 {
		ok = false
		t.Fatalf(
			"Expected at least one verification error, "+
				"but got none when verifying the model loaded from: %s",
			source,
		)
	}

	parts := make([]string, len(errors))
	for i, verErr := range errors {
		parts[i] = fmt.Sprintf(
			"%s: %s",
			verErr.PathString(),
			verErr.Message,
		)
	}

	// Add a newline for POSIX systems
	got := strings.Replace(strings.Join(parts, ";\n"), "\r\n", "\n", -1) + "\n"

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
			t.Fatalf("Failed to write to %s: %s", expectedPth, err.Error())
		}
	} else {
		b, err := os.ReadFile(expectedPth)
		if err != nil {
			t.Fatalf("Failed to read from %s: %s", expectedPth, err.Error())
		}
		expected := strings.ReplaceAll(string(b), "\r\n", "\n")

		if expected != got {
			ok = false
			t.Fatalf(
				"The expected verification errors (read from %s) in the model "+
					"de-serialized from %s do not match the obtained ones. "+
					"Expected:\n%s\nGot:\n%s",
				expectedPth, source, expected, got,
			)
		}
	}
	return
}
