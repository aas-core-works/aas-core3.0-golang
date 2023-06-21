"""Generate the code to test the xmlization of concrete classes."""

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
)

import test_codegen.common
import test_codegen.test_data_io


def _generate_for(
    cls: intermediate.ConcreteClass, container_cls: intermediate.ConcreteClass
) -> List[Stripped]:
    """Generate the tests for a self-contained class."""
    model_type_literal = golang_common.string_literal(
        aas_core_codegen.naming.json_model_type(cls.name)
    )

    if cls is container_cls:
        deserialization_function = golang_naming.function_name(
            Identifier(f"Unmarshal_{cls.name}")
        )
        contained_in_literal = golang_common.string_literal("SelfContained")
    else:
        deserialization_function = golang_naming.function_name(
            Identifier(f"Unmarshal_{container_cls.name}")
        )
        contained_in_literal = golang_common.string_literal(
            f"ContainedIn{container_cls.name}"
        )

    blocks = []  # type: List[Stripped]

    test_name = golang_naming.function_name(
        Identifier(f"Test_{cls.name}_round_trip_OK")
    )

    blocks.append(
        Stripped(
            f"""\
func {test_name}(t *testing.T) {{
{I}pths := aastesting.FindFilesBySuffixRecursively(
{II}filepath.Join(
{III}aastesting.TestDataDir,
{III}"Xml",
{III}{contained_in_literal},
{III}"Expected",
{III}{model_type_literal},
{II}),
{II}".xml",
{I})
{I}sort.Strings(pths)

{I}for _, pth := range pths {{
{II}bb, err := os.ReadFile(pth)
{II}if err != nil {{
{III}t.Fatalf("Failed to read the file %s: %s", pth, err.Error())
{III}return
{II}}}
{II}text := string(bb)

{II}decoder := xml.NewDecoder(strings.NewReader(text))

{II}deserialized, deseriaErr := aasxmlization.{deserialization_function}(decoder)
{II}ok := assertNoDeserializationError(t, deseriaErr, pth)
{II}if !ok {{
{III}return
{II}}}

{II}buf := &bytes.Buffer{{}}
{II}encoder := xml.NewEncoder(buf)
{II}encoder.Indent("", "\\t")

{II}seriaErr := aasxmlization.Marshal(encoder, deserialized, true)
{II}ok = assertNoSerializationError(t, seriaErr, pth)
{II}if !ok {{
{III}return
{II}}}

{II}roundTrip := string(buf.Bytes())

{II}ok = assertSerializationEqualsDeserialization(
{III}t,
{III}text,
{III}roundTrip,
{III}pth,
{II})
{II}if !ok {{
{III}return
{II}}}
{I}}}
}}"""
        )
    )

    test_name = golang_naming.function_name(
        Identifier(f"Test_{cls.name}_deserialization_fail")
    )

    blocks.append(
        Stripped(
            f"""\
func {test_name}(t *testing.T) {{
{I}for _, cause := range causesForDeserializationFailure {{
{II}pths := aastesting.FindFilesBySuffixRecursively(
{III}filepath.Join(
{IIII}aastesting.TestDataDir,
{IIII}"Xml",
{IIII}{contained_in_literal},
{IIII}"Unexpected",
{IIII}cause,
{IIII}{model_type_literal},
{III}),
{III}".xml",
{II})
{II}sort.Strings(pths)

{II}for _, pth := range pths {{
{III}relPth, err := filepath.Rel(aastesting.TestDataDir, pth)
{III}if err != nil {{
{IIII}panic(
{IIIII}fmt.Sprintf(
{IIIII}{I}"Failed to compute the relative path of %s to %s: %s",
{IIIII}{I}aastesting.TestDataDir, pth, err.Error(),
{IIIII}),
{IIII})
{III}}}

{III}expectedPth := filepath.Join(
{IIII}aastesting.TestDataDir,
{IIII}"DeserializationError",
{IIII}filepath.Dir(relPth),
{IIII}filepath.Base(relPth)+".error",
{III})

{III}bb, err := os.ReadFile(pth)
{III}if err != nil {{
{IIII}t.Fatalf("Failed to read the file %s: %s", pth, err.Error())
{IIII}return
{III}}}
{III}text := string(bb)

{III}decoder := xml.NewDecoder(strings.NewReader(text))

{III}_, deseriaErr := aasxmlization.UnmarshalEnvironment(decoder)
{III}ok := assertIsDeserializationErrorAndEqualsExpectedOrRecord(
{IIII}t, deseriaErr, pth, expectedPth,
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
        Stripped("package xmlization_test"),
        warning,
        Stripped(
            f"""\
import (
{I}"bytes"
{I}"path/filepath"
{I}"fmt"
{I}"os"
{I}"sort"
{I}"strings"
{I}"testing"
{I}"encoding/xml"
{I}aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
{I}aasxmlization "github.com/aas-core-works/aas-core3.0-golang/xmlization"
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

        blocks.extend(_generate_for(cls=our_type, container_cls=container_cls))

    blocks.append(warning)

    writer = io.StringIO()
    for i, block in enumerate(blocks):
        if i > 0:
            writer.write("\n\n")

        writer.write(block)

    writer.write("\n")

    target_pth = repo_root / "xmlization/test/generated_of_concrete_classes_test.go"
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
