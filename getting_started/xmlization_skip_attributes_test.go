package getting_started_test

import (
	"encoding/xml"
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
)

// TokenReaderNoAttributes reads tokens from a reader and
// removes any attributes in the start elements.
type TokenReaderNoAttributes struct {
	r xml.TokenReader
}

func NewTokenReaderNoAttributes(r xml.TokenReader) *TokenReaderNoAttributes {
	return &TokenReaderNoAttributes{r: r}
}

func (trna *TokenReaderNoAttributes) Token() (xml.Token, error) {
	var token xml.Token
	var err error

	for {
		token, err = trna.r.Token()
		if err != nil {
			return token, err
		}

		var startElement xml.StartElement
		startElement, isStartElement := token.(xml.StartElement)

		if !isStartElement {
			return token, err
		}

		startElement.Attr = []xml.Attr{}
		return startElement, err
	}
}

func Example_xmlizationSkipAttributes() {
	text := `<environment
	xmlns="https://admin-shell.io/aas/3/0"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="https://admin-shell.io/aas/3/0 AAS.xsd"
>
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

	// Use a decoder which skips the XML declarations
	decoder := xml.NewTokenDecoder(
		NewTokenReaderNoAttributes(
			xml.NewDecoder(reader),
		),
	)

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
