package getting_started_test

import (
	"encoding/json"
	"fmt"
	aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

func Example_jsonizationTo() {
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

	// Serialize to jsonable
	var jsonable map[string]interface{}
	var seriaErr *aasjsonization.SerializationError
	jsonable, seriaErr = aasjsonization.ToJsonable(environment)
	if seriaErr != nil {
		panic(seriaErr.Error())
	}

	// Serialize jsonable to string
	var bb []byte
	var err error
	bb, err = json.MarshalIndent(jsonable, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	text := string(bb)

	fmt.Println(text)
	// Output:
	// {
	//   "submodels": [
	//     {
	//       "id": "some-unique-global-identifier",
	//       "modelType": "Submodel",
	//       "submodelElements": [
	//         {
	//           "idShort": "someProperty",
	//           "modelType": "Property",
	//           "value": "some-value",
	//           "valueType": "xs:string"
	//         }
	//       ]
	//     }
	//   ]
	// }
}
