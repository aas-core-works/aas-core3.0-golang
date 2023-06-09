// Test [aastypes.Descend] functions.

package types_descend_test

// This code has been automatically generated by:
// _dev_scripts/test_codegen/generate_descend_test.py
// Do NOT edit or append.

import (
	aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
	"path/filepath"
	"testing"
)

func TestDescendOnAnInstanceOfExtension(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalExtension()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Extension",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfAdministrativeInformation(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalAdministrativeInformation()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"AdministrativeInformation",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfQualifier(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalQualifier()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Qualifier",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfAssetAdministrationShell(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalAssetAdministrationShell()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"AssetAdministrationShell",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfAssetInformation(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalAssetInformation()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"AssetInformation",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfResource(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalResource()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Resource",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfSpecificAssetID(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalSpecificAssetID()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"SpecificAssetId",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfSubmodel(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalSubmodel()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Submodel",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfRelationshipElement(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalRelationshipElement()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"RelationshipElement",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfSubmodelElementList(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalSubmodelElementList()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"SubmodelElementList",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfSubmodelElementCollection(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalSubmodelElementCollection()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"SubmodelElementCollection",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfProperty(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalProperty()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Property",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfMultiLanguageProperty(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalMultiLanguageProperty()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"MultiLanguageProperty",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfRange(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalRange()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Range",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfReferenceElement(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalReferenceElement()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"ReferenceElement",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfBlob(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalBlob()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Blob",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfFile(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalFile()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"File",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfAnnotatedRelationshipElement(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalAnnotatedRelationshipElement()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"AnnotatedRelationshipElement",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfEntity(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalEntity()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Entity",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfEventPayload(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalEventPayload()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"EventPayload",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfBasicEventElement(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalBasicEventElement()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"BasicEventElement",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfOperation(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalOperation()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Operation",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfOperationVariable(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalOperationVariable()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"OperationVariable",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfCapability(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalCapability()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Capability",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfConceptDescription(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalConceptDescription()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"ConceptDescription",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfReference(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalReference()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Reference",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfKey(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalKey()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Key",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfLangStringNameType(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalLangStringNameType()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"LangStringNameType",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfLangStringTextType(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalLangStringTextType()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"LangStringTextType",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfEnvironment(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalEnvironment()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"Environment",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfEmbeddedDataSpecification(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalEmbeddedDataSpecification()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"EmbeddedDataSpecification",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfLevelType(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalLevelType()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"LevelType",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfValueReferencePair(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalValueReferencePair()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"ValueReferencePair",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfValueList(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalValueList()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"ValueList",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfLangStringPreferredNameTypeIEC61360(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalLangStringPreferredNameTypeIEC61360()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"LangStringPreferredNameTypeIec61360",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfLangStringShortNameTypeIEC61360(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalLangStringShortNameTypeIEC61360()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"LangStringShortNameTypeIec61360",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfLangStringDefinitionTypeIEC61360(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalLangStringDefinitionTypeIEC61360()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"LangStringDefinitionTypeIec61360",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

func TestDescendOnAnInstanceOfDataSpecificationIEC61360(
	t *testing.T,
) {
	instance := aastesting.MustLoadMaximalDataSpecificationIEC61360()

	expectedPth := filepath.Join(
		aastesting.TestDataDir,
		"Descend",
		"DataSpecificationIec61360",
		"maximal.json.trace",
	)

	onlyOnce := false

	message := compareOrRerecordTrace(
		instance,
		expectedPth,
		onlyOnce,
	)
	if message != nil {
		t.Fatal(*message)
	}
}

// This code has been automatically generated by:
// _dev_scripts/test_codegen/generate_descend_test.py
// Do NOT edit or append.
