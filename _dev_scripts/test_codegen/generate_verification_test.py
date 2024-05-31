"""Generate the code to test the verification."""

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
    INDENT3 as III,
    INDENT4 as IIII,
    INDENT5 as IIIII,
    INDENT6 as IIIIII,
)

import test_codegen.common
import test_codegen.test_data_io


def _generate_for_cls(
    cls: intermediate.ConcreteClass, container_cls: intermediate.ConcreteClass
) -> List[Stripped]:
    """Generate the tests for a class."""
    model_type_literal = golang_common.string_literal(
        aas_core_codegen.naming.json_model_type(cls.name)
    )

    contained_in_literal = golang_common.string_literal(
        "SelfContained" if cls is container_cls else f"ContainedIn{container_cls.name}"
    )

    deserialization_function = golang_naming.function_name(
        Identifier(f"{cls.name}_from_jsonable")
        if cls is container_cls
        else Identifier(f"{container_cls.name}_from_jsonable")
    )

    blocks = []  # type: List[Stripped]

    test_name = golang_naming.function_name(Identifier(f"Test_{cls.name}_OK"))

    blocks.append(
        Stripped(
            f"""\
func {test_name}(t *testing.T) {{
{I}pths := aastesting.FindFilesBySuffixRecursively(
{II}filepath.Join(
{III}aastesting.TestDataDir,
{III}"Json",
{III}{contained_in_literal},
{III}"Expected",
{III}{model_type_literal},
{II}),
{II}".json",
{I})
{I}sort.Strings(pths)

{I}for _, pth := range pths {{
{II}jsonable := aastesting.MustReadJsonable(
{III}pth,
{II})

{II}deserialized, deseriaErr := aasjsonization.{deserialization_function}(
{III}jsonable,
{II})
{II}if deseriaErr != nil {{
{III}t.Fatalf(
{IIII}"Unexpected deserialization error from %s: %s",
{IIII}pth, deseriaErr.Error(),
{III})
{III}return
{II}}}

{II}var errors []*aasverification.VerificationError
{II}aasverification.Verify(
{III}deserialized,
{III}func(veriErr *aasverification.VerificationError) (abort bool) {{
{IIII}errors = append(errors, veriErr)
{IIII}return
{III}}},
{II})

{II}ok := assertNoVerificationErrors(
{III}t,
{III}deserialized,
{III}pth,
{II})
{II}if !ok {{
{III}return
{II}}}
{I}}}
}}"""
        )
    )

    test_name = golang_naming.function_name(Identifier(f"Test_{cls.name}_fail"))

    blocks.append(
        Stripped(
            f"""\
func {test_name}(t *testing.T) {{
{I}pattern := filepath.Join(
{II}aastesting.TestDataDir,
{II}"Json",
{II}{contained_in_literal},
{II}"Unexpected",
{II}"Invalid",
{II}"*",  // This asterisk represents the cause.
{II}{model_type_literal},
{I})
 
{I}causeDirs, err := filepath.Glob(pattern)
{I}if err != nil {{
{II}panic(
{III}fmt.Sprintf(
{IIII}"Failed to find cause directories matching %s: %s",
{IIII}pattern, err.Error(),
{III}),
{II})
{I}}}

{I}for _, causeDir := range causeDirs {{
{II}pths := aastesting.FindFilesBySuffixRecursively(
{III}causeDir,
{III}".json",
{II})
{II}sort.Strings(pths)

{II}for _, pth := range pths {{
{III}jsonable := aastesting.MustReadJsonable(
{IIII}pth,
{III})

{III}relPth, err := filepath.Rel(aastesting.TestDataDir, pth)
{III}if err != nil {{
{IIII}panic(
{IIIII}fmt.Sprintf(
{IIIIII}"Failed to compute the relative path of %s to %s: %s",
{IIIIII}aastesting.TestDataDir, pth, err.Error(),
{IIIII}),
{IIII})
{III}}}

{III}expectedPth := filepath.Join(
{IIII}aastesting.TestDataDir,
{IIII}"VerificationError",
{IIII}filepath.Dir(relPth),
{IIII}filepath.Base(relPth)+".errors",
{III})

{III}deserialized, deseriaErr := aasjsonization.{deserialization_function}(
{IIII}jsonable,
{III})
{III}if deseriaErr != nil {{
{IIII}t.Fatalf(
{IIIII}"Unexpected deserialization error from %s: %s",
{IIIII}pth, deseriaErr.Error(),
{IIII})
{IIII}return
{III}}}

{III}var errors []*aasverification.VerificationError
{III}aasverification.Verify(
{IIII}deserialized,
{IIII}func(err *aasverification.VerificationError) (abort bool) {{
{IIIII}errors = append(errors, err)
{IIIII}return
{IIII}}},
{III})

{III}ok := assertEqualsExpectedOrRerecordVerificationErrors(
{IIII}t,
{IIII}errors,
{IIII}pth,
{IIII}expectedPth,
{III})
{III}if !ok {{
{IIII}return
{III}}}
{II}}}
{I}}}
}}"""
        )
    )

    return blocks


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
        Stripped("package verification_test"),
        warning,
        Stripped(
            f"""\
import (
{I}"fmt"
{I}"path/filepath"
{I}"sort"
{I}"testing"
{I}aasjsonization "github.com/aas-core-works/aas-core3.0-golang/jsonization"
{I}aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
{I}aasverification "github.com/aas-core-works/aas-core3.0-golang/verification"
)"""
        ),
    ]  # type: List[Stripped]

    environment_cls = symbol_table.must_find_concrete_class(Identifier("Environment"))
    test_data_dir = repo_root / "testdata"

    for our_type in symbol_table.our_types:
        if not isinstance(our_type, intermediate.ConcreteClass):
            continue

        # fmt: off
        container_cls = (
            test_codegen.test_data_io.determine_container_class(
                cls=our_type,
                test_data_dir=test_data_dir,
                environment_cls=environment_cls
            )
        )
        # fmt: on

        blocks.extend(_generate_for_cls(cls=our_type, container_cls=container_cls))

    blocks.append(warning)

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = repo_root / "verification/test/generated_test.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
