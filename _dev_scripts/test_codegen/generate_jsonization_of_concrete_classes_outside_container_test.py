"""Generate the test code for the jsonization of classes outside a container."""

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
from test_codegen import test_data_io


def main() -> int:
    """Execute the main routine."""
    symbol_table = test_codegen.common.load_symbol_table()

    this_path = pathlib.Path(os.path.realpath(__file__))
    repo_root = this_path.parent.parent.parent

    test_data_dir = repo_root / "testdata"

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

    environment_cls = symbol_table.must_find_concrete_class(
        aas_core_codegen.common.Identifier("Environment")
    )

    for our_type in symbol_table.our_types:
        if not isinstance(our_type, intermediate.ConcreteClass):
            continue

        # fmt: off
        container_cls = (
            test_data_io.determine_container_class(
                cls=our_type,
                test_data_dir=test_data_dir,
                environment_cls=environment_cls
            )
        )
        # fmt: on

        if container_cls is our_type:
            # NOTE (mristin, 2023-06-07):
            # These classes are tested already in
            # jsonization/of_concrete_classes_test.go. We only need to test for class
            # instances contained in a container.
            continue

        must_load_maximal_name = golang_naming.function_name(
            Identifier(f"must_load_maximal_{our_type.name}")
        )

        model_type = aas_core_codegen.naming.json_model_type(our_type.name)

        deserialization_function = golang_naming.function_name(
            Identifier(f"{our_type.name}_from_jsonable")
        )

        test_name = golang_naming.function_name(
            (Identifier(f"test_{our_type.name}_round_trip_OK_outside_container"))
        )

        blocks.append(
            Stripped(
                f"""\
func {test_name}(t *testing.T) {{
{I}instance := aastesting.{must_load_maximal_name}()

{I}jsonable, err := aasjsonization.ToJsonable(instance)
{I}if err != nil {{
{II}t.Fatalf(
{III}"Failed to serialize the maximal {model_type}: %s",
{III}err.Error(),
{II})
{II}return
{I}}}

{I}source := "<maximal {model_type}>"

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

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = (
        repo_root
        / "jsonization/test/generated_of_concrete_classes_outside_container_test.go"
    )
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
