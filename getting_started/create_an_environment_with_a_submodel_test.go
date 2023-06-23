package getting_started_test

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

func Example_createAnEnvironmentWithASubmodel() {
	// Create a new instance of the `value` and return the pointer to it.
	NewString := func(value string) *string {
		return &value
	}

	// Create the first element
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	someElement.SetIDShort(
		NewString("someProperty"),
	)
	someElement.SetValue(
		NewString("some-value"),
	)

	// Create the second element
	anotherElement := aastypes.NewBlob(
		"application/octet-stream",
	)
	anotherElement.SetIDShort(
		NewString("someBlob"),
	)
	anotherElement.SetValue(
		[]byte{0xDE, 0xAD, 0xBE, 0xEF},
	)

	// Nest the elements in a submodel
	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
			anotherElement,
		},
	)

	// Now create the environment to wrap it all up
	environment := aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// You can set the properties.
	environment.Submodels()[0].SubmodelElements()[0].(aastypes.IProperty).SetValue(
		NewString("changed-value"),
	)

	// You can access the properties from the children.
	fmt.Printf(
		"%v\n",
		*environment.Submodels()[0].SubmodelElements()[0].(aastypes.IProperty).Value(),
	)
	// Output:
	// changed-value
}
