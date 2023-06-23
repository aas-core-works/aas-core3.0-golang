package getting_started_test

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasverification "github.com/aas-core-works/aas-core3.0-golang/verification"
)

func Example_verification() {
	// Create a new instance of the `value` and return the pointer to it.
	NewString := func(value string) *string {
		return &value
	}

	// Prepare the environment
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	// The ID-shorts must be proper variable names,
	// but there is a dash (`-`) in this ID-short.
	someElement.SetIDShort(
		NewString("some-property"),
	)

	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
		},
	)

	environment := aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Verify
	aasverification.Verify(
		environment,
		func(err *aasverification.VerificationError) (abort bool) {
			fmt.Printf("%s\n", err.Error())
			return
		},
	)
	// Output:
	// Submodels[0].SubmodelElements[0].IDShort: ID-short of Referables shall only feature letters, digits, underscore (``_``); starting mandatory with a letter. *I.e.* ``[a-zA-Z][a-zA-Z0-9_]*``.
}
