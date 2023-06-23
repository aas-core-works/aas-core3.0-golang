"""Generate the code to test ``Is*`` functions."""

import io
import os
import pathlib
import sys
from typing import List

import aas_core_codegen.naming
from aas_core_codegen import intermediate
from aas_core_codegen.common import Stripped, Identifier
from aas_core_codegen.golang import (
    naming as golang_naming,
    common as golang_common,
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
    test_function_name = golang_naming.function_name(
        Identifier(f"test_descend_once_on_an_instance_of_{cls.name}")
    )

    must_load_maximal_name = golang_naming.function_name(
        Identifier(f"must_load_maximal_{cls.name}")
    )

    model_type_literal = golang_common.string_literal(
        aas_core_codegen.naming.json_model_type(cls.name)
    )

    return Stripped(
        f"""\
func {test_function_name}(
{I}t *testing.T,
) {{
{I}instance := aastesting.{must_load_maximal_name}()

{I}expectedPth := filepath.Join(
{II}aastesting.TestDataDir,
{II}"DescendOnce",
{II}{model_type_literal},
{II}"maximal.json.trace",
{I})

{I}onlyOnce := true

{I}message := compareOrRerecordTrace(
{II}instance,
{II}expectedPth,
{II}onlyOnce,
{I})
{I}if message != nil {{
{II}t.Fatal(*message)
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
        Stripped("package types_descend_test"),
        warning,
        Stripped(
            f"""\
import (
{I}"path/filepath"
{I}"testing"
{I}aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
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

    target_pth = repo_root / "types/descend_test/generated_descend_once_test.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
