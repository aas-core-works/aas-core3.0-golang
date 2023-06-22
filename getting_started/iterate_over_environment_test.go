package getting_started_test

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	"strings"
)

func Example_iterateOverEnvironment() {
	// Create a new instance of the `value` and return the pointer to it.
	NewString := func(value string) *string {
		return &value
	}

	// Prepare the environment
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	someElement.SetIDShort(
		NewString("someProperty"),
	)

	anotherElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	anotherElement.SetIDShort(
		NewString("anotherProperty"),
	)

	yetAnotherElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	yetAnotherElement.SetIDShort(
		NewString("yetAnotherProperty"),
	)

	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
			anotherElement,
			yetAnotherElement,
		},
	)

	environment := aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Iterate using ``Descend``
	environment.Descend(
		func(that aastypes.IClass) (abort bool) {
			if aastypes.IsProperty(that) {
				idShort := that.(aastypes.IProperty).IDShort()
				if idShort != nil &&
					strings.Contains(strings.ToLower(*idShort), "another") {
					fmt.Printf("%s\n", *idShort)
				}
			}
			return
		},
	)
	// Output:
	// anotherProperty
	// yetAnotherProperty
}
