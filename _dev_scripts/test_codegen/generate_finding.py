"""Generate the code to find the first occurrence of an instance in a container."""

import io
import os
import pathlib
import sys
from typing import List

from aas_core_codegen import intermediate
from aas_core_codegen.common import Stripped, Identifier
from aas_core_codegen.golang import (
    naming as golang_naming,
)
from aas_core_codegen.golang.common import (
    INDENT as I,
    INDENT2 as II,
    INDENT3 as III,
    INDENT4 as IIII,
)

import test_codegen.common


def _generate_find_function(cls: intermediate.ClassUnion) -> Stripped:
    """Generate the code of the finding function for ``cls``."""
    function_name = golang_naming.function_name(Identifier(f"must_find_{cls.name}"))
    interface_name = golang_naming.interface_name(cls.name)

    is_function_name = golang_naming.function_name(Identifier(f"is_{cls.name}"))

    return Stripped(
        f"""\
// Find the first instance of [aastypes.{interface_name}] in `container`
// including the container itself.
//
// If no instance could be find, panic.
func {function_name}(
{I}container aastypes.IClass,
) (result aastypes.{interface_name}) {{
{I}ok := aastypes.{is_function_name}(container)
{I}if ok {{
{II}result = container.(aastypes.{interface_name})
{II}return
{I}}}

{I}container.Descend(func(that aastypes.IClass) (abort bool) {{
{II}abort = aastypes.{is_function_name}(that)
{II}if abort {{
{III}result = that.(aastypes.{interface_name})
{II}}}
{II}return
{I}}})

{I}if result == nil {{
{II}panic(
{III}fmt.Sprintf(
{IIII}"Could not find an instance of {interface_name} " +
{IIII}"in the container of type %T: %v",
{IIII}container, container,
{III}),
{II})
{I}}}
{I}return
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
        Stripped("package aastesting"),
        warning,
        Stripped(
            f"""\
import (
{I}"fmt"
{I}aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)"""
        ),
    ]  # type: List[Stripped]

    for our_type in symbol_table.our_types:
        if not isinstance(
            our_type, (intermediate.AbstractClass, intermediate.ConcreteClass)
        ):
            continue

        blocks.append(_generate_find_function(cls=our_type))

    blocks.append(warning)

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = repo_root / "aastesting/finding.generated.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
