"""
Generate the mapping AAS_submodel_elements 🠒 ``Is X`` function and print it to stdout.

Use this script to write implementation-specific snippet
``submodel_element_is_of_type.go``.
"""

import sys
from typing import List

from aas_core_codegen.common import Identifier, Stripped, indent_but_first_line
from aas_core_codegen.golang import naming as golang_naming
from aas_core_codegen.golang.common import (
    INDENT as I,
    INDENT2 as II,
)

import test_codegen.common


def main() -> int:
    """Execute the main routine."""
    symbol_table = test_codegen.common.load_symbol_table()

    aas_submodel_elements = symbol_table.must_find_enumeration(
        Identifier("AAS_submodel_elements")
    )

    case_stmts = []  # type: List[str]

    for literal in aas_submodel_elements.literals:
        literal_name = golang_naming.enum_literal_name(
            enumeration_name=aas_submodel_elements.name, literal_name=literal.name
        )

        function_name = golang_naming.function_name(Identifier(f"is_{literal.name}"))

        case_stmts.append(
            Stripped(
                f"""\
case aastypes.{literal_name}:
{I}return aastypes.{function_name}(
{II}element,
{I})"""
            )
        )

    case_stmts_joined = "\n".join(case_stmts)

    switch_stmt = Stripped(
        f"""\
switch expectedType {{
{case_stmts_joined}
}}"""
    )

    print(
        Stripped(
            f"""\
// Check that `element` is an instance of the interface corresponding to
// `expectedType`.
func SubmodelElementIsOfType(
{I}element aastypes.ISubmodelElement,
{I}expectedType aastypes.AASSubmodelElements,
) bool {{
{I}{indent_but_first_line(switch_stmt, I)}
{I}return false
}}"""
        )
    )

    return 0


if __name__ == "__main__":
    sys.exit(main())
