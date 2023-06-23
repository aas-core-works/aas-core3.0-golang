package getting_started_test

import (
	"encoding/json"
	"fmt"
	aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

func Example_jsonizationFrom() {
	text := `
{
  "submodels": [
    {
      "id": "some-unique-global-identifier",
      "modelType": "Submodel",
      "submodelElements": [
        {
          "idShort": "someProperty",
          "modelType": "Property",
          "value": "some-value",
          "valueType": "xs:string"
        }
      ]
    }
  ]
}`

	bb := []byte(text)

	var jsonable map[string]interface{}
	var err error
	err = json.Unmarshal(bb, &jsonable)
	if err != nil {
		panic(err.Error())
	}

	var environment aastypes.IEnvironment
	var deseriaErr *aasjsonization.DeserializationError
	environment, deseriaErr = aasjsonization.EnvironmentFromJsonable(
		jsonable,
	)
	if deseriaErr != nil {
		panic(deseriaErr.Error())
	}

	environment.Descend(
		func(that aastypes.IClass) (abort bool) {
			fmt.Printf(
				"%s\n",
				aasstringification.MustModelTypeToString(that.ModelType()),
			)
			return
		},
	)

	// Output:
	// Submodel
	// Property
}
