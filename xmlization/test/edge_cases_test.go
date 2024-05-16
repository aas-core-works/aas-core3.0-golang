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
	text := ("<environment xmlns=\"https://admin-shell.io/aas/3/0\">" +
		"<submodels>" +
		"<submodel><id>some-unique-global-identifier</id></submodel>" +
		"</submodels>" +
		"</environment>")

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
	// Note that both `dataSpecificationContent` and `dataSpecificationIec61360` are
	// on the same line (right in the middle).
	text := `<environment xmlns="https://admin-shell.io/aas/3/0">
  <conceptDescriptions>
    <conceptDescription>
      <id>0173-1#02-AAR529#004</id>
      <embeddedDataSpecifications>
        <embeddedDataSpecification>
          <dataSpecification>
            <type>ExternalReference</type>
            <keys>
              <key>
                <type>GlobalReference</type>
                <value>https://admin-shell.io/DataSpecificationTemplates/DataSpecificationIEC61360/3/0</value>
              </key>
            </keys>
          </dataSpecification>
          <dataSpecificationContent><dataSpecificationIec61360>
              <preferredName>
                <langStringPreferredNameTypeIec61360>
                  <language>en</language>
                  <text>kilogram</text>
                </langStringPreferredNameTypeIec61360>
              </preferredName>
            </dataSpecificationIec61360>
          </dataSpecificationContent>
        </embeddedDataSpecification>
      </embeddedDataSpecifications>
    </conceptDescription>
  </conceptDescriptions>
</environment>`

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
	<conceptDescriptions>
		<conceptDescription>
			<id>0173-1#02-AAR529#004</id>
			<embeddedDataSpecifications>
				<embeddedDataSpecification>
					<dataSpecificationContent>
						<dataSpecificationIec61360>
							<preferredName>
								<langStringPreferredNameTypeIec61360>
									<language>en</language>
									<text>kilogram</text>
								</langStringPreferredNameTypeIec61360>
							</preferredName>
						</dataSpecificationIec61360>
					</dataSpecificationContent>
					<dataSpecification>
						<type>ExternalReference</type>
						<keys>
							<key>
								<type>GlobalReference</type>
								<value>https://admin-shell.io/DataSpecificationTemplates/DataSpecificationIEC61360/3/0</value>
							</key>
						</keys>
					</dataSpecification>
				</embeddedDataSpecification>
			</embeddedDataSpecifications>
		</conceptDescription>
	</conceptDescriptions>
</environment>`

	if expected != roundTrip {
		t.Fatalf("Expected `%v`, got `%v`", expected, roundTrip)
	}
}
