package getting_started_test

import (
	"encoding/xml"
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
)

func Example_xmlizationNamespaceAlias() {
	text := `<aas:environment
	xmlns:aas="https://admin-shell.io/aas/3/0"
>
  <aas:submodels>
    <aas:submodel>
      <aas:id>some-unique-global-identifier</aas:id>
      <aas:submodelElements>
        <aas:property>
          <aas:idShort>someProperty</aas:idShort>
          <aas:valueType>xs:string</aas:valueType>
          <aas:value>some-value</aas:value>
        </aas:property>
      </aas:submodelElements>
    </aas:submodel>
  </aas:submodels>
</aas:environment>`

	reader := strings.NewReader(text)

	decoder := xml.NewDecoder(reader)

	var instance aastypes.IClass
	var err error
	instance, err = aasxmlization.Unmarshal(
		decoder,
	)
	if err != nil {
		panic(err.Error())
	}

	instance.Descend(
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
