"""Generate the test code for the enhancing of instances."""

import io
import os
import pathlib
import sys

from aas_core_codegen import intermediate
from aas_core_codegen.common import Stripped, Identifier
from aas_core_codegen.golang import naming as golang_naming
from aas_core_codegen.golang.common import (
    INDENT as I,
    INDENT2 as II,
    INDENT3 as III,
    INDENT4 as IIII,
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
        Stripped("package enhancing_test"),
        warning,
        Stripped(
            f"""\
import (
{I}"testing"
{I}aasenhancing "github.com/aas-core-works/aas-core3.0-golang/enhancing"
{I}aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
{I}aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)"""
        ),
        Stripped(
            f"""\
type Enhancement struct {{
{I}ID int
}}"""
        ),
        Stripped(
            f"""\
func collectIDsAndAssertTheyAreConsecutiveAndTheirCountEqualsNextID(
{I}t *testing.T,
{I}wrapped aastypes.IClass,
{I}nextID int,
) {{
{I}var ids []int

{I}instanceEnh := aasenhancing.MustUnwrap[*Enhancement](wrapped)
{I}ids = append(ids, instanceEnh.ID)

{I}wrapped.Descend(
{II}func(that aastypes.IClass) (abort bool) {{
{III}enh := aasenhancing.MustUnwrap[*Enhancement](that)
{III}ids = append(ids, enh.ID)
{III}return
{II}}},
{I})

{I}if len(ids) != nextID {{
{II}t.Fatalf("Expected to collect %d IDs, but got: %d", len(ids), nextID)
{II}return
{I}}}

{I}for i, id := range ids {{
{II}if id != i {{
{III}t.Fatalf(
{IIII}"Unexpected ID at index %d (starting from 0); expected %d, got %d",
{IIII}i, i, id,
{III})
{II}}}
{I}}}
}}"""
        ),
    ]

    for our_type in symbol_table.our_types:
        if not isinstance(our_type, intermediate.ConcreteClass):
            continue

        must_load_maximal_name = golang_naming.function_name(
            Identifier(f"must_load_maximal_{our_type.name}")
        )

        test_name = golang_naming.function_name(
            (Identifier(f"test_{our_type.name}_wrapped"))
        )

        blocks.append(
            Stripped(
                f"""\
func {test_name}(t *testing.T) {{
{I}instance := aastesting.{must_load_maximal_name}()

{I}nextID := 0
{I}wrapped := aasenhancing.Wrap[*Enhancement](
{II}instance,
{II}func(that aastypes.IClass) (enh *Enhancement, should bool) {{
{III}enh = &Enhancement{{}}
{III}enh.ID = nextID
{III}should = true

{III}nextID++
{III}return
{II}}},
{I})

{I}if !aastesting.DeepEqual(instance, wrapped) {{
{II}t.Fatalf(
{III}"Deep equality failed between the instance and the wrapped: %v %v",
{III}instance, wrapped,
{II})
{I}}}

{I}collectIDsAndAssertTheyAreConsecutiveAndTheirCountEqualsNextID(
{II}t, wrapped, nextID,
{I})
}}"""
            )
        )

        test_name = golang_naming.function_name(
            (Identifier(f"test_{our_type.name}_nothing_wrapped"))
        )

        blocks.append(
            Stripped(
                f"""\
func {test_name}(t *testing.T) {{
{I}instance := aastesting.{must_load_maximal_name}()

{I}wrapped := aasenhancing.Wrap[*Enhancement](
{II}instance,
{II}func(that aastypes.IClass) (enh *Enhancement, should bool) {{
{III}should = false
{III}return
{II}}},
{I})

{I}if !aastesting.DeepEqual(instance, wrapped) {{
{II}t.Fatalf(
{III}"Deep equality failed between the instance and the wrapped: %v %v",
{III}instance, wrapped,
{II})
{I}}}

{I}// Wrapped should be equal to instance by reference as our enhancement factory
{I}// did not wrap anything.
{I}if wrapped != instance {{
{II}t.Fatalf("Unexpected inequality between %v and %v", wrapped, instance)
{I}}}

{I}wrapped.Descend(func (that aastypes.IClass) (abort bool) {{
{II}_, ok := aasenhancing.Unwrap[*Enhancement](that)
{II}if ok {{
{III}t.Fatalf("Unexpected wrapped descendant: %v", that)
{II}}}
{II}return
{I}}})
}}"""
            )
        )

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = repo_root / "enhancing/test/generated_test.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
