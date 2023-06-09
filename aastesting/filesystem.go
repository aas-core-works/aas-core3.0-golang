package aastesting

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Read the content of `pth` and parse it as a JSON.
//
// If any errors, panic.
func MustReadJsonable(pth string) (jsonable interface{}) {
	bb, err := os.ReadFile(pth)
	if err != nil {
		panic(
			fmt.Sprintf(
				"Failed to read the content of: %s",
				pth,
			),
		)
	}

	err = json.Unmarshal(bb, &jsonable)
	if err != nil {
		panic(
			fmt.Sprintf(
				"Failed to parse the content of %s as JSON: %s",
				pth, err.Error(),
			),
		)
	}
	return
}

func FindFilesBySuffixRecursively(root, suffix string) []string {
	var a []string

	_, statErr := os.Stat(root)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			return a
		}

		panic(fmt.Sprintf("Failed to stat %s: %s", root, statErr.Error()))
	}

	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == suffix {
			a = append(a, s)
		}
		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to walk %s: %s", root, err.Error()))
	}
	return a
}
