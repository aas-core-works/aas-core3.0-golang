package getting_started_test

import (
	"fmt"
	aasenhancing "github.com/aas-core-works/aas-core3.0-golang/enhancing"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

type ParentEnhancement struct {
	Parent aastypes.IClass
}

type stack []aastypes.IClass

func (s *stack) Push(v aastypes.IClass) {
	*s = append(*s, v)
}

func (s *stack) Pop() aastypes.IClass {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func Example_enhancingParentChild() {
	// Create a new instance of the `value` and return the pointer to it.
	NewString := func(value string) *string {
		return &value
	}

	// Prepare the environment
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	someElement.SetIDShort(
		NewString("someProperty"),
	)
	someElement.SetValue(
		NewString("some-value"),
	)

	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
		},
	)

	var environment aastypes.IEnvironment
	environment = aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Wrap everything
	factory := func(that aastypes.IClass) (enh *ParentEnhancement, shouldEnh bool) {
		enh = &ParentEnhancement{}
		shouldEnh = true
		return
	}

	environment = aasenhancing.Wrap[*ParentEnhancement](
		environment,
		factory,
	).(aastypes.IEnvironment)

	// Initialize the parents
	var s stack
	s.Push(environment)
	for len(s) > 0 {
		instance := s.Pop()
		instance.DescendOnce(
			func(child aastypes.IClass) (abort bool) {
				enh := aasenhancing.MustUnwrap[*ParentEnhancement](child)
				enh.Parent = instance

				s.Push(child)
				return
			},
		)
	}

	// Retrieve the parent of the first submodel
	parent := aasenhancing.MustUnwrap[*ParentEnhancement](
		environment.Submodels()[0],
	).Parent

	fmt.Printf("%v\n", parent == environment)
	// Output:
	// true
}
