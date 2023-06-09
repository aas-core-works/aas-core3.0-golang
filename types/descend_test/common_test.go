package types_descend_test

import (
	"fmt"
	aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	"os"
	"path/filepath"
	"strings"
)

// Trace the `instance` and compare the trace against the golden one from
// the test data, or re-record the trace if [aastesting.RecordMode] is set.
//
// If `onlyOnce`, trace the `instance` with [aastypes.DescendOnce]. Otherwise,
// trace with [aastypes.Descend].
//
// If we are comparing, and not recording, return the error message if
// the expected and the obtained trace differ.
func compareOrRerecordTrace(
	instance aastypes.IClass,
	expectedPath string,
	onlyOnce bool,
) (message *string) {
	lines := []string{aastesting.TraceMark(instance)}

	if onlyOnce {
		instance.DescendOnce(func(descendant aastypes.IClass) (abort bool) {
			lines = append(lines, aastesting.TraceMark(descendant))
			return
		})
	} else {
		instance.Descend(func(descendant aastypes.IClass) (abort bool) {
			lines = append(lines, aastesting.TraceMark(descendant))
			return
		})
	}

	got := strings.Join(lines, "\n")

	// Add a new line for POSIX systems.
	got += "\n"

	if aastesting.RecordMode {
		parent := filepath.Dir(expectedPath)
		err := os.MkdirAll(parent, os.ModePerm)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to create the directory %s: %s", parent, err.Error(),
				),
			)
		}

		err = os.WriteFile(expectedPath, []byte(got), 0644)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to write to the file %s: %s", expectedPath, err.Error(),
				),
			)
		}
	} else {
		bb, err := os.ReadFile(expectedPath)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to read from file %s: %s", expectedPath, err.Error(),
				),
			)
		}

		expected := string(bb)

		// NOTE (mristin, 2023-06-07):
		// Git automatically strips and adds `\r`, so we have to remove it here
		// to obtain a canonical text.
		expected = strings.Replace(expected, "\r", "", -1)

		if expected != got {
			text := fmt.Sprintf(
				"What we got differs from the expected in %s. "+
					"We got:\n%s\nWe expected:\n%s",
				expectedPath, got, expected,
			)
			message = &text
		}
	}
	return
}
