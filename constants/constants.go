// Package constants provides constant values of the meta-model.

package constants

// This code has been automatically generated by aas-core-codegen.
// Do NOT edit or append.

import (
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

// Categories for [aastypes.IDataElement] as defined in Constraint AASd-090
var ValidCategoriesForDataElement = map[string]struct{}{
	"CONSTANT":  {},
	"PARAMETER": {},
	"VARIABLE":  {},
}

// Enumeration of all identifiable elements within an asset administration shell.
var GenericFragmentKeys = map[aastypes.KeyTypes]struct{}{
	aastypes.KeyTypesFragmentReference: {},
}

// Enumeration of different key value types within a key.
var GenericGloballyIdentifiables = map[aastypes.KeyTypes]struct{}{
	aastypes.KeyTypesGlobalReference: {},
}

// Enumeration of different key value types within a key.
var AASIdentifiables = map[aastypes.KeyTypes]struct{}{
	aastypes.KeyTypesAssetAdministrationShell: {},
	aastypes.KeyTypesConceptDescription:       {},
	aastypes.KeyTypesIdentifiable:             {},
	aastypes.KeyTypesSubmodel:                 {},
}

// Enumeration of all submodel elements within an asset administration shell.
var AASSubmodelElementsAsKeys = map[aastypes.KeyTypes]struct{}{
	aastypes.KeyTypesAnnotatedRelationshipElement: {},
	aastypes.KeyTypesBasicEventElement:            {},
	aastypes.KeyTypesBlob:                         {},
	aastypes.KeyTypesCapability:                   {},
	aastypes.KeyTypesDataElement:                  {},
	aastypes.KeyTypesEntity:                       {},
	aastypes.KeyTypesEventElement:                 {},
	aastypes.KeyTypesFile:                         {},
	aastypes.KeyTypesMultiLanguageProperty:        {},
	aastypes.KeyTypesOperation:                    {},
	aastypes.KeyTypesProperty:                     {},
	aastypes.KeyTypesRange:                        {},
	aastypes.KeyTypesReferenceElement:             {},
	aastypes.KeyTypesRelationshipElement:          {},
	aastypes.KeyTypesSubmodelElement:              {},
	aastypes.KeyTypesSubmodelElementCollection:    {},
	aastypes.KeyTypesSubmodelElementList:          {},
}

// Enumeration of different fragment key value types within a key.
var AASReferableNonIdentifiables = map[aastypes.KeyTypes]struct{}{
	aastypes.KeyTypesAnnotatedRelationshipElement: {},
	aastypes.KeyTypesBasicEventElement:            {},
	aastypes.KeyTypesBlob:                         {},
	aastypes.KeyTypesCapability:                   {},
	aastypes.KeyTypesDataElement:                  {},
	aastypes.KeyTypesEntity:                       {},
	aastypes.KeyTypesEventElement:                 {},
	aastypes.KeyTypesFile:                         {},
	aastypes.KeyTypesMultiLanguageProperty:        {},
	aastypes.KeyTypesOperation:                    {},
	aastypes.KeyTypesProperty:                     {},
	aastypes.KeyTypesRange:                        {},
	aastypes.KeyTypesReferenceElement:             {},
	aastypes.KeyTypesRelationshipElement:          {},
	aastypes.KeyTypesSubmodelElement:              {},
	aastypes.KeyTypesSubmodelElementCollection:    {},
	aastypes.KeyTypesSubmodelElementList:          {},
}

// Enumeration of referables. We need this to check that model references refer to a Referable. For example, the observed attribute of the Basic Event Element object must be a model reference to a Referable.
var AASReferables = map[aastypes.KeyTypes]struct{}{
	aastypes.KeyTypesAssetAdministrationShell:     {},
	aastypes.KeyTypesConceptDescription:           {},
	aastypes.KeyTypesIdentifiable:                 {},
	aastypes.KeyTypesSubmodel:                     {},
	aastypes.KeyTypesAnnotatedRelationshipElement: {},
	aastypes.KeyTypesBasicEventElement:            {},
	aastypes.KeyTypesBlob:                         {},
	aastypes.KeyTypesCapability:                   {},
	aastypes.KeyTypesDataElement:                  {},
	aastypes.KeyTypesEntity:                       {},
	aastypes.KeyTypesEventElement:                 {},
	aastypes.KeyTypesFile:                         {},
	aastypes.KeyTypesMultiLanguageProperty:        {},
	aastypes.KeyTypesOperation:                    {},
	aastypes.KeyTypesProperty:                     {},
	aastypes.KeyTypesRange:                        {},
	aastypes.KeyTypesReferenceElement:             {},
	aastypes.KeyTypesReferable:                    {},
	aastypes.KeyTypesRelationshipElement:          {},
	aastypes.KeyTypesSubmodelElement:              {},
	aastypes.KeyTypesSubmodelElementCollection:    {},
	aastypes.KeyTypesSubmodelElementList:          {},
}

// Enumeration of all referable elements within an asset administration shell
var GloballyIdentifiables = map[aastypes.KeyTypes]struct{}{
	aastypes.KeyTypesGlobalReference:          {},
	aastypes.KeyTypesAssetAdministrationShell: {},
	aastypes.KeyTypesConceptDescription:       {},
	aastypes.KeyTypesIdentifiable:             {},
	aastypes.KeyTypesSubmodel:                 {},
}

// Enumeration of different key value types within a key.
var FragmentKeys = map[aastypes.KeyTypes]struct{}{
	aastypes.KeyTypesAnnotatedRelationshipElement: {},
	aastypes.KeyTypesBasicEventElement:            {},
	aastypes.KeyTypesBlob:                         {},
	aastypes.KeyTypesCapability:                   {},
	aastypes.KeyTypesDataElement:                  {},
	aastypes.KeyTypesEntity:                       {},
	aastypes.KeyTypesEventElement:                 {},
	aastypes.KeyTypesFile:                         {},
	aastypes.KeyTypesFragmentReference:            {},
	aastypes.KeyTypesMultiLanguageProperty:        {},
	aastypes.KeyTypesOperation:                    {},
	aastypes.KeyTypesProperty:                     {},
	aastypes.KeyTypesRange:                        {},
	aastypes.KeyTypesReferenceElement:             {},
	aastypes.KeyTypesRelationshipElement:          {},
	aastypes.KeyTypesSubmodelElement:              {},
	aastypes.KeyTypesSubmodelElementCollection:    {},
	aastypes.KeyTypesSubmodelElementList:          {},
}

// IEC 61360 data types for concept descriptions categorized with PROPERTY or VALUE.
var DataTypeIEC61360ForPropertyOrValue = map[aastypes.DataTypeIEC61360]struct{}{
	aastypes.DataTypeIEC61360Date:               {},
	aastypes.DataTypeIEC61360String:             {},
	aastypes.DataTypeIEC61360StringTranslatable: {},
	aastypes.DataTypeIEC61360IntegerMeasure:     {},
	aastypes.DataTypeIEC61360IntegerCount:       {},
	aastypes.DataTypeIEC61360IntegerCurrency:    {},
	aastypes.DataTypeIEC61360RealMeasure:        {},
	aastypes.DataTypeIEC61360RealCount:          {},
	aastypes.DataTypeIEC61360RealCurrency:       {},
	aastypes.DataTypeIEC61360Boolean:            {},
	aastypes.DataTypeIEC61360Rational:           {},
	aastypes.DataTypeIEC61360RationalMeasure:    {},
	aastypes.DataTypeIEC61360Time:               {},
	aastypes.DataTypeIEC61360Timestamp:          {},
}

// IEC 61360 data types for concept descriptions categorized with REFERENCE.
var DataTypeIEC61360ForReference = map[aastypes.DataTypeIEC61360]struct{}{
	aastypes.DataTypeIEC61360String: {},
	aastypes.DataTypeIEC61360IRI:    {},
	aastypes.DataTypeIEC61360IRDI:   {},
}

// IEC 61360 data types for concept descriptions categorized with DOCUMENT.
var DataTypeIEC61360ForDocument = map[aastypes.DataTypeIEC61360]struct{}{
	aastypes.DataTypeIEC61360File: {},
	aastypes.DataTypeIEC61360Blob: {},
	aastypes.DataTypeIEC61360HTML: {},
}

// These data types imply that the unit is defined in the data specification.
var IEC61360DataTypesWithUnit = map[aastypes.DataTypeIEC61360]struct{}{
	aastypes.DataTypeIEC61360IntegerMeasure:  {},
	aastypes.DataTypeIEC61360RealMeasure:     {},
	aastypes.DataTypeIEC61360RationalMeasure: {},
	aastypes.DataTypeIEC61360IntegerCurrency: {},
	aastypes.DataTypeIEC61360RealCurrency:    {},
}

// This code has been automatically generated by aas-core-codegen.
// Do NOT edit or append.
