package aastesting

import (
	"fmt"
	"os"
	"strings"
)

// NOTE (mristin, 2023-05-24):
// It is tedious to record manually all the expected error messages. Therefore we
// include this variable to steer the automatic recording. We intentionally
// intertwine the recording code with the test code to keep them close to each other
// so that they are easier to maintain.
var rM = os.Getenv("AAS_CORE_3_0_GOLANG_RECORD_MODE")
var RecordMode = rM == "1" || strings.ToLower(rM) == "true" || strings.ToLower(rM) == "on"

func getTestDataDir() string {
	variable := "AAS_CORE_3_0_GOLANG_TEST_DATA_DIR"
	val, ok := os.LookupEnv(variable)
	if !ok {
		panic(
			fmt.Sprintf(
				"Expected the environment variable to be set, but it was not: %s",
				variable,
			),
		)
	}
	return val
}

var TestDataDir = getTestDataDir()
