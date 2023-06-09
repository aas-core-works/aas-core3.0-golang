"""Generate the code to test ``Is*`` functions."""

import io
import os
import pathlib
import sys
from typing import List, Optional

import aas_core_codegen.naming
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

    # noinspection PyListCreation
    blocks = [
        Stripped("// Test `XxxOrDefault` functions."),
        Stripped("package types_xxx_or_default_test"),
        warning,
        Stripped(
            f"""\
import (
{I}"path/filepath"
{I}"fmt"
{I}"encoding/json"
{I}"os"
{I}"reflect"
{I}"strings"
{I}"testing"
{I}aasstringification "github.com/aas-core-works/aas-core3.0-golang/stringification"
{I}aastesting "github.com/aas-core-works/aas-core3.0-golang/aastesting"
{I}aastypes "github.com/aas-core-works/aas-core3.0-golang/types"
)"""
        ),
        Stripped(
            f"""\
// Represent explicitly a literal of an enumeration.
type enumerationLiteral struct {{
{I}enumerationName string
{I}literalName string
}}

func (el *enumerationLiteral) String() string {{
{I}return fmt.Sprintf("%s.%s", el.enumerationName, el.literalName)
}}

// Marshal the value as JSON, or panic otherwise.
func mustJSONMarshal(value interface{{}}) string {{
{I}bb, err := json.Marshal(value)
{I}if err != nil {{
{II}panic(
{III}fmt.Sprintf(
{IIII}"Failed to marshal the value %v to JSON: %s",
{IIII}value, err.Error(),
{III}),
{II})
{I}}}

{I}return string(bb)
}}

func stringify(value interface{{}}) (got string) {{
{I}if value == nil {{
{II}got = mustJSONMarshal(value)
{I}}} else {{
{II}// See: https://stackoverflow.com/questions/38748098/golang-type-switch-how-to-match-a-generic-slice-array-map-chan
{II}reflected := reflect.ValueOf(value)

{II}if reflected.Kind() == reflect.Slice {{
{III}parts := make([]string, reflected.Len())

{III}for i := 0; i < reflected.Len(); i++ {{
{IIII}item := reflected.Index(i)
{IIII}parts[i] = stringify(item)
{III}}}

{III}got = fmt.Sprintf("[%s]", strings.Join(parts, ", "))
{II}}} else {{
{III}switch casted := value.(type) {{
{III}case bool:
{IIII}got = mustJSONMarshal(casted)
{III}case int:
{IIII}got = mustJSONMarshal(casted)
{III}case string:
{IIII}got = mustJSONMarshal(casted)
{III}case []byte:
{IIII}got = fmt.Sprintf("%d byte(s)", len(casted))
{III}case *enumerationLiteral:
{IIII}got = casted.String()
{III}case aastypes.IClass:
{IIII}got = aastesting.TraceMark(casted)
{III}default:
{IIII}panic(
{IIII}{I}fmt.Sprintf(
{IIII}{II}"We do not know hot to represent the value of type %T: %v",
{IIII}{II}value, value,
{IIII}{I}),
{IIII})
{III}}}
{II}}}
{I}}}

{I}return
}}

// Represent `value` such that we can immediately check whether it is the default value
// or the set one.
//
// We compare it against the recorded golden file, if not [aastesting.RecordMode].
// If there are differences, a `message` is set.
//
// Otherwise, when [aastesting.RecordMode] is set, we re-record the golden file.
func compareOrRerecordValue(
{I}value interface{{}},
{I}expectedPath string,
) (message *string) {{
{I}got := stringify(value)

{I}// NOTE (mristin, 2023-06-07):
{I}// Add a new line for POSIX systems.
{I}got += "\\n"

{I}if aastesting.RecordMode {{
{II}parent := filepath.Dir(expectedPath)
{II}err := os.MkdirAll(parent, os.ModePerm)
{II}if err != nil {{
{III}panic(
{IIII}fmt.Sprintf(
{IIII}{I}"Failed to create the directory %s: %s", parent, err.Error(),
{IIII}),
{III})
{II}}}

{II}err = os.WriteFile(expectedPath, []byte(got), 0644)
{II}if err != nil {{
{III}panic(
{IIII}fmt.Sprintf(
{IIII}{I}"Failed to write to the file %s: %s", expectedPath, err.Error(),
{IIII}),
{III})
{II}}}
{I}}} else {{
{II}bb, err := os.ReadFile(expectedPath)
{II}if err != nil {{
{III}panic(
{IIII}fmt.Sprintf(
{IIII}{I}"Failed to read from file %s: %s", expectedPath, err.Error(),
{IIII}),
{III})
{II}}}

{II}expected := string(bb)

{II}// NOTE (mristin, 2023-06-07):
{II}// Git automatically strips and adds `\\r`, so we have to remove it here
{II}// to obtain a canonical text.
{II}expected = strings.Replace(expected, "\\r", "", -1)

{II}if expected != got {{
{III}text := fmt.Sprintf(
{IIII}"What we got differs from the expected in %s. " +
{IIII}"We got:\\n%s\\nWe expected:\\n%s",
{IIII}expectedPath, got, expected,
{III})
{III}message = &text
{II}}}
{I}}}

{I}return
}}"""
        ),
    ]  # type: List[Stripped]

    for our_type in symbol_table.our_types:
        if not isinstance(our_type, intermediate.ConcreteClass):
            continue

        x_or_default_methods = [
            method for method in our_type.methods if method.name.endswith("_or_default")
        ]  # type: List[intermediate.Method]

        model_type = aas_core_codegen.naming.json_model_type(our_type.name)

        for method in x_or_default_methods:
            method_name = golang_naming.method_name(method.name)

            result_enum = None  # type: Optional[intermediate.Enumeration]
            assert method.returns is not None, (
                f"Expected all X_or_default to return something, "
                f"but got None for {our_type}.{method.name}"
            )

            if isinstance(
                method.returns, intermediate.OurTypeAnnotation
            ) and isinstance(method.returns.our_type, intermediate.Enumeration):
                result_enum = method.returns.our_type

            if result_enum is None:
                value_assignment_snippet = Stripped(
                    f"value := instance.{method_name}()"
                )
            else:
                enum_to_string_name = golang_naming.function_name(
                    Identifier(f"must_{result_enum.name}_to_string")
                )

                value_assignment_snippet = Stripped(
                    f"""\
value := &enumerationLiteral{{
{I}enumerationName: "QualifierKind",
{I}literalName: aasstringification.{enum_to_string_name}(
{II}instance.{method_name}(),
{I}),
}}"""
                )

            test_function_name = golang_naming.function_name(
                Identifier(f"Test_{our_type.name}_{method.name}_default")
            )

            must_load_minimal_name = golang_naming.function_name(
                Identifier(f"must_load_minimal_{our_type.name}")
            )

            # noinspection SpellCheckingInspection
            blocks.append(
                Stripped(
                    f"""\
func {test_function_name}(t *testing.T) {{
{I}instance := aastesting.{must_load_minimal_name}()

{I}{indent_but_first_line(value_assignment_snippet, I)}

{I}expectedPth := filepath.Join(
{II}aastesting.TestDataDir,
{II}"XxxOrDefault",
{II}{golang_common.string_literal(model_type)},
{II}"{method_name}.default.txt",
{I})

{I}message := compareOrRerecordValue(
{II}value,
{II}expectedPth,
{I})

{I}if message != nil {{
{II}t.Fatal(*message)
{I}}}
}}"""
                )
            )

            test_function_name = golang_naming.function_name(
                Identifier(f"Test_{our_type.name}_{method.name}_non_default")
            )

            must_load_maximal_name = golang_naming.function_name(
                Identifier(f"must_load_maximal_{our_type.name}")
            )

            # noinspection SpellCheckingInspection
            blocks.append(
                Stripped(
                    f"""\
func {test_function_name}(t *testing.T) {{
{I}instance := aastesting.{must_load_maximal_name}()

{I}{indent_but_first_line(value_assignment_snippet, I)}

{I}expectedPth := filepath.Join(
{II}aastesting.TestDataDir,
{II}"XxxOrDefault",
{II}{golang_common.string_literal(model_type)},
{II}"{method_name}.non-default.txt",
{I})

{I}message := compareOrRerecordValue(
{II}value,
{II}expectedPth,
{I})

{I}if message != nil {{
{II}t.Fatal(*message)
{I}}}
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

    target_pth = (
        repo_root / "types/xxx_or_default_test/generated_xxx_or_default_test.go"
    )
    target_pth.write_text(writer.getvalue(), encoding="utf-8")

    return 0


if __name__ == "__main__":
    sys.exit(main())
