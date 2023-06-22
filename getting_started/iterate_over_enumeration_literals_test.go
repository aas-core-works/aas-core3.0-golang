package getting_started_test

import (
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

func Example_iterateOverEnumerationLiterals() {
	for _, literal := range aastypes.LiteralsOfModellingKind {
		fmt.Printf(
			"Literal as number: %d, literal as string: %s\n",
			literal, aasstringification.MustModellingKindToString(literal),
		)
	}
	// Output:
	// Literal as number: 0, literal as string: Template
	// Literal as number: 1, literal as string: Instance
}
