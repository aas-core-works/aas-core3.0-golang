package jsonization_test

// This code has been automatically generated by:
// _dev_scripts/test_codegen/generate_jsonization_of_classes_with_descendants_test.py
// Do NOT edit or append.

import (
	aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
	aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
	"testing"
)

func TestHasSemanticsRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalRelationshipElement()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal RelationshipElement: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal RelationshipElement>"

	deserialized, deseriaErr := aasjsonization.HasSemanticsFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestHasSemanticsDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.HasSemanticsFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestHasExtensionsRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalRelationshipElement()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal RelationshipElement: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal RelationshipElement>"

	deserialized, deseriaErr := aasjsonization.HasExtensionsFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestHasExtensionsDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.HasExtensionsFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestReferableRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalRelationshipElement()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal RelationshipElement: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal RelationshipElement>"

	deserialized, deseriaErr := aasjsonization.ReferableFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestReferableDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.ReferableFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestIdentifiableRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalAssetAdministrationShell()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal AssetAdministrationShell: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal AssetAdministrationShell>"

	deserialized, deseriaErr := aasjsonization.IdentifiableFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestIdentifiableDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.IdentifiableFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestHasKindRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalSubmodel()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal Submodel: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal Submodel>"

	deserialized, deseriaErr := aasjsonization.HasKindFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestHasKindDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.HasKindFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestHasDataSpecificationRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalRelationshipElement()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal RelationshipElement: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal RelationshipElement>"

	deserialized, deseriaErr := aasjsonization.HasDataSpecificationFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestHasDataSpecificationDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.HasDataSpecificationFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestQualifiableRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalRelationshipElement()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal RelationshipElement: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal RelationshipElement>"

	deserialized, deseriaErr := aasjsonization.QualifiableFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestQualifiableDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.QualifiableFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestSubmodelElementRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalRelationshipElement()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal RelationshipElement: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal RelationshipElement>"

	deserialized, deseriaErr := aasjsonization.SubmodelElementFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestSubmodelElementDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.SubmodelElementFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestRelationshipElementRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalAnnotatedRelationshipElement()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal AnnotatedRelationshipElement: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal AnnotatedRelationshipElement>"

	deserialized, deseriaErr := aasjsonization.RelationshipElementFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestDataElementRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalBlob()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal Blob: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal Blob>"

	deserialized, deseriaErr := aasjsonization.DataElementFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestDataElementDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.DataElementFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestEventElementRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalBasicEventElement()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal BasicEventElement: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal BasicEventElement>"

	deserialized, deseriaErr := aasjsonization.EventElementFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestEventElementDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.EventElementFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}

func TestDataSpecificationContentRoundTripOKOverDescendant(t *testing.T) {
	instance := aastesting.MustLoadMinimalDataSpecificationIEC61360()

	jsonable, err := aasjsonization.ToJsonable(instance)
	if err != nil {
		t.Fatalf(
			"Failed to serialize the minimal DataSpecificationIec61360: %s",
			err.Error(),
		)
		return
	}

	source := "<minimal DataSpecificationIec61360>"

	deserialized, deseriaErr := aasjsonization.DataSpecificationContentFromJsonable(
		jsonable,
	)
	ok := assertNoDeserializationError(t, deseriaErr, source)
	if !ok {
		return
	}

	anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
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

func TestDataSpecificationContentDeserializationFail(t *testing.T) {
	jsonable := any("this is not an object")

	_, deseriaErr := aasjsonization.DataSpecificationContentFromJsonable(
		jsonable,
	)

	if deseriaErr == nil {
		t.Fatal("Expected a deserialization error, but got none.")
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
		"Expected a JSON object, but got string"

	if deseriaErr.Message != expectedMessage {
		t.Fatalf(
			"Expected the deserialization error:\n%s\n, but got:\n%s",
			expectedMessage, deseriaErr.Message,
		)
		return
	}
}
