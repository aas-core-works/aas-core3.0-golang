"""Generate the code to test ``Is*`` functions."""

import io
import os
import pathlib
import sys
from typing import List

from aas_core_codegen import intermediate
from aas_core_codegen.common import Stripped, Identifier, indent_but_first_line
from aas_core_codegen.golang import (
    naming as golang_naming,
)
from aas_core_codegen.golang.common import (
    INDENT as I,
    INDENT2 as II,
)

import test_codegen.common


def _generate_for_cls(
    cls: intermediate.ConcreteClass, symbol_table: intermediate.SymbolTable
) -> Stripped:
    """Generate the test function."""
    must_load_minimal_name = golang_naming.function_name(
        Identifier(f"must_load_minimal_{cls.name}")
    )

    interface_name = golang_naming.interface_name(cls.name)

    block = [
        Stripped(
            f"""\
instance := aastesting.{must_load_minimal_name}()"""
        )
    ]  # type: List[Stripped]

    for our_type in symbol_table.our_types:
        if not isinstance(
            our_type, (intermediate.AbstractClass, intermediate.ConcreteClass)
        ):
            continue

        is_function_name = golang_naming.function_name(
            Identifier(f"is_{our_type.name}")
        )

        if cls.is_subclass_of(our_type):
            block.append(
                Stripped(
                    f"""\
if !aastypes.{is_function_name}(instance) {{
{I}t.Errorf(
{II}"Expected {is_function_name} to be true on an instance " +
{II}"of {interface_name} with runtime type %T and with model type %v",
{II}instance, instance.ModelType(),
{I})
}}"""
                )
            )
        else:
            block.append(
                Stripped(
                    f"""\
if aastypes.{is_function_name}(instance) {{
{I}t.Errorf(
{II}"Expected {is_function_name} to be false on an instance " +
{II}"of {interface_name} with runtime type %T and with model type %v",
{II}instance, instance.ModelType(),
{I})
}}"""
                )
            )

    body = "\n\n".join(block)

    test_name = golang_naming.function_name(
        Identifier(f"test_is_Xxx_on_an_instance_of_{cls.name}")
    )

    return Stripped(
        f"""\
func {test_name}(t *testing.T) {{
{I}{indent_but_first_line(body, I)}
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
        Stripped("// Test `IsXxx` functions."),
        Stripped("package types_is_xxx_test"),
        warning,
        Stripped(
            f"""\
import (
{I}"testing"
{I}aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
{I}aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)"""
        ),
    ]  # type: List[Stripped]

    for our_type in symbol_table.our_types:
        if not isinstance(our_type, intermediate.ConcreteClass):
            continue

        blocks.append(_generate_for_cls(cls=our_type, symbol_table=symbol_table))

    blocks.append(warning)

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = repo_root / "types/is_xxx_test/generated_is_model_type_test.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
