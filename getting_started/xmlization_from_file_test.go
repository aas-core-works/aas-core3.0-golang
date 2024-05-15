package getting_started_test

import (
	"encoding/xml"
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"log"
	"strings"
)

// TokenReaderSkipDeclaration reads tokens from a reader and
// skips the XML declarations.
type TokenReaderSkipDeclaration struct {
	r xml.TokenReader
}

func NewTokenReaderSkipDeclaration(r xml.TokenReader) *TokenReaderSkipDeclaration {
	return &TokenReaderSkipDeclaration{r: r}
}

func (trsd *TokenReaderSkipDeclaration) Token() (xml.Token, error) {
	var token xml.Token
	var err error

	for {
		token, err = trsd.r.Token()
		if err != nil {
			return token, err
		}

		if _, isProcInst := token.(xml.ProcInst); !isProcInst {
			return token, err
		}
	}
}

func Example_xmlizationFromTextWithBOMAndXMLDeclaration() {
	// "\uFEFF" is a Byte Order Mark.
	text := "\uFEFF" + `<?xml version="1.0" encoding="UTF-8" ?>
<environment xmlns="https://admin-shell.io/aas/3/0">
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

	// Try to read the Byte Order Mark
	var err error
	rune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	if rune != '\uFEFF' {
		reader.UnreadRune()
	}

	// Use a decoder which skips the XML declarations
	decoder := xml.NewTokenDecoder(
		NewTokenReaderSkipDeclaration(
			xml.NewDecoder(reader),
		),
	)

	var instance aastypes.IClass
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
