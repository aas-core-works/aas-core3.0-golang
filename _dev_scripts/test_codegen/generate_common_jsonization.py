"""Generate the code to load minimal and maximal examples."""

import io
import os
import pathlib
import sys
from typing import List

import aas_core_codegen.naming
from aas_core_codegen import intermediate
from aas_core_codegen.common import Stripped, Identifier, indent_but_first_line
from aas_core_codegen.golang import (
    common as golang_common,
    naming as golang_naming,
)
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

    environment_cls = symbol_table.must_find_concrete_class(Identifier("Environment"))

    test_data_dir = repo_root / "testdata"

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
{I}"path"
{I}aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
{I}aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)"""
        ),
    ]  # type: List[Stripped]

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

        interface_name = golang_naming.interface_name(our_type.name)

        cls_name_json = aas_core_codegen.naming.json_model_type(our_type.name)

        # NOTE (mristin, 2023-05-31):
        # The class is self-contained if the container class is equal to our type.
        if container_cls is our_type:
            deserialization_function = golang_naming.function_name(
                Identifier(f"{our_type.name}_from_jsonable")
            )

            deserialization_snippet = Stripped(
                f"""\
instance, err := aasjsonization.{deserialization_function}(
{I}jsonable,
)
if err != nil {{
{I}panic(
{II}fmt.Sprintf(
{III}"Failed to de-serialize an instance of {interface_name} " +
{III}"from %s: %s",
{III}pth, err.Error(),
{II}),
{I})
}}
var ok bool
result, ok = instance.(aastypes.{interface_name})
if !ok {{
{I}panic(
{II}fmt.Sprintf(
{III}"Expected to find an instance of {interface_name} at %s, " +
{III}"but got an instance of %T: %v",
{III}pth, instance, instance,
{II}),
{I})
}}"""
            )
            container_kind_directory = "SelfContained"
        else:
            deserialization_function = golang_naming.function_name(
                Identifier(f"{container_cls.name}_from_jsonable")
            )

            container_interface_name = golang_naming.interface_name(container_cls.name)

            must_find_name = golang_naming.function_name(
                Identifier(f"must_find_{our_type.name}")
            )

            deserialization_snippet = Stripped(
                f"""\
container, err := aasjsonization.{deserialization_function}(
{I}jsonable,
)
if err != nil {{
{I}panic(
{II}fmt.Sprintf(
{III}"Failed to de-serialize the container as {container_interface_name} " +
{III}"from %s: %s",
{III}pth, err.Error(),
{II}),
{I})
}}
result = {must_find_name}(
{I}container,
)"""
            )

            assert (
                container_cls.name == "Environment"
            ), "Necessary for the container kind directory"
            container_kind_directory = "ContainedInEnvironment"

        must_load_maximal_name = golang_naming.function_name(
            Identifier(f"must_load_maximal_{our_type.name}")
        )

        must_load_minimal_name = golang_naming.function_name(
            Identifier(f"must_load_minimal_{our_type.name}")
        )

        blocks.append(
            Stripped(
                f"""\
// Load a maximal example of [aastypes.{interface_name}] from
// the test data directory.
//
// If there is any error, panic.
func {must_load_maximal_name}(
) (result aastypes.{interface_name}) {{
{I}pth := path.Join(
{II}TestDataDir,
{II}"Json",
{II}{golang_common.string_literal(container_kind_directory)},
{II}"Expected",
{II}{golang_common.string_literal(cls_name_json)},
{II}"maximal.json",
{I})

{I}jsonable := MustReadJsonable(pth)

{I}{indent_but_first_line(deserialization_snippet, I)}
{I}return
}}"""
            )
        )

        blocks.append(
            Stripped(
                f"""\
// Load a minimal example of [aastypes.{interface_name}] from
// the test data directory.
//
// If there is any error, panic.
func {must_load_minimal_name}(
) (result aastypes.{interface_name}) {{
{I}pth := path.Join(
{II}TestDataDir,
{II}"Json",
{II}{golang_common.string_literal(container_kind_directory)},
{II}"Expected",
{II}{golang_common.string_literal(cls_name_json)},
{II}"minimal.json",
{I})

{I}jsonable := MustReadJsonable(pth)

{I}{indent_but_first_line(deserialization_snippet, I)}
{I}return
}}"""
            )
        )

    blocks.append(warning)

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = repo_root / "aastesting/common_jsonization.generated.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
