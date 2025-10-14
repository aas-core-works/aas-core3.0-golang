"""Generate the code to perform a comparison of deep equality on instances."""

import io
import os
import pathlib
import sys
from typing import List, Optional

from aas_core_codegen import intermediate
from aas_core_codegen.common import (
    Stripped,
    Identifier,
    assert_never,
    indent_but_first_line,
)
from aas_core_codegen.golang import (
    naming as golang_naming,
    pointering as golang_pointering,
)
from aas_core_codegen.golang.common import (
    INDENT as I,
    INDENT2 as II,
)

import test_codegen.common


def _generate_for_cls(cls: intermediate.ConcreteClass) -> Stripped:
    """Generate the code for the deep equality function for ``cls``."""
    blocks = []  # type: List[Stripped]

    for prop in cls.properties:
        that_var = golang_naming.variable_name(Identifier(f"that_{prop.name}"))
        other_var = golang_naming.variable_name(Identifier(f"other_{prop.name}"))

        getter_name = golang_naming.getter_name(prop.name)

        subblocks = [
            Stripped(
                f"""\
{that_var} := that.{getter_name}()
{other_var} := other.{getter_name}()"""
            )
        ]  # type: List[Stripped]

        if isinstance(prop.type_annotation, intermediate.OptionalTypeAnnotation):
            subblocks.append(
                Stripped(
                    f"""\
if
{I}({that_var} == nil && {other_var} != nil) ||
{I}({that_var} != nil && {other_var} == nil) {{
{I}return false
}}"""
                )
            )

        cmp_subblock = None  # type: Optional[Stripped]

        type_anno = intermediate.beneath_optional(prop.type_annotation)

        if isinstance(type_anno, intermediate.PrimitiveTypeAnnotation) or (
            isinstance(type_anno, intermediate.OurTypeAnnotation)
            and isinstance(
                type_anno.our_type,
                (intermediate.ConstrainedPrimitive, intermediate.Enumeration),
            )
        ):
            primitive_type = intermediate.try_primitive_type(type_anno)

            if golang_pointering.is_pointer_type(prop.type_annotation):
                cmp_subblock = Stripped(
                    f"""\
if *{that_var} != *{other_var} {{
{I}return false
}}"""
                )
            else:
                if primitive_type is intermediate.PrimitiveType.BYTEARRAY:
                    cmp_subblock = Stripped(
                        f"""\
if !bytes.Equal(
{I}{that_var},
{I}{other_var},
) {{
{I}return false
}}"""
                    )
                else:
                    cmp_subblock = Stripped(
                        f"""\
if {that_var} != {other_var} {{
{I}return false
}}"""
                    )

        elif isinstance(type_anno, intermediate.OurTypeAnnotation):
            if isinstance(type_anno.our_type, intermediate.Enumeration):
                raise AssertionError("Should have been handled before")

            elif isinstance(type_anno.our_type, intermediate.ConstrainedPrimitive):
                raise AssertionError("Should have been handled before")

            elif isinstance(
                type_anno.our_type,
                (intermediate.AbstractClass, intermediate.ConcreteClass),
            ):
                cmp_subblock = Stripped(
                    f"""\
if !DeepEqual(
{I}{that_var},
{I}{other_var},
) {{
{I}return false
}}"""
                )

            else:
                assert_never(type_anno.our_type)
        elif isinstance(type_anno, intermediate.ListTypeAnnotation):
            assert isinstance(
                type_anno.items, intermediate.OurTypeAnnotation
            ) and isinstance(
                type_anno.items.our_type,
                (intermediate.AbstractClass, intermediate.ConcreteClass),
            ), (
                f"NOTE (mristin, 2023-03-29): We expect only lists of classes "
                f"at the moment, but you specified {type_anno}. "
                f"Please contact the developers if you need this feature."
            )

            cmp_subblock = Stripped(
                f"""\
if 
{I}len({that_var}) !=
{I}len({other_var}) {{
{I}return false
}}
for i := range {that_var} {{
{I}if !DeepEqual(
{II}{that_var}[i],
{II}{other_var}[i],
{I}) {{
{II}return false
{I}}}
}}"""
            )
        else:
            assert_never(type_anno)

        assert cmp_subblock is not None

        if isinstance(prop.type_annotation, intermediate.OptionalTypeAnnotation):
            cmp_subblock = Stripped(
                f"""\
if {that_var} != nil {{
{I}{indent_but_first_line(cmp_subblock, I)}
}}"""
            )

        subblocks.append(cmp_subblock)

        blocks.append(Stripped("\n".join(subblocks)))

    blocks.append(Stripped("return true"))

    body = "\n\n".join(blocks)

    interface_name = golang_naming.interface_name(cls.name)

    function_name = golang_naming.private_function_name(
        Identifier(f"deep_equal_{cls.name}")
    )

    return Stripped(
        f"""\
// Perform a comparison for deep equality between `that` and `other` instance.
//
// The deep equality means that all the properties are checked for equality recursively.
func {function_name}(
{I}that aastypes.{interface_name},
{I}other aastypes.{interface_name},
) bool {{
{I}{indent_but_first_line(body, I)}
}}"""
    )


def _generate_dispatch_function(symbol_table: intermediate.SymbolTable) -> Stripped:
    """Generate the main ``DeepEqual`` function which then dispatches as necessary."""
    blocks = [
        Stripped(
            f"""\
if that.ModelType() != other.ModelType() {{
{I}return false
}}"""
        )
    ]

    cases = []  # type: List[Stripped]
    for cls in symbol_table.concrete_classes:
        model_type_literal = golang_naming.enum_literal_name(
            enumeration_name=Identifier("Model_type"), literal_name=cls.name
        )

        deep_equal_func_name = golang_naming.private_function_name(
            Identifier(f"deep_equal_{cls.name}")
        )

        interface_name = golang_naming.interface_name(cls.name)

        cases.append(
            Stripped(
                f"""\
case aastypes.{model_type_literal}:
{I}return {deep_equal_func_name}(
{II}that.(aastypes.{interface_name}),
{II}other.(aastypes.{interface_name}),
{I})"""
            )
        )

    cases_joined = "\n".join(cases)

    switch_stmt = Stripped(
        f"""\
switch that.ModelType() {{
{cases_joined}
}}"""
    )

    blocks.append(switch_stmt)

    blocks.append(
        Stripped(
            """\
panic(fmt.Sprintf("Unexpected model type: %d", that.ModelType()))"""
        )
    )

    body = "\n\n".join(blocks)

    return Stripped(
        f"""\
func DeepEqual(
{I}that aastypes.IClass,
{I}other aastypes.IClass,
) bool {{
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
        Stripped("package aastesting"),
        warning,
        Stripped(
            f"""\
import (
{I}"bytes"
{I}"fmt"
{I}aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)"""
        ),
    ]  # type: List[Stripped]

    for cls in symbol_table.concrete_classes:
        blocks.append(_generate_for_cls(cls=cls))

    blocks.append(_generate_dispatch_function(symbol_table=symbol_table))

    blocks.append(warning)

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = repo_root / "aastesting/deep_equal.generated.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
