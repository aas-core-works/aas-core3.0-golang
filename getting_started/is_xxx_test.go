package getting_started_test

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

func Example_isXxx() {
	// Create the first element
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDInt,
	)

	fmt.Printf("%v\n", aastypes.IsSubmodelElement(someElement))
	fmt.Printf("%v\n", aastypes.IsProperty(someElement))
	fmt.Printf("%v\n", aastypes.IsBlob(someElement))
	// Output:
	// true
	// true
	// false
}
