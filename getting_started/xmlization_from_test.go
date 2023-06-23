package getting_started_test

import (
	"encoding/xml"
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
)

func Example_xmlizationFrom() {
	text := `<environment xmlns="https://admin-shell.io/aas/3/0">
  <submodels>
    <submodel>
      <id>some-unique-global-identifier</id>
      <submodelElements>
        <property>
          <idShort>someProperty</idShort>
          <valueType>xs:string</valueType>
          <value>some-value</value>
        </property>
      </submodelElements>
    </submodel>
  </submodels>
</environment>`

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
