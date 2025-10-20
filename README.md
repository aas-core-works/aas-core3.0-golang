# aas-core3.0-golang

Manipulate, verify and de/serialize asset administration shells in Go.

[![CI](https://github.com/aas-core-works/aas-core3.0-golang/actions/workflows/ci.yml/badge.svg)](https://github.com/aas-core-works/aas-core3.0-golang/actions/workflows/ci.yml)

[![CI-for-dev-scripts](https://github.com/aas-core-works/aas-core3.0-golang/actions/workflows/ci-for-dev-scripts.yml/badge.svg)](https://github.com/aas-core-works/aas-core3.0-golang/actions/workflows/ci-for-dev-scripts.yml)

[![Coverage Status](https://coveralls.io/repos/github/aas-core-works/aas-core3.0-golang/badge.svg?branch=main)](https://coveralls.io/github/aas-core-works/aas-core3.0-golang?branch=main)

This is a software development kit (SDK) to:

* manipulate,
* verify, and
* de/serialize to and from JSON

… Asset Administration Shells based on the version 3.0 of the meta-model.

For a brief introduction, see [Getting Started].

For a detailed documentation of the API, see [API Documentation].

We documented most of the rationale behind the implementation and interface choices in the section [Design Decisions]. 

If you want to contribute, see our [Contributing Guide].

The history of the module is listed in the [Change Log].

[Getting Started]: #getting-started
[API documentation]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang
[Design Decisions]: #design-decisions
[Contributing Guide]: #contributing-guide
[Change Log]: #change-log

## Getting Started

Here's a quick intro to get you started with the SDK.
See how you can:

* [Install the SDK](#install-the-sdk),
* [Programmatically create, get and set properties of an AAS model](#create-get-and-set),
* [Switch on runtime types of instances](#switch-on-runtime-types),
* [Iterate over, copy and transform a model](#iterate),
* [Verify a model](#verify),
* [De/serialize a model from and to JSON](#json-deserialization), and
* [De/serialize a model from and to XML](#xml-deserialization), and
* [Enhance instances with your custom data](#enhancing). 

### Install the SDK

The SDK is available as a module `github.com/aas-core-works/aas-core3.0-golang`.

Install it using `go get`:

```
go get github.com/aas-core-works/aas-core3.0-golang
```

### Create, Get and Set

The package [`types`] defines all the data types of the meta-model.
This includes structs, interfaces and enumerations.

[`types`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types

#### Creation

We model each meta-model class, abstract and concrete alike, as Go interface.
You should prefer interfaces to structs so that you can use [`enhancing`] package (see below in Section [Enhancing](#enhancing)).

The most general interface [`types.IClass`] represents an instance of the AAS model.
All other interfaces adopt it.

We use constructors to create an AAS model.
They are marked as `New*`.
For example, [`types.NewEnvironment`].
Usually you start bottom-up, all the way up to the [`types.Environment`].

[`types.IClass`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#IClass
[`types.NewEnvironment`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#NewEnvironment
[`types.Environment`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#Environment

#### Getting and Setting Properties

All properties of the classes are modeled as getter and setter methods.
The properties which are not set should be assigned a `nil`. 

The lists are modeled as slices.
For example, [`types.Environment.Submodels`]:

```go
func (e *Environment) Submodels() []ISubmodel
```

[`types.Environment.Submodels`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#Environment.Submodels

Byte arrays are modeled as slices of `byte`.
For example, [`types.Blob.Value`]:

```go
func (b *Blob) Value() []byte
```

[`types.Blob.Value`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#Blob.Value

#### Getters with a Default Value

For optional properties which come with a default value, we provide special getters, `{property name}OrDefault`.
If the property is `nil`, this getter will give you the default value.
Otherwise, if the property is set, the actual value of the property will be returned.

For example, see [`types.IHasKind.KindOrDefault`].

[`types.IHasKind.kindOrDefault`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#IHasKind

#### Example: Create an Environment with a Submodel

Here is a very rudimentary example where we show how to create an environment which contains a submodel.

The submodel will contain two elements, a property and a blob.

```go
package main

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

// Create a new instance of the `value` and return the pointer to it.
func NewString(value string) *string {
	return &value
}

func main() {
	// Create the first element
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	someElement.SetIDShort(
		NewString("someProperty"),
	)
	someElement.SetValue(
		NewString("some-value"),
	)

	// Create the second element
	anotherElement := aastypes.NewBlob(
		"application/octet-stream",
	)
	anotherElement.SetIDShort(
		NewString("someBlob"),
	)
	anotherElement.SetValue(
		[]byte{0xDE, 0xAD, 0xBE, 0xEF},
	)

	// Nest the elements in a submodel
	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
			anotherElement,
		},
	)

	// Now create the environment to wrap it all up
	environment := aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// You can set the properties.
	environment.Submodels()[0].SubmodelElements()[0].(aastypes.IProperty).SetValue(
		NewString("changed-value"),
	)

	// You can access the properties from the children.
	fmt.Printf(
		"%v\n",
		*environment.Submodels()[0].SubmodelElements()[0].(aastypes.IProperty).Value(),
	)
	// Output:
	// changed-value
}
``` 

(See: [Example CreateAnEnvironmentWithASubmodel](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-CreateAnEnvironmentWithASubmodel))

### Switch on Runtime Types

As we noted in Section [Creation](#creation), the classes of the meta-model are specified as Go interfaces.

Go uses [structural typing] so any struct satisfying an interface automatically implements that interface.
This breaks runtime [type switches], as you can not exactly infer the exact runtime type of instance as soon as it satisfies multiple interfaces.
To that end, every instance is provided with `ModelType()` method that provides the exact model type at runtime as [`types.ModelType`].

[structural typing]: https://en.wikipedia.org/wiki/Structural_type_system
[type switches]: https://go.dev/tour/methods/16
[`types.ModelType`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#ModelType

This has an additional benefit for computational efficiency.
Applying type switches is more complex than applying a switch on an integer enumeration.
A switch on an enumerator can be optimized by compiler as a [jump table], while type switches are not as straight-forward, especially due to multiple inheritance. 

[jump table]: https://en.wikipedia.org/wiki/Branch_table

We also provide functions `Is*` in [`types`] package to allow for runtime type checks which honour the inheritance.
For example, see [`types.IsSubmodelElement`].

[`types.IsSubmodelElement`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#IsSubmodelElement
[`types.IsProperty`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#IsProperty
[`types.IsBlob`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#IsBlob


Here is a short example with [`types.IsSubmodelElement`], [`types.IsProperty`] and [`types.IsBlob`]:

```go
package main

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

func main() {
	// Create the first element
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDInt,
	)

	fmt.Printf("%v\n", aastypes.IsSubmodelElement(someElement))
	fmt.Printf("%v\n", aastypes.IsProperty(someElement))
	fmt.Printf("%v\n", aastypes.IsBlob(someElement))
	// Output:
    // true
    // true
    // false
}
```

(See: [Example IsXxx](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-IsXxx))

### Iterate

Looping through the instances of a model is tedious to write manually, especially when you want to recursively iterate over a model.
The SDK provides two methods for all the structs implementing [`types.IClass`], [`DescendOnce` and `Descend`], which you can use to loop through the instances.

[`DescendOnce` and `Descend`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#IClass

Both `DescendOnce` and `Descend` iterate over referenced children of an instance of [`types.IClass`].
The method [`DescendOnce`], as it names suggests, stops after all the immediate children has been iterated over.
The method [`Descend`] continues recursively to grand-children, grand-grand-children *etc.*

You have to supply a callback function which is applied on every instance that we iterate over.
If the callback function returns `true` (as its return argument `abort`), the iteration stops.

Here is a short example which shows how you can get all the properties from an environment whose ID-short contains the word `another`:

```go
package main

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	"strings"
)

// Create a new instance of the `value` and return the pointer to it.
func NewString(value string) *string {
	return &value
}

func main() {
	// Prepare the environment
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	someElement.SetIDShort(
		NewString("someProperty"),
	)

	anotherElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	anotherElement.SetIDShort(
		NewString("anotherProperty"),
	)

	yetAnotherElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	yetAnotherElement.SetIDShort(
		NewString("yetAnotherProperty"),
	)

	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
			anotherElement,
			yetAnotherElement,
		},
	)

	environment := aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Iterate using ``Descend``
	environment.Descend(
		func(that aastypes.IClass) (abort bool) {
			if aastypes.IsProperty(that) {
				idShort := that.(aastypes.IProperty).IDShort()
				if idShort != nil &&
					strings.Contains(strings.ToLower(*idShort), "another") {
					fmt.Printf("%s\n", *idShort)
				}
			}
			return
		},
	)
}
```

(See: [Example IterateOverEnvironment](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-IterateOverEnvironment))

#### Iterate over Enumeration Literals

Go does not treat enumerations as collections, but as a fixed tuple of constants.
This means that there is no "native" way to iterate over enumerations using for-range-loop.

If the constants are consecutive numbers (using [`iota`]), you can loop through a segment of literals by using for-loop with an incrementing integer.
The enumerations in our SDK are indeed always defined as integers using [`iota`] (see Section [design decisions]), so that is one option.

[`iota`]: https://go.dev/ref/spec#Iota
[design-decisions]: #design-decisions

Such an incrementing for-loop can be confusing for the reader, and potentially buggy.
For example, if the order of your start and end literal in the enumeration ever change in the future, you end up with a non-loop, which is clearly a bug.

To avoid these bugs, we provide `LiteralsOf*` slices that you can readily use in your code.
Though Go does not provide a concept of immutability, these slices are meant to be constant.
You should not change them in your code, only read them.

If you want to obtain the string representation of the literal, we provide the [`stringification`] package.
The functions `stringification.{enumeration name}ToString` give you back the string representation of the literal, and an `ok` which is set to `false` if the literal was an invalid number.
For the client's convenience, our SDK also implements the functions `stringification.Must{enumeration name}ToString` which returns the string representation, or panics.
If you are certain that your code deals with only correct literals, `Stringification.must{enumeration name}ToString` will spare you an is-ok check.
For example, see [`stringification.ModellingKindToString`] and [`stringification.MustModellingKindToString`].

[`stringification.ModellingKindToString`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/stringification#ModellingKindToString
[`stringification.MustModellingKindToString`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/stringification#MustModellingKindToString

Here is a short example that illustrates how to loop over enumeration literals of the enumeration [`types.ModelingKind`] using the slice [`types.LiteralsOfModellingKind`]:

[`types.ModelingKind`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/stringification#ModellingKindToString
[`types.LiteralsOfModellingKind`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#pkg-variables

```go
package main

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	"strings"
)

// Create a new instance of the `value` and return the pointer to it.
func NewString(value string) *string {
	return &value
}

func main() {
	for _, literal := range aastypes.LiteralsOfModellingKind {
		fmt.Printf(
			"Literal as number: %d, literal as string: %s\n",
			literal, aasstringification.MustModellingKindToString(literal),
		)
	}
	// Output:
	// Literal as number: 0, literal as string: Template
	// Literal as number: 1, literal as string: Instance
}
```

(See: [Example IterateOverEnumerationLiterals](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-IterateOverEnumerationLiterals))

### Verify

Our SDK allows you to verify that a model satisfies the constraints of the meta-model.

The verification logic is concentrated in the package [`verification`], and all it takes is a call to [`verification.Verify`] function.
The function [`verification.Verify`] will check that constraints in the given model element are satisfied, including the recursion into children elements.

You have to pass in a callback function to [`verification.Verify`] which is applied on each reported error.
The verification stops if the callback function ever returns `true` (as its only return variable `abort`).
This is useful, for example, if you want to report only a certain number of errors.

[`verification`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/verification
[`verification.Verify`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/verification#Verify

Here is a short example snippet:

```go
package main

import (
	"fmt"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasverification "github.com/aas-core-works/aas-core3.0-golang/verification"
)

// Create a new instance of the `value` and return the pointer to it.
func NewString(value string) *string {
	return &value
}

func main() {
	// Prepare the environment
	someElement := aastypes.NewProperty(
		aastypes.DataTypeDefXSDString,
	)
	// The ID-shorts must be proper variable names,
	// but there is a dash (`-`) in this ID-short.
	someElement.SetIDShort(
		NewString("some-property"),
	)

	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
		},
	)

	environment := aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Verify
	aasverification.Verify(
		environment,
		func(err *aasverification.VerificationError) (abort bool) {
			fmt.Printf("%s\n", err.Error())
			return
		},
	)
	// Output:
	// Submodels[0].SubmodelElements[0].IDShort: ID-short of Referables
	// shall only feature letters, digits, underscore (``_``); starting mandatory
	// with a letter. *I.e.* ``[a-zA-Z][a-zA-Z0-9_]*``.
}
```

(See: [Example Verification](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-Verification))

#### Omitted Constraints

Not all constraints specified in the meta-model can be verified.
Some constraints require external dependencies such as an AAS registry.
Verifying the constraints with external dependencies is out-of-scope of our SDK, as we still lack standardized interfaces to those dependencies.

However, all the constraints which need *no* external dependency are verified.
For a full list of exception, please see the description of the package [`types`].

### JSON de/serialization

Our SDK handles the de/serialization of the AAS models from and to JSON format through the package [`jsonization`].

[`jsonization`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/jsonization

Instead of de/serializing to and from strings or arrays of bytes, we de/serialize from JSON-able structures such as `map[string]interface{}` or `[]map[string]interface{}`.
This allows the de/serialization to be more versatile so you are not restricted to JSON, but you can also use JSON-like formats such as [YAML] or [binary JSON].

[YAML]: https://en.wikipedia.org/wiki/YAML
[binary JSON]: https://en.wikipedia.org/wiki/BSON

#### Serialize to JSON

To serialize, you call the function [`jsonization.ToJsonable`] on an instance of [`types.IClass`] which will convert it to a JSON-able `map[string]interface{}`.

[`jsonization.ToJsonable`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/jsonization#ToJsonable

Here is a snippet that converts the environment first into a JSON-able object, and next converts the JSON-able object to text:

```go
package main

import (
	"encoding/json"
	"fmt"
	aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

func main() {
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

	environment := aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Serialize to jsonable
	var jsonable map[string]interface{}
	var err error
	jsonable, err = aasjsonization.ToJsonable(environment)
	if err != nil {
		panic(err.Error())
	}

	// Serialize jsonable to string
	var bb []byte
	bb, err = json.MarshalIndent(jsonable, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	text := string(bb)

	fmt.Println(text)
	// Output:
	// {
	//   "submodels": [
	//     {
	//       "id": "some-unique-global-identifier",
	//       "modelType": "Submodel",
	//       "submodelElements": [
	//         {
	//           "idShort": "someProperty",
	//           "modelType": "Property",
	//           "value": "some-value",
	//           "valueType": "xs:string"
	//         }
	//       ]
	//     }
	//   ]
	// }
}
```

(See: [Example JsonizationTo](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-JsonizationTo))

#### De-serialize from JSON

Our SDK can convert a JSON-able structure back to an instance of [`types.IClass`]. 
To that end, you call the appropriate function `jsonization.{class name}FromJsonable`.
For example, if you want to de-serialize an instance of [`types.IEnvironment`], call [`jsonization.EnvironmentFromJsonable`].

[`types.IEnvironment`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#IEnvironment
[`jsonization.EnvironmentFromJsonable`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/jsonization#EnvironmentFromJsonable

Note that the SDK cannot de-serialize classes automatically as the discriminator property `modelType` is not included in the serializations for *all* the classes.
Without the discriminator property provided, we thus cannot know the actual type of the instance just from the serialization.
See [this sections on discriminators in AAS Specs] for more details. 

[this sections on discriminators in AAS Specs]: https://github.com/admin-shell-io/aas-specs/tree/master/schemas/json#discriminator

Here is an example snippet to show you how to de-serialize an instance of [`types.IEnvironment`]:

```go
package main

import (
	"encoding/json"
	"fmt"
	aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

func main() {
	text := `
{
  "submodels": [
    {
      "id": "some-unique-global-identifier",
      "modelType": "Submodel",
      "submodelElements": [
        {
          "idShort": "someProperty",
          "modelType": "Property",
          "value": "some-value",
          "valueType": "xs:string"
        }
      ]
    }
  ]
}`

	bb := []byte(text)

	var jsonable map[string]interface{}
	var err error
	err = json.Unmarshal(bb, &jsonable)
	if err != nil {
		panic(err.Error())
	}

	var environment aastypes.IEnvironment
	environment, err = aasjsonization.EnvironmentFromJsonable(
		jsonable,
	)
	if err != nil {
		panic(err.Error())
	}

	environment.Descend(
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
```

(See: [Example JsonizationFrom](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-JsonizationFrom))

### XML de/serialization

The de/serialization of the AAS models is handled by the package [`xmlization`].

[`xmlization`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/xmlization

#### Serialize to XML

The XML serialization lives in [`xmlization`] package.

[`xmlization`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/xmlization

We serialize the instances by writing tokens to [`xml.Encoder`].
While we could immediately return a string or write to a [`io.Writer`], writing tokens to [`xml.Encoder`] allows you to better steer the format of the output.
For example, you can adjust the indention by calling [`xml.Encoder.Indent`].

[`xml.Encoder`]: https://pkg.go.dev/encoding/xml#Encoder
[`io.Writer`]: https://pkg.go.dev/io#Writer
[`xml.Encoder.Indent`]: https://pkg.go.dev/encoding/xml#Encoder.Indent

Given an encoder, call the function [`xmlization.Marshal`] on it together with your instance of [`types.IClass`] that you want to serialize.

Here is a snippet that serializes an environment to XML:

```go
package main

import (
	"encoding/xml"
	"fmt"

	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"

	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
)

func main() {
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

	environment := aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Serialize to XML
	builder := new(strings.Builder)
	encoder := xml.NewEncoder(builder)
	encoder.Indent("", "  ")

	// We want to include the namespace in the root XML element.
	withNamespace := true

	var err error
	err = aasxmlization.Marshal(encoder, environment, withNamespace)
	if err != nil {
		panic(err.Error())
	}

	text := builder.String()

	fmt.Println(text)
	// Output:
	// <environment xmlns="https://admin-shell.io/aas/3/0">
	//   <submodels>
	//     <submodel>
	//       <id>some-unique-global-identifier</id>
	//       <submodelElements>
	//         <property>
	//           <idShort>someProperty</idShort>
	//           <valueType>xs:string</valueType>
	//           <value>some-value</value>
	//         </property>
	//       </submodelElements>
	//     </submodel>
	//   </submodels>
	// </environment>
}
```

(See: [Example XmlizationTo](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-XmlizationTo))

#### De-serialize from XML
 
For efficient one-pass de-serialization, we directly read tokens from [`xml.Decoder`].
The function [`xmlization.Unmarshal`] de-serializes an instance from the given XML decoder. 

[`xml.Decoder`]: https://pkg.go.dev/encoding/xml#Decoder
[`xmlization.Unmarshal`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/xmlization#Unmarshal

Here is an example snippet to show you how to de-serialize an instance of [`types.IEnvironment`]:

```go
package main

import (
	"encoding/xml"
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
)

func main() {
	text := `<environment xmlns="https://admin-shell.io/aas/3/0">
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
	decoder := xml.NewDecoder(reader)

	var instance aastypes.IClass
	var err error
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
```

(See: [Example XmlizationFrom](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-XmlizationFrom))

##### Versatility to Different Sources

We do not assume to know the source of the XML, and choose to be versatile with handling different sources (wire, file, element embedded in a larger XML document *etc.*).
To that end, we expect that [`xml.Decoder`] already points to the root XML element that you want to de-serialize from.

For example, if you are decoding from a file that starts with a [Byte Order Mark (BOM)] and an XML declaration, you first need to read those yourself.
Here is an example:

[Byte Order Mark (BOM)]: https://en.wikipedia.org/wiki/Byte_order_mark

```go
package main

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

func main() {
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
```

(See: [Example XmlizationFromTextWithBOMAndXMLDeclaration](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-XmlizationFromTextWithBOMAndXMLDeclaration))

##### No Attributes Expected

The XML specification of the AAS meta-model expects no attributes in the XML elements.
Consequently, we follow the specification and throw an error if there are any XML attributes present in a start element.
There is one exception, so we do allow for `xmlns` and `xmlns:*` attributes as they are necessary for the correct namespacing.

If there are attributes in your XML documents, make sure you wrap the [`xml.TokenReader`] to undo them on the fly.
Here is an example snippet:

```go
package main

import (
	"encoding/xml"
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
)

// TokenReaderNoAttributes reads tokens from a reader and
// removes any attributes in the start elements.
type TokenReaderNoAttributes struct {
	r xml.TokenReader
}

func NewTokenReaderNoAttributes(r xml.TokenReader) *TokenReaderNoAttributes {
	return &TokenReaderNoAttributes{r: r}
}

func (trna *TokenReaderNoAttributes) Token() (xml.Token, error) {
	var token xml.Token
	var err error

	for {
		token, err = trna.r.Token()
		if err != nil {
			return token, err
		}

		var startElement xml.StartElement
		startElement, isStartElement := token.(xml.StartElement)

		if !isStartElement {
			return token, err
		}

		startElement.Attr = []xml.Attr{}
		return startElement, err
	}
}

func main() {
	text := `<environment
	xmlns="https://admin-shell.io/aas/3/0"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="https://admin-shell.io/aas/3/0 AAS.xsd"
>
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

	// Use a decoder which skips the XML declarations
	decoder := xml.NewTokenDecoder(
		NewTokenReaderNoAttributes(
			xml.NewDecoder(reader),
		),
	)

	var instance aastypes.IClass
	var err error
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
```

(See: [Example XmlizationSkipAttributes](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-XmlizationSkipAttributes))

##### Namespace Aliases

We rely on [`xml.Decoder`] to handle namespacing of the elements.

Here is an example snippet where we use a namespace alias:

```go
package main

import (
	"encoding/xml"
	"fmt"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
	aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
	"strings"
)

func main() {
	text := `<aas:environment
	xmlns:aas="https://admin-shell.io/aas/3/0"
>
  <aas:submodels>
    <aas:submodel>
      <aas:id>some-unique-global-identifier</aas:id>
      <aas:submodelElements>
        <aas:property>
          <aas:idShort>someProperty</aas:idShort>
          <aas:valueType>xs:string</aas:valueType>
          <aas:value>some-value</aas:value>
        </aas:property>
      </aas:submodelElements>
    </aas:submodel>
  </aas:submodels>
</aas:environment>`

	reader := strings.NewReader(text)

	decoder := xml.NewDecoder(reader)

	var instance aastypes.IClass
	var err error
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
```

(See: [Example XmlizationNamespaceAlias](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-XmlizationNamespaceAlias))

### Enhancing

In any complex application, creating, modifying and de/serializing AAS instances is not enough.
You have to insert your custom application-specific data to the model in order for the model to be useful.

Take, for example, parent-child relationship.
The current library ignores it, and there is no easy way for you to find out to which [`types.ISubmodel`] a particular [`types.ISubmodelElement`] belongs to.

[`types.ISubmodel`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#ISubmodel
[`types.ISubmodelElement`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/types#ISubmodelElement

We did want to keep the types as simple as possible — the parent-child relationships can get tricky very soon if you have multiple environments with shared submodels *etc.*
Instead of overcomplicating the code and making it barely traceable, we decided to keep it simple and frugal in features.

However, that is little solace if you are developing an GUI editor where you know for sure that there will be only one environment, and where parent-child relationships are crucial for so many tasks.
What is more, parent-child relationships are not the only data that need to be intertwined — you probably want history, localized caches *etc.*

#### Hashtable?

There are different ways how application-specific data can be synced with the model.
One popular technique is to use [Hashtable]'s and simply map model instances to your custom nuggets of data.
This works well if the data is read-only, and you can spare the cycles for the lookups (which is often acceptable as they run on average in time complexity `O(1)` anyhow).

[Hashtable]: https://learn.microsoft.com/en-US/dotnet/api/system.collections.hashtable

Otherwise, if you need to modify the data, maintaining the consistency between the [Hashtable] and your nuggets becomes difficult.
For example, if you forget to remove the entries from the [Hashtable] when you remove the instances from the model, you might clog your garbage collector.

#### Wrapping

Hence, if you modify the data, you need to keep it close to the model instance.
In dynamic languages, such as Python and JavaScript, you can simply add your custom fields to the object.
This does not work in such a static language like Go.

One solution, usually called [Decorator pattern], is to *wrap* or *decorate* the instances with your application-specific data.
The decorated objects should satisfy both the interface of the original model and provide a way to retrieve your custom nuggets of information.

[Decorator pattern]: https://en.wikipedia.org/wiki/Decorator_pattern

Writing wrappers for many classes in the AAS meta-model is a tedious task.
We therefore pre-generated the most of the boilerplate code in the package [`enhancing`].

[`enhancing`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/enhancing

In the context of decoration, we call your specific data *enhancements*.
First, you need to specify how individual instances are enhanced, *i.e.* how to produce enhancements for each one of them.
We call this an *enhancement factory*.
Second, you need to recursively wrap your instances with the given enhancement factory.

The enhancing is generic and can work with any form of enhancement classes.
You need to specify your enhancement factory as a function which takes an instance of [`types.IClass`] as input and returns either an enhancement, or `false` as `shouldEnhance` return value if you do not want to enhance the particular instance.

The methods [`enhancing.Wrap`] and [`enhancing.Unwrap`] perform the wrapping and unwrapping, respectively.
The method [`enhancing.MustUnwrap`] is a shortcut method that spares you to write a non-nil check of [`enhancing.Unwrap`] and the related panic if the instance has not been wrapped.

[`enhancing.Wrap`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/enhancing#Wrap
[`enhancing.Unwrap`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/enhancing#Unwrap
[`enhancing.MustUnwrap`]: https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/enhancing#MustUnwrap

#### Example: Parent-Child Enhancement

Let us now consider the aforementioned example.
We want to keep track of parent-child relationships in a model.

The following code snippets first constructs an environment for illustration.
Then we specify the enhancement such that each instance is initialized with the parent set to `nil`.
Finally, we modify the enhancements such that they reflect the parent-child relationships.

```go
package main

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

func main() {
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
```

(See: [Example EnhancingParentChild](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-EnhancingParentChild))

Note that this approach is indeed more maintainable than the one with [Hashtable], but you still need to take extra care.
If you create new submodels and insert them into the environment, you have to make sure that you wrap them appropriately.
If you move a submodel from one environment to another, you have to update the parent link manually *etc.*

#### Example: Selective Enhancement

We demonstrate now how you can selectively enhance only some instances in an [`types.IEnvironment`].

For example, let us assign a unique identifier to all instances which are referable.
All the other instances are not enhanced.

```go
package main

import (
	"fmt"
	aasenhancing "github.com/aas-core-works/aas-core3.0-golang/enhancing"
	aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
	aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)

type ReferableEnhancement struct {
	ID int
}

func main() {
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

	administrativeInfo := aastypes.NewAdministrativeInformation()
	administrativeInfo.SetVersion(
		NewString("1.0"),
	)

	submodel := aastypes.NewSubmodel(
		"some-unique-global-identifier",
	)
	submodel.SetSubmodelElements(
		[]aastypes.ISubmodelElement{
			someElement,
		},
	)
	submodel.SetAdministration(administrativeInfo)

	var environment aastypes.IEnvironment
	environment = aastypes.NewEnvironment()
	environment.SetSubmodels(
		[]aastypes.ISubmodel{
			submodel,
		},
	)

	// Wrap everything
	nextID := 0

	factory := func(that aastypes.IClass) (enh *ReferableEnhancement, shouldEnh bool) {
		if aastypes.IsReferable(that) {
			enh = &ReferableEnhancement{ID: nextID}
			shouldEnh = true

			nextID++
		}

		return
	}

	environment = aasenhancing.Wrap[*ReferableEnhancement](
		environment,
		factory,
	).(aastypes.IEnvironment)

	environment.Descend(
		func(that aastypes.IClass) (abort bool) {
			enh, ok := aasenhancing.Unwrap[*ReferableEnhancement](that)
			if ok {
				fmt.Printf(
					"%s ID: %d\n",
					aasstringification.MustModelTypeToString(that.ModelType()),
					enh.ID,
				)
			} else {
				fmt.Printf(
					"%s No ID\n",
					aasstringification.MustModelTypeToString(that.ModelType()),
				)
			}
			return
		},
	)
    // Output:
	// Submodel ID: 0
	// AdministrativeInformation No ID
	// Property ID: 1
}
```

(See: [Example EnhancingReferable](https://pkg.go.dev/github.com/aas-core-works/aas-core3.0-golang/getting_started#example-package-EnhancingReferable))

#### No Re-wraps Allowed

We panic on re-wraps of already wrapped instances to avoid costly iterations over the object trees.
Additionally, we want to prevent bugs in many settings where the enhancement factory assigns unique identifiers to instances or performs non-idempotent operations.

Please let us know by [creating an issue] if you need re-wraps to be allowed, and please tell us more about your particular scenario.

[create an issue]: https://github.com/aas-core-works/aas-core3.0-golang/issues/new

## API

For a detailed documentation of the API, see [API documentation].

## Design Decisions

We present here some of the choices we made during the design and implementation of the SDK.
While it is not necessary to understand our thread of thought to *use* the SDK, we explain the rationale here behind why we structured and programmed the SDK the way we did.
This should hopefully clear up some confusion, or ease the frustration, if you prefer certain features to be implemented differently.

### Enumeration Literals as Numbers

We optimize the enumerations for look-ups and comparisons instead of string representation.
Thus, we implement literals as numbers (instead of strings).
For example, this makes lookups faster as hash values are directly computed on a numeric literal involving usually only a few arithmetic operations.
In contrast, if the enumeration literals were listed as strings, the hash value of the literal would need to be computed by iterating through *all the characters* of the string.

### Inheritance Hierarchy

The AAS meta-model uses multiple inheritance.
However, Go supports no inheritance.

Instead of multiple inheritance we use interfaces and provide `Is*` functions to dynamically decide the instance type at runtime.
All the interfaces inherit from the most general interface [`types.IClass`]. 
Please see Section [Switch on Runtime Types] how you can determine the model type at runtime.

[Switch on Runtime Types]: #switch-on-runtime-types

### Interface Names

It is common in Go to call the interfaces with an "-er" suffix (`Reader`, `Writer`, *etc.*), see the book ["Effective Go"].
This works well when you write code by hand, and can be creative.
In our setting where the code is generated mostly automatically, we could not easily avoid naming conflicts if we added the suffix "-er" indiscriminately.
Therefore, we opted to call all the interface with the prefix "I-" (`IClass`, `IEnvironment` *etc.*).

["Effective Go"]: https://go.dev/doc/effective_go#interface-names

## Contributing Guide

### Issues

Please report bugs or feature requests by [creating GitHub issues].

[creating GitHub issues]: https://github.com/aas-core-works/aas-core3.0-golang/issues/new/choose

### In Code

If you want to contribute in code, pull requests are welcome!

Please do [create a new issue] before you dive into coding.
It can well be that we already started working on the feature, or that there are upstream or downstream complexities involved which you might not be aware of.

[create a new issue]: https://github.com/aas-core-works/aas-core3.0-golang/issues/new/choose

### SDK Code Generation

The biggest part of the code has been automatically generated by [aas-core-codegen].
It probably makes most sense to change the generator rather than add new functionality.
However, this needs to be decided on a case-by-case basis.

[aas-core-codegen]: https://github.com/aas-core-works/aas-core-codegen

### Test Code Generation

The code of the unit tests has been automatically generated using the Python scripts in the `_dev_scripts/test_codegen/` directory.

To re-generate the test code, first create a virtual environment at the root of the repository:

```
python -m venv venv
```

Activate the virtual environment (in Windows):

```
venv\Scripts\activate
```

or in Linux:
```
source venv/bin/activate
```

Then install the dependencies:

```
pip3 install -e . _dev_scripts
```

Now you can run the generation scripts:

```
python _dev_scripts/test_codegen/generate_all.py
```

### Test Data

The test data is automatically generated by [aas-core3.0-testgen], and copied to this repository on every change.

[aas-core3.0-testgen]: https://github.com/aas-core-works/aas-core3.0-testgen

### Pull Requests

**Feature branches**.
We develop using the feature branches, see [this section of the Git book].

[this section of the Git book]: https://git-scm.com/book/en/v2/Git-Branching-Branching-Workflows 

If you are a member of the development team, create a feature branch directly within the repository.

Otherwise, if you are a non-member contributor, fork the repository and create the feature branch in your forked repository. See [this GitHub tuturial] for more guidance. 

[this GitHub tutorial]: https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request-from-a-fork

**Branch Prefix**.
Please prefix the branch with your Github user name (*e.g.,* `mristin/Add-some-feature`).

**Continuous Integration**. 
GitHub will run the continuous integration (CI) automatically through GitHub actions.
The CI includes checking the formatting, vetting the code, running the tests, *etc.*

### Commit Messages

The commit messages follow the guidelines from https://chris.beams.io/posts/git-commit:

* Separate subject from body with a blank line,
* Limit the subject line to 50 characters,
* Capitalize the subject line,
* Do not end the subject line with a period,
* Use the imperative mood in the subject line,
* Wrap the body at 72 characters, and
* Use the body to explain *what* and *why* (instead of *how*).

## Change Log

### v1.0.6 (2025-10-20)

This is a patch release that propagates a fix for references index constraint
where indices in references were erroneously assumed to be positive integers.

### v1.0.5 (2025-05-14)

We propagate the changes and fixes for V3.0.2; please refer to:
* https://github.com/aas-core-works/aas-core-meta/pull/335
* https://github.com/aas-core-works/aas-core-meta/pull/341
* https://github.com/aas-core-works/aas-core-meta/pull/343
* https://github.com/aas-core-works/aas-core-meta/pull/353
* https://github.com/aas-core-works/aas-core-meta/pull/365
* https://github.com/aas-core-works/aas-core-meta/pull/368


### v1.0.4 (2024-06-19)

This is a patch release concerning the fixes for the bugs revealed in
the field by the pioneer users.

* Ignore attributes prefixed with `xmlns` (#27)
* Fix XML deserialization for over-consumption (#29)
* Return only abstract errors (#30)

### v1.0.3 (2024-04-16)

The `dataSpecification` field in `EmbeddedDataSpecification` is made
optional, according to the book.

### v1.0.2 (2024-03-23)

* Update to aas-core-meta, codegen, testgen cb28d18, c414f32, 6ff39c260 (#17)

  We propagate the fix from abnf-to-regex related to maximum qualifiers 
  which had been mistakenly represented as exact repetition before.

### v1.0.1 (2024-03-13)

* Update to aas-core-meta, codegen, testgen 79314c6, 94399e1, e1087880 (#15)

  This patch release brings about the fix for patterns concerning dates and
  date-times with zone offset `14:00` which previously allowed for
  a concatenation without a plus sign.

### v1.0.0 (2024-02-02)

This is the first stable release. The release candidates stood
the test of time, so we are now confident to publish a stable
version.

### v1.0.0-rc3 (2023-09-08)

* Update to aas-core-meta, codegen, testgen  4d7e59e, 7e264a0 and
  9b43de2e (#11)

  Notably, this fixes constraints AASd-131 and AASc-3a-010,
  propagating the changes from aas-core-meta.

### v1.0.0-rc2 (2023-06-28)

* Update to aas-core-meta, codegen, testgen 44756fb, 607f65c,
  bf3720d7 (#9)

  This is an important patch propagating pull request 275 in
  aas-core-meta which affected the constraints and their documentation.

### v1.0.0-rc1 (2023-06-23)

* This is the initial version.