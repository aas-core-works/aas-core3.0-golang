package aastesting

import (
	"encoding/json"
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

// Represent `that` instance as a human-readable line of an iteration trace.
func TraceMark(that aastypes.IClass) string {
	modelTypeText := aasstringification.MustModelTypeToString(
		that.ModelType(),
	)

	if aastypes.IsIdentifiable(that) {
		identifiable := that.(aastypes.IIdentifiable)
		idBytes, err := json.Marshal(identifiable.ID())
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to convert ID %s to JSON: %s",
					identifiable.ID(), err.Error(),
				),
			)
		}

		return fmt.Sprintf("%s with ID %s", modelTypeText, string(idBytes))
	}

	if aastypes.IsReferable(that) {
		referable := that.(aastypes.IReferable)
		idShortBytes, err := json.Marshal(referable.IDShort())
		if err != nil {
			panic(
				fmt.Sprintf(
					"Failed to convert ID-short %v to JSON: %s",
					referable.IDShort(), err.Error(),
				),
			)
		}

		return fmt.Sprintf("%s with ID-short %s", modelTypeText, string(idShortBytes))
	}

	return modelTypeText
}
