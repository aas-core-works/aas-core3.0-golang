"""Generate the test code for the jsonization of classes with descendants."""

import io
import os
import pathlib
import sys

import aas_core_codegen
import aas_core_codegen.common
import aas_core_codegen.naming
import aas_core_codegen.parse
import aas_core_codegen.run
from aas_core_codegen import intermediate
from aas_core_codegen.common import Stripped, Identifier
from aas_core_codegen.golang import naming as golang_naming
from aas_core_codegen.golang.common import (
    INDENT as I,
    INDENT2 as II,
    INDENT3 as III,
)

import test_codegen.common


def main() -> int:
    """Execute the main routine."""
    symbol_table = test_codegen.common.load_symbol_table()

    this_path = pathlib.Path(os.path.realpath(__file__))
    repo_root = this_path.parent.parent.parent

    warning = test_codegen.common.generate_warning_comment(
        this_path.relative_to(repo_root)
    )

    blocks = [
        Stripped("package jsonization_test"),
        warning,
        Stripped(
            f"""\
import (
{I}"testing"
{I}aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
{I}aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
)"""
        ),
    ]

    for our_type in symbol_table.our_types:
        if not isinstance(
            our_type, (intermediate.AbstractClass, intermediate.ConcreteClass)
        ):
            continue

        # NOTE (mristin, 2023-06-08):
        # We can only de-serialize instances of concrete descendants which carry
        # a model type in their serializations.
        concrete_descendants_with_model_type = [
            concrete_descendant
            for concrete_descendant in our_type.concrete_descendants
            if concrete_descendant.serialization.with_model_type
        ]

        if len(concrete_descendants_with_model_type) == 0:
            continue

        descendant_cls = concrete_descendants_with_model_type[0]

        must_load_minimal_name = golang_naming.function_name(
            Identifier(f"must_load_minimal_{descendant_cls.name}")
        )

        deserialization_function = golang_naming.function_name(
            Identifier(f"{our_type.name}_from_jsonable")
        )

        test_name = golang_naming.function_name(
            (Identifier(f"test_{our_type.name}_round_trip_OK_over_descendant"))
        )

        model_type = aas_core_codegen.naming.json_model_type(descendant_cls.name)

        blocks.append(
            Stripped(
                f"""\
func {test_name}(t *testing.T) {{
{I}instance := aastesting.{must_load_minimal_name}()

{I}jsonable, err := aasjsonization.ToJsonable(instance)
{I}if err != nil {{
{II}t.Fatalf(
{III}"Failed to serialize the minimal {model_type}: %v",
{III}err,
{II})
{II}return
{I}}}

{I}source := "<minimal {model_type}>"

{I}deserialized, deseriaErr := aasjsonization.{deserialization_function}(
{II}jsonable,
{I})
{I}ok := assertNoDeserializationError(t, deseriaErr, source)
{I}if !ok {{
{II}return
{I}}}

{I}anotherJsonable, seriaErr := aasjsonization.ToJsonable(deserialized)
{I}ok = assertNoSerializationError(t, seriaErr, source)
{I}if !ok {{
{II}return
{I}}}

{I}ok = assertSerializationEqualsDeserialization(
{II}t,
{II}jsonable,
{II}anotherJsonable,
{II}source,
{I})
{I}if !ok {{
{II}return
{I}}}
}}"""
            )
        )

        # NOTE (mristin, 2023-06-08):
        # We test here only abstract classes as the concrete classes are going
        # to be already tested in ``generated_of_concrete_classes_test.go``
        if isinstance(our_type, intermediate.AbstractClass):
            test_name = golang_naming.function_name(
                (Identifier(f"test_{our_type.name}_deserialization_fail"))
            )

            blocks.append(
                Stripped(
                    f"""\
func {test_name}(t *testing.T) {{
{I}jsonable := any("this is not an object")

{I}_, err := aasjsonization.{deserialization_function}(
{II}jsonable,
{I})

{I}if err == nil {{
{II}t.Fatal("Expected an error, but got none.")
{II}return
{I}}}

{I}deseriaErr, ok := err.(*aasjsonization.DeserializationError)
{I}if !ok {{
{II}t.Fatalf("Expected a de-serialization error, but got: %v", err)
{II}return
{I}}}

{I}pathString := deseriaErr.PathString()
{I}if len(pathString) != 0 {{
{II}t.Fatalf(
{III}"Expected an empty path in error, but got: %s",
{III}pathString,
{II})
{II}return
{I}}}

{I}expectedMessage :=
{II}"Expected a JSON object, but got string"

{I}if deseriaErr.Message != expectedMessage {{
{II}t.Fatalf(
{III}"Expected the deserialization error:\\n%s\\n, but got:\\n%s",
{III}expectedMessage, deseriaErr.Message,
{II})
{II}return
{I}}}
}}"""
                )
            )

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = (
        repo_root / "jsonization/test/generated_of_classes_with_descendants_test.go"
    )
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
