"""Generate the code to test de/serialization of enumeration literals."""

import io
import os
import pathlib
import sys
from typing import List

from aas_core_codegen import intermediate
from aas_core_codegen.common import Stripped, Identifier, indent_but_first_line
from aas_core_codegen.golang import (
    naming as golang_naming,
    common as golang_common,
)
from aas_core_codegen.golang.common import (
    INDENT as I,
    INDENT2 as II,
    INDENT3 as III,
)
from icontract import require

import test_codegen.common


# fmt: off
@require(
    lambda enum: len(enum.literals) > 0,
    "Enumeration must have at least one literal as enumerations without literals "
    "can not be tested"
)
# fmt: on
def _generate_round_trip_test_for_enum(
    enum: intermediate.Enumeration,
) -> Stripped:
    """Generate the test that the de-serialization equals the serialization."""
    test_name = golang_naming.function_name(
        Identifier(f"test_{enum.name}_round_trip_OK")
    )

    literals_joined = "\n".join(
        f"{golang_common.string_literal(literal.value)}," for literal in enum.literals
    )

    deserialization_function = golang_naming.function_name(
        Identifier(f"{enum.name}_from_jsonable")
    )

    serialization_function = golang_naming.function_name(
        Identifier(f"{enum.name}_to_jsonable")
    )

    return Stripped(
        f"""\
func {test_name}(t *testing.T) {{
{I}literals := []string{{
{II}{indent_but_first_line(literals_joined, II)}
{I}}}

{I}for _, literal := range literals {{
{II}source := fmt.Sprintf("<string literal %s>", literal)
{II}jsonable := any(literal)

{II}deserialized, deseriaErr := aasjsonization.{deserialization_function}(
{III}jsonable,
{II})
{II}ok := assertNoDeserializationError(t, deseriaErr, source)
{II}if !ok {{
{III}return
{II}}}

{II}anotherJsonable, seriaErr :=
{III}aasjsonization.{serialization_function}(deserialized)
{II}ok = assertNoSerializationError(t, seriaErr, source)
{II}if !ok {{
{III}return
{II}}}

{II}ok = assertSerializationEqualsDeserialization(
{III}t,
{III}jsonable,
{III}anotherJsonable,
{III}source,
{II})
{II}if !ok {{
{III}return
{II}}}
{I}}}
}}"""
    )


def _generate_deserialization_fail_for_enum(enum: intermediate.Enumeration) -> Stripped:
    """Generate the test for de-serialization of an invalid value."""
    test_name = golang_naming.function_name(
        Identifier(f"test_{enum.name}_deserialization_fail")
    )

    deserialization_function = golang_naming.function_name(
        Identifier(f"{enum.name}_from_jsonable")
    )

    enum_name = golang_naming.enum_name(enum.name)

    return Stripped(
        f"""\
func {test_name}(t *testing.T) {{
{I}jsonable := any("THIS-CANNOT-POSSIBLY-BE-VALID")

{I}_, err := aasjsonization.{deserialization_function}(
{II}jsonable,
{I})

{I}if err == nil {{
{II}t.Fatal("Expected a deserialization error, but got none.")
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
{II}"Expected a string representation of {enum_name}, " +
{II}"but got THIS-CANNOT-POSSIBLY-BE-VALID"

{I}if deseriaErr.Message != expectedMessage {{
{II}t.Fatalf(
{III}"Expected the deserialization error:\\n%s\\n, but got:\\n%s",
{III}expectedMessage, deseriaErr.Message,
{II})
{II}return
{I}}}
}}"""
    )


def main() -> int:
    """Execute the main routine."""
    symbol_table = test_codegen.common.load_symbol_table()

    this_path = pathlib.Path(os.path.realpath(__file__))
    repo_root = this_path.parent.parent.parent

    warning = test_codegen.common.generate_warning_comment(
        this_path.relative_to(repo_root)
    )

    # noinspection PyListCreation
    blocks = [
        Stripped("package jsonization_test"),
        warning,
        Stripped(
            f"""\
import (
{I}"fmt"
{I}"testing"
{I}aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
)"""
        ),
    ]  # type: List[Stripped]

    for our_type in symbol_table.our_types:
        if not isinstance(our_type, intermediate.Enumeration):
            continue

        blocks.append(_generate_round_trip_test_for_enum(enum=our_type))
        blocks.append(_generate_deserialization_fail_for_enum(enum=our_type))

    blocks.append(warning)

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = repo_root / "jsonization/test/generated_of_enums_test.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
