package getting_started_test

import (
	"encoding/xml"
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
)

func Example_xmlizationTo() {
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

	// Serialize to XML
	builder := new(strings.Builder)
	encoder := xml.NewEncoder(builder)
	encoder.Indent("", "  ")

	// We want to include the namespace in the root XML element.
	withNamespace := true

	var err error
	err = aasxmlization.Marshal(encoder, environment, withNamespace)
	if err != nil {
		panic(err.Error())
	}

	text := builder.String()

	fmt.Println(text)
	// Output:
	// <environment xmlns="https://admin-shell.io/aas/3/0">
	//   <submodels>
	//     <submodel>
	//       <id>some-unique-global-identifier</id>
	//       <submodelElements>
	//         <property>
	//           <idShort>someProperty</idShort>
	//           <valueType>xs:string</valueType>
	//           <value>some-value</value>
	//         </property>
	//       </submodelElements>
	//     </submodel>
	//   </submodels>
	// </environment>
}
