package xmlization_test

import (
	"bytes"
	"encoding/xml"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
	"testing"
)

func TestRobustToNewline(t *testing.T) {
	// See:
	// https://github.com/aas-core-works/aas-core3.0-golang/issues/24
	//
	// We relied erroneously on whitespace between the elements which eclipsed the bug
	// where we read into the content of a class property instead of letting the called
	// function move the cursor.
	text := "<environment xmlns=\"https://admin-shell.io/aas/3/0\">" +
		"<submodels>" +
		"<submodel><id>some-unique-global-identifier</id></submodel>" +
		"</submodels>" +
		"</environment>"

	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	deserialized, deseriaErr := aasxmlization.Unmarshal(decoder)

	source := "Minimal example to reproduce robustness to newlines"

	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	if _, ok := deserialized.(aastypes.IEnvironment); !ok {
		t.Fatalf(
			"Expected an instance of IEnvironment, "+
				"but got %T: %v",
			deserialized, deserialized,
		)
		return
	}

	buf := &bytes.Buffer{}
	encoder := xml.NewEncoder(buf)
	encoder.Indent("", "\t")

	seriaErr := aasxmlization.Marshal(encoder, deserialized, true)
	ok = assertNoSerializationError(t, seriaErr, source)
	if !ok {
		return
	}

	roundTrip := string(buf.Bytes())

	expected := `<environment xmlns="https://admin-shell.io/aas/3/0">
	<submodels>
		<submodel>
			<id>some-unique-global-identifier</id>
		</submodel>
	</submodels>
</environment>`
	if roundTrip != expected {
		t.Fatalf(
			"Expected round-trip serialization to be `%v`, but got `%v`",
			expected, roundTrip,
		)
	}
}

func TestRobustToNewlineOnDispatch(t *testing.T) {
	// See:
	// https://github.com/aas-core-works/aas-core3.0-golang/issues/24
	//
	// Note that both `idShort` and `valueType` are
	// on the same line (right in the middle).
	text := `<environment xmlns="https://admin-shell.io/aas/3/0">
	<submodels>
		<submodel>
			<id>something_48c66017</id>
			<submodelElements>
				<property><idShort>something3fdd3eb4</idShort><valueType>xs:decimal</valueType>
				</property>
			</submodelElements>
		</submodel>
	</submodels>
</environment>
`

	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	deserialized, deseriaErr := aasxmlization.Unmarshal(decoder)

	pth := "Issue #24"

	ok := assertNoDeserializationError(t, deseriaErr, pth)
	if !ok {
		return
	}

	if _, ok := deserialized.(aastypes.IEnvironment); !ok {
		t.Fatalf(
			"Expected an instance of IEnvironment, "+
				"but got %T: %v",
			deserialized, deserialized,
		)
		return
	}

	buf := &bytes.Buffer{}
	encoder := xml.NewEncoder(buf)
	encoder.Indent("", "\t")

	seriaErr := aasxmlization.Marshal(encoder, deserialized, true)
	ok = assertNoSerializationError(t, seriaErr, pth)
	if !ok {
		return
	}

	roundTrip := string(buf.Bytes())

	expected := `<environment xmlns="https://admin-shell.io/aas/3/0">
	<submodels>
		<submodel>
			<id>something_48c66017</id>
			<submodelElements>
				<property>
					<idShort>something3fdd3eb4</idShort>
					<valueType>xs:decimal</valueType>
				</property>
			</submodelElements>
		</submodel>
	</submodels>
</environment>`

	if expected != roundTrip {
		t.Fatalf("Expected `%v`, got `%v`", expected, roundTrip)
	}
}
