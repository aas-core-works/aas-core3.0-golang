package jsonization_test

// This code has been automatically generated by:
// _dev_scripts/test_codegen/generate_jsonization_of_enums_test.py
// Do NOT edit or append.

import (
	"fmt"
	aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
	"testing"
)

func TestModellingKindRoundTripOK(t *testing.T) {
	literals := []string{
		"Template",
		"Instance",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.ModellingKindFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.ModellingKindToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestModellingKindDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.ModellingKindFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of ModellingKind, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestQualifierKindRoundTripOK(t *testing.T) {
	literals := []string{
		"ValueQualifier",
		"ConceptQualifier",
		"TemplateQualifier",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.QualifierKindFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.QualifierKindToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestQualifierKindDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.QualifierKindFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of QualifierKind, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestAssetKindRoundTripOK(t *testing.T) {
	literals := []string{
		"Type",
		"Instance",
		"NotApplicable",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.AssetKindFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.AssetKindToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestAssetKindDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.AssetKindFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of AssetKind, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestAASSubmodelElementsRoundTripOK(t *testing.T) {
	literals := []string{
		"AnnotatedRelationshipElement",
		"BasicEventElement",
		"Blob",
		"Capability",
		"DataElement",
		"Entity",
		"EventElement",
		"File",
		"MultiLanguageProperty",
		"Operation",
		"Property",
		"Range",
		"ReferenceElement",
		"RelationshipElement",
		"SubmodelElement",
		"SubmodelElementList",
		"SubmodelElementCollection",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.AASSubmodelElementsFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.AASSubmodelElementsToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestAASSubmodelElementsDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.AASSubmodelElementsFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of AASSubmodelElements, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestEntityTypeRoundTripOK(t *testing.T) {
	literals := []string{
		"CoManagedEntity",
		"SelfManagedEntity",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.EntityTypeFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.EntityTypeToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestEntityTypeDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.EntityTypeFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of EntityType, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestDirectionRoundTripOK(t *testing.T) {
	literals := []string{
		"input",
		"output",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.DirectionFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.DirectionToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestDirectionDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.DirectionFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of Direction, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestStateOfEventRoundTripOK(t *testing.T) {
	literals := []string{
		"on",
		"off",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.StateOfEventFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.StateOfEventToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestStateOfEventDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.StateOfEventFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of StateOfEvent, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestReferenceTypesRoundTripOK(t *testing.T) {
	literals := []string{
		"ExternalReference",
		"ModelReference",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.ReferenceTypesFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.ReferenceTypesToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestReferenceTypesDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.ReferenceTypesFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of ReferenceTypes, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestKeyTypesRoundTripOK(t *testing.T) {
	literals := []string{
		"AnnotatedRelationshipElement",
		"AssetAdministrationShell",
		"BasicEventElement",
		"Blob",
		"Capability",
		"ConceptDescription",
		"DataElement",
		"Entity",
		"EventElement",
		"File",
		"FragmentReference",
		"GlobalReference",
		"Identifiable",
		"MultiLanguageProperty",
		"Operation",
		"Property",
		"Range",
		"Referable",
		"ReferenceElement",
		"RelationshipElement",
		"Submodel",
		"SubmodelElement",
		"SubmodelElementCollection",
		"SubmodelElementList",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.KeyTypesFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.KeyTypesToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestKeyTypesDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.KeyTypesFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of KeyTypes, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestDataTypeDefXSDRoundTripOK(t *testing.T) {
	literals := []string{
		"xs:anyURI",
		"xs:base64Binary",
		"xs:boolean",
		"xs:byte",
		"xs:date",
		"xs:dateTime",
		"xs:decimal",
		"xs:double",
		"xs:duration",
		"xs:float",
		"xs:gDay",
		"xs:gMonth",
		"xs:gMonthDay",
		"xs:gYear",
		"xs:gYearMonth",
		"xs:hexBinary",
		"xs:int",
		"xs:integer",
		"xs:long",
		"xs:negativeInteger",
		"xs:nonNegativeInteger",
		"xs:nonPositiveInteger",
		"xs:positiveInteger",
		"xs:short",
		"xs:string",
		"xs:time",
		"xs:unsignedByte",
		"xs:unsignedInt",
		"xs:unsignedLong",
		"xs:unsignedShort",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.DataTypeDefXSDFromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.DataTypeDefXSDToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestDataTypeDefXSDDeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.DataTypeDefXSDFromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of DataTypeDefXSD, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestDataTypeIEC61360RoundTripOK(t *testing.T) {
	literals := []string{
		"DATE",
		"STRING",
		"STRING_TRANSLATABLE",
		"INTEGER_MEASURE",
		"INTEGER_COUNT",
		"INTEGER_CURRENCY",
		"REAL_MEASURE",
		"REAL_COUNT",
		"REAL_CURRENCY",
		"BOOLEAN",
		"IRI",
		"IRDI",
		"RATIONAL",
		"RATIONAL_MEASURE",
		"TIME",
		"TIMESTAMP",
		"FILE",
		"HTML",
		"BLOB",
	}

	for _, literal := range literals {
		source := fmt.Sprintf("<string literal %s>", literal)
		jsonable := any(literal)

		deserialized, deseriaErr := aasjsonization.DataTypeIEC61360FromJsonable(
			jsonable,
		)
		ok := assertNoDeserializationError(t, deseriaErr, source)
		if !ok {
			return
		}

		anotherJsonable, seriaErr :=
			aasjsonization.DataTypeIEC61360ToJsonable(deserialized)
		ok = assertNoSerializationError(t, seriaErr, source)
		if !ok {
			return
		}

		ok = assertSerializationEqualsDeserialization(
			t,
			jsonable,
			anotherJsonable,
			source,
		)
		if !ok {
			return
		}
	}
}

func TestDataTypeIEC61360DeserializationFail(t *testing.T) {
	jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

	_, err := aasjsonization.DataTypeIEC61360FromJsonable(
		jsonable,
	)

	if err == nil {
		t.Fatal("Expected a deserialization error, but got none.")
		return
	}

	deseriaErr, ok := err.(*aasjsonization.DeserializationError)
	if !ok {
		t.Fatalf("Expected a de-serialization error, but got: %v", err)
		return
	}

	pathString := deseriaErr.PathString()
	if len(pathString) != 0 {
		t.Fatalf(
			"Expected an empty path in error, but got: %s",
			pathString,
		)
		return
	}

	expectedMessage :=
		"Expected a string representation of DataTypeIEC61360, " +
			"but got THIS-CANNOT-POSSIBLY-BE-VALID"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

// This code has been automatically generated by:
// _dev_scripts/test_codegen/generate_jsonization_of_enums_test.py
// Do NOT edit or append.
