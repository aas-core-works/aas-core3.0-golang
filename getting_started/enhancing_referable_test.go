package getting_started_test

import (
	"fmt"
	aasenhancing "github.com/aas-core-works/aas-core3.0-golang/enhancing"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

type ReferableEnhancement struct {
	ID int
}

func Example_enhancingReferable() {
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
	someElement.SetValue(
		NewString("some-value"),
	)

	administrativeInfo := aastypes.NewAdministrativeInformation()
	administrativeInfo.SetVersion(
		NewString("1.0"),
	)

	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
		},
	)
	submodel.SetAdministration(administrativeInfo)

	var environment aastypes.IEnvironment
	environment = aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Wrap everything
	nextID := 0

	factory := func(that aastypes.IClass) (enh *ReferableEnhancement, shouldEnh bool) {
		if aastypes.IsReferable(that) {
			enh = &ReferableEnhancement{ID: nextID}
			shouldEnh = true

			nextID++
		}

		return
	}

	environment = aasenhancing.Wrap[*ReferableEnhancement](
		environment,
		factory,
	).(aastypes.IEnvironment)

	environment.Descend(
		func(that aastypes.IClass) (abort bool) {
			enh, ok := aasenhancing.Unwrap[*ReferableEnhancement](that)
			if ok {
				fmt.Printf(
					"%s ID: %d\n",
					aasstringification.MustModelTypeToString(that.ModelType()),
					enh.ID,
				)
			} else {
				fmt.Printf(
					"%s No ID\n",
					aasstringification.MustModelTypeToString(that.ModelType()),
				)
			}
			return
		},
	)
	// Output:
	// Submodel ID: 0
	// AdministrativeInformation No ID
	// Property ID: 1
}
