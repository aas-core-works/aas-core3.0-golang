package aastesting

// This code has been automatically generated by:
// _dev_scripts/test_codegen/generate_finding.py
// Do NOT edit or append.

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

// Find the first instance of [aastypes.IHasSemantics] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindHasSemantics(
	container aastypes.IClass,
) (result aastypes.IHasSemantics) {
	ok := aastypes.IsHasSemantics(container)
	if ok {
		result = container.(aastypes.IHasSemantics)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsHasSemantics(that)
		if abort {
			result = that.(aastypes.IHasSemantics)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IHasSemantics "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IExtension] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindExtension(
	container aastypes.IClass,
) (result aastypes.IExtension) {
	ok := aastypes.IsExtension(container)
	if ok {
		result = container.(aastypes.IExtension)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsExtension(that)
		if abort {
			result = that.(aastypes.IExtension)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IExtension "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IHasExtensions] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindHasExtensions(
	container aastypes.IClass,
) (result aastypes.IHasExtensions) {
	ok := aastypes.IsHasExtensions(container)
	if ok {
		result = container.(aastypes.IHasExtensions)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsHasExtensions(that)
		if abort {
			result = that.(aastypes.IHasExtensions)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IHasExtensions "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IReferable] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindReferable(
	container aastypes.IClass,
) (result aastypes.IReferable) {
	ok := aastypes.IsReferable(container)
	if ok {
		result = container.(aastypes.IReferable)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsReferable(that)
		if abort {
			result = that.(aastypes.IReferable)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IReferable "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IIdentifiable] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindIdentifiable(
	container aastypes.IClass,
) (result aastypes.IIdentifiable) {
	ok := aastypes.IsIdentifiable(container)
	if ok {
		result = container.(aastypes.IIdentifiable)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsIdentifiable(that)
		if abort {
			result = that.(aastypes.IIdentifiable)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IIdentifiable "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IHasKind] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindHasKind(
	container aastypes.IClass,
) (result aastypes.IHasKind) {
	ok := aastypes.IsHasKind(container)
	if ok {
		result = container.(aastypes.IHasKind)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsHasKind(that)
		if abort {
			result = that.(aastypes.IHasKind)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IHasKind "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IHasDataSpecification] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindHasDataSpecification(
	container aastypes.IClass,
) (result aastypes.IHasDataSpecification) {
	ok := aastypes.IsHasDataSpecification(container)
	if ok {
		result = container.(aastypes.IHasDataSpecification)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsHasDataSpecification(that)
		if abort {
			result = that.(aastypes.IHasDataSpecification)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IHasDataSpecification "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IAdministrativeInformation] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindAdministrativeInformation(
	container aastypes.IClass,
) (result aastypes.IAdministrativeInformation) {
	ok := aastypes.IsAdministrativeInformation(container)
	if ok {
		result = container.(aastypes.IAdministrativeInformation)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsAdministrativeInformation(that)
		if abort {
			result = that.(aastypes.IAdministrativeInformation)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IAdministrativeInformation "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IQualifiable] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindQualifiable(
	container aastypes.IClass,
) (result aastypes.IQualifiable) {
	ok := aastypes.IsQualifiable(container)
	if ok {
		result = container.(aastypes.IQualifiable)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsQualifiable(that)
		if abort {
			result = that.(aastypes.IQualifiable)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IQualifiable "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IQualifier] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindQualifier(
	container aastypes.IClass,
) (result aastypes.IQualifier) {
	ok := aastypes.IsQualifier(container)
	if ok {
		result = container.(aastypes.IQualifier)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsQualifier(that)
		if abort {
			result = that.(aastypes.IQualifier)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IQualifier "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IAssetAdministrationShell] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindAssetAdministrationShell(
	container aastypes.IClass,
) (result aastypes.IAssetAdministrationShell) {
	ok := aastypes.IsAssetAdministrationShell(container)
	if ok {
		result = container.(aastypes.IAssetAdministrationShell)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsAssetAdministrationShell(that)
		if abort {
			result = that.(aastypes.IAssetAdministrationShell)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IAssetAdministrationShell "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IAssetInformation] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindAssetInformation(
	container aastypes.IClass,
) (result aastypes.IAssetInformation) {
	ok := aastypes.IsAssetInformation(container)
	if ok {
		result = container.(aastypes.IAssetInformation)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsAssetInformation(that)
		if abort {
			result = that.(aastypes.IAssetInformation)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IAssetInformation "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IResource] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindResource(
	container aastypes.IClass,
) (result aastypes.IResource) {
	ok := aastypes.IsResource(container)
	if ok {
		result = container.(aastypes.IResource)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsResource(that)
		if abort {
			result = that.(aastypes.IResource)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IResource "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ISpecificAssetID] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindSpecificAssetID(
	container aastypes.IClass,
) (result aastypes.ISpecificAssetID) {
	ok := aastypes.IsSpecificAssetID(container)
	if ok {
		result = container.(aastypes.ISpecificAssetID)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsSpecificAssetID(that)
		if abort {
			result = that.(aastypes.ISpecificAssetID)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ISpecificAssetID "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ISubmodel] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindSubmodel(
	container aastypes.IClass,
) (result aastypes.ISubmodel) {
	ok := aastypes.IsSubmodel(container)
	if ok {
		result = container.(aastypes.ISubmodel)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsSubmodel(that)
		if abort {
			result = that.(aastypes.ISubmodel)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ISubmodel "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ISubmodelElement] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindSubmodelElement(
	container aastypes.IClass,
) (result aastypes.ISubmodelElement) {
	ok := aastypes.IsSubmodelElement(container)
	if ok {
		result = container.(aastypes.ISubmodelElement)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsSubmodelElement(that)
		if abort {
			result = that.(aastypes.ISubmodelElement)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ISubmodelElement "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IRelationshipElement] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindRelationshipElement(
	container aastypes.IClass,
) (result aastypes.IRelationshipElement) {
	ok := aastypes.IsRelationshipElement(container)
	if ok {
		result = container.(aastypes.IRelationshipElement)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsRelationshipElement(that)
		if abort {
			result = that.(aastypes.IRelationshipElement)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IRelationshipElement "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ISubmodelElementList] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindSubmodelElementList(
	container aastypes.IClass,
) (result aastypes.ISubmodelElementList) {
	ok := aastypes.IsSubmodelElementList(container)
	if ok {
		result = container.(aastypes.ISubmodelElementList)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsSubmodelElementList(that)
		if abort {
			result = that.(aastypes.ISubmodelElementList)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ISubmodelElementList "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ISubmodelElementCollection] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindSubmodelElementCollection(
	container aastypes.IClass,
) (result aastypes.ISubmodelElementCollection) {
	ok := aastypes.IsSubmodelElementCollection(container)
	if ok {
		result = container.(aastypes.ISubmodelElementCollection)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsSubmodelElementCollection(that)
		if abort {
			result = that.(aastypes.ISubmodelElementCollection)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ISubmodelElementCollection "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IDataElement] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindDataElement(
	container aastypes.IClass,
) (result aastypes.IDataElement) {
	ok := aastypes.IsDataElement(container)
	if ok {
		result = container.(aastypes.IDataElement)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsDataElement(that)
		if abort {
			result = that.(aastypes.IDataElement)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IDataElement "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IProperty] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindProperty(
	container aastypes.IClass,
) (result aastypes.IProperty) {
	ok := aastypes.IsProperty(container)
	if ok {
		result = container.(aastypes.IProperty)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsProperty(that)
		if abort {
			result = that.(aastypes.IProperty)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IProperty "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IMultiLanguageProperty] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindMultiLanguageProperty(
	container aastypes.IClass,
) (result aastypes.IMultiLanguageProperty) {
	ok := aastypes.IsMultiLanguageProperty(container)
	if ok {
		result = container.(aastypes.IMultiLanguageProperty)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsMultiLanguageProperty(that)
		if abort {
			result = that.(aastypes.IMultiLanguageProperty)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IMultiLanguageProperty "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IRange] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindRange(
	container aastypes.IClass,
) (result aastypes.IRange) {
	ok := aastypes.IsRange(container)
	if ok {
		result = container.(aastypes.IRange)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsRange(that)
		if abort {
			result = that.(aastypes.IRange)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IRange "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IReferenceElement] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindReferenceElement(
	container aastypes.IClass,
) (result aastypes.IReferenceElement) {
	ok := aastypes.IsReferenceElement(container)
	if ok {
		result = container.(aastypes.IReferenceElement)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsReferenceElement(that)
		if abort {
			result = that.(aastypes.IReferenceElement)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IReferenceElement "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IBlob] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindBlob(
	container aastypes.IClass,
) (result aastypes.IBlob) {
	ok := aastypes.IsBlob(container)
	if ok {
		result = container.(aastypes.IBlob)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsBlob(that)
		if abort {
			result = that.(aastypes.IBlob)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IBlob "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IFile] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindFile(
	container aastypes.IClass,
) (result aastypes.IFile) {
	ok := aastypes.IsFile(container)
	if ok {
		result = container.(aastypes.IFile)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsFile(that)
		if abort {
			result = that.(aastypes.IFile)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IFile "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IAnnotatedRelationshipElement] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindAnnotatedRelationshipElement(
	container aastypes.IClass,
) (result aastypes.IAnnotatedRelationshipElement) {
	ok := aastypes.IsAnnotatedRelationshipElement(container)
	if ok {
		result = container.(aastypes.IAnnotatedRelationshipElement)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsAnnotatedRelationshipElement(that)
		if abort {
			result = that.(aastypes.IAnnotatedRelationshipElement)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IAnnotatedRelationshipElement "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IEntity] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindEntity(
	container aastypes.IClass,
) (result aastypes.IEntity) {
	ok := aastypes.IsEntity(container)
	if ok {
		result = container.(aastypes.IEntity)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsEntity(that)
		if abort {
			result = that.(aastypes.IEntity)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IEntity "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IEventPayload] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindEventPayload(
	container aastypes.IClass,
) (result aastypes.IEventPayload) {
	ok := aastypes.IsEventPayload(container)
	if ok {
		result = container.(aastypes.IEventPayload)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsEventPayload(that)
		if abort {
			result = that.(aastypes.IEventPayload)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IEventPayload "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IEventElement] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindEventElement(
	container aastypes.IClass,
) (result aastypes.IEventElement) {
	ok := aastypes.IsEventElement(container)
	if ok {
		result = container.(aastypes.IEventElement)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsEventElement(that)
		if abort {
			result = that.(aastypes.IEventElement)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IEventElement "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IBasicEventElement] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindBasicEventElement(
	container aastypes.IClass,
) (result aastypes.IBasicEventElement) {
	ok := aastypes.IsBasicEventElement(container)
	if ok {
		result = container.(aastypes.IBasicEventElement)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsBasicEventElement(that)
		if abort {
			result = that.(aastypes.IBasicEventElement)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IBasicEventElement "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IOperation] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindOperation(
	container aastypes.IClass,
) (result aastypes.IOperation) {
	ok := aastypes.IsOperation(container)
	if ok {
		result = container.(aastypes.IOperation)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsOperation(that)
		if abort {
			result = that.(aastypes.IOperation)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IOperation "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IOperationVariable] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindOperationVariable(
	container aastypes.IClass,
) (result aastypes.IOperationVariable) {
	ok := aastypes.IsOperationVariable(container)
	if ok {
		result = container.(aastypes.IOperationVariable)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsOperationVariable(that)
		if abort {
			result = that.(aastypes.IOperationVariable)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IOperationVariable "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ICapability] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindCapability(
	container aastypes.IClass,
) (result aastypes.ICapability) {
	ok := aastypes.IsCapability(container)
	if ok {
		result = container.(aastypes.ICapability)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsCapability(that)
		if abort {
			result = that.(aastypes.ICapability)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ICapability "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IConceptDescription] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindConceptDescription(
	container aastypes.IClass,
) (result aastypes.IConceptDescription) {
	ok := aastypes.IsConceptDescription(container)
	if ok {
		result = container.(aastypes.IConceptDescription)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsConceptDescription(that)
		if abort {
			result = that.(aastypes.IConceptDescription)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IConceptDescription "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IReference] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindReference(
	container aastypes.IClass,
) (result aastypes.IReference) {
	ok := aastypes.IsReference(container)
	if ok {
		result = container.(aastypes.IReference)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsReference(that)
		if abort {
			result = that.(aastypes.IReference)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IReference "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IKey] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindKey(
	container aastypes.IClass,
) (result aastypes.IKey) {
	ok := aastypes.IsKey(container)
	if ok {
		result = container.(aastypes.IKey)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsKey(that)
		if abort {
			result = that.(aastypes.IKey)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IKey "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IAbstractLangString] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindAbstractLangString(
	container aastypes.IClass,
) (result aastypes.IAbstractLangString) {
	ok := aastypes.IsAbstractLangString(container)
	if ok {
		result = container.(aastypes.IAbstractLangString)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsAbstractLangString(that)
		if abort {
			result = that.(aastypes.IAbstractLangString)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IAbstractLangString "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ILangStringNameType] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindLangStringNameType(
	container aastypes.IClass,
) (result aastypes.ILangStringNameType) {
	ok := aastypes.IsLangStringNameType(container)
	if ok {
		result = container.(aastypes.ILangStringNameType)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsLangStringNameType(that)
		if abort {
			result = that.(aastypes.ILangStringNameType)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ILangStringNameType "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ILangStringTextType] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindLangStringTextType(
	container aastypes.IClass,
) (result aastypes.ILangStringTextType) {
	ok := aastypes.IsLangStringTextType(container)
	if ok {
		result = container.(aastypes.ILangStringTextType)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsLangStringTextType(that)
		if abort {
			result = that.(aastypes.ILangStringTextType)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ILangStringTextType "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IEnvironment] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindEnvironment(
	container aastypes.IClass,
) (result aastypes.IEnvironment) {
	ok := aastypes.IsEnvironment(container)
	if ok {
		result = container.(aastypes.IEnvironment)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsEnvironment(that)
		if abort {
			result = that.(aastypes.IEnvironment)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IEnvironment "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IDataSpecificationContent] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindDataSpecificationContent(
	container aastypes.IClass,
) (result aastypes.IDataSpecificationContent) {
	ok := aastypes.IsDataSpecificationContent(container)
	if ok {
		result = container.(aastypes.IDataSpecificationContent)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsDataSpecificationContent(that)
		if abort {
			result = that.(aastypes.IDataSpecificationContent)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IDataSpecificationContent "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IEmbeddedDataSpecification] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindEmbeddedDataSpecification(
	container aastypes.IClass,
) (result aastypes.IEmbeddedDataSpecification) {
	ok := aastypes.IsEmbeddedDataSpecification(container)
	if ok {
		result = container.(aastypes.IEmbeddedDataSpecification)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsEmbeddedDataSpecification(that)
		if abort {
			result = that.(aastypes.IEmbeddedDataSpecification)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IEmbeddedDataSpecification "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ILevelType] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindLevelType(
	container aastypes.IClass,
) (result aastypes.ILevelType) {
	ok := aastypes.IsLevelType(container)
	if ok {
		result = container.(aastypes.ILevelType)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsLevelType(that)
		if abort {
			result = that.(aastypes.ILevelType)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ILevelType "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IValueReferencePair] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindValueReferencePair(
	container aastypes.IClass,
) (result aastypes.IValueReferencePair) {
	ok := aastypes.IsValueReferencePair(container)
	if ok {
		result = container.(aastypes.IValueReferencePair)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsValueReferencePair(that)
		if abort {
			result = that.(aastypes.IValueReferencePair)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IValueReferencePair "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IValueList] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindValueList(
	container aastypes.IClass,
) (result aastypes.IValueList) {
	ok := aastypes.IsValueList(container)
	if ok {
		result = container.(aastypes.IValueList)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsValueList(that)
		if abort {
			result = that.(aastypes.IValueList)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IValueList "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ILangStringPreferredNameTypeIEC61360] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindLangStringPreferredNameTypeIEC61360(
	container aastypes.IClass,
) (result aastypes.ILangStringPreferredNameTypeIEC61360) {
	ok := aastypes.IsLangStringPreferredNameTypeIEC61360(container)
	if ok {
		result = container.(aastypes.ILangStringPreferredNameTypeIEC61360)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsLangStringPreferredNameTypeIEC61360(that)
		if abort {
			result = that.(aastypes.ILangStringPreferredNameTypeIEC61360)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ILangStringPreferredNameTypeIEC61360 "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ILangStringShortNameTypeIEC61360] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindLangStringShortNameTypeIEC61360(
	container aastypes.IClass,
) (result aastypes.ILangStringShortNameTypeIEC61360) {
	ok := aastypes.IsLangStringShortNameTypeIEC61360(container)
	if ok {
		result = container.(aastypes.ILangStringShortNameTypeIEC61360)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsLangStringShortNameTypeIEC61360(that)
		if abort {
			result = that.(aastypes.ILangStringShortNameTypeIEC61360)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ILangStringShortNameTypeIEC61360 "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.ILangStringDefinitionTypeIEC61360] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindLangStringDefinitionTypeIEC61360(
	container aastypes.IClass,
) (result aastypes.ILangStringDefinitionTypeIEC61360) {
	ok := aastypes.IsLangStringDefinitionTypeIEC61360(container)
	if ok {
		result = container.(aastypes.ILangStringDefinitionTypeIEC61360)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsLangStringDefinitionTypeIEC61360(that)
		if abort {
			result = that.(aastypes.ILangStringDefinitionTypeIEC61360)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of ILangStringDefinitionTypeIEC61360 "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// Find the first instance of [aastypes.IDataSpecificationIEC61360] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func MustFindDataSpecificationIEC61360(
	container aastypes.IClass,
) (result aastypes.IDataSpecificationIEC61360) {
	ok := aastypes.IsDataSpecificationIEC61360(container)
	if ok {
		result = container.(aastypes.IDataSpecificationIEC61360)
		return
	}

	container.Descend(func(that aastypes.IClass) (abort bool) {
		abort = aastypes.IsDataSpecificationIEC61360(that)
		if abort {
			result = that.(aastypes.IDataSpecificationIEC61360)
		}
		return
	})

	if result == nil {
		panic(
			fmt.Sprintf(
				"Could not find an instance of IDataSpecificationIEC61360 "+
					"in the container of type %T: %v",
				container, container,
			),
		)
	}
	return
}

// This code has been automatically generated by:
// _dev_scripts/test_codegen/generate_finding.py
// Do NOT edit or append.
