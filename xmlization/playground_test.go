package xmlization

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"testing"
)

func dispatchOnXMLTokensToWriter(decoder *xml.Decoder, w io.Writer) (err error) {
	for {
		token, tokenErr := decoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			} else {
				err = tokenErr
				return
			}
		}

		var writeErr error

		switch et := token.(type) {
		case xml.StartElement:
			_, writeErr = io.WriteString(
				w,
				fmt.Sprintf(
					"Start element: space: %s, local: %s\n",
					et.Name.Space, et.Name.Local,
				),
			)
			if writeErr != nil {
				panic(writeErr.Error())
			}
		case xml.EndElement:
			_, writeErr = io.WriteString(
				w,
				fmt.Sprintf(
					"End element: space: %s, local: %s\n", et.Name.Space, et.Name.Local,
				),
			)
			if writeErr != nil {
				panic(writeErr.Error())
			}
		case xml.CharData:
			text := string(et)

			textJsonBytes, writeErr := json.Marshal(text)
			if writeErr != nil {
				panic(writeErr.Error())
			}

			_, writeErr = io.WriteString(
				w,
				fmt.Sprintf(
					"Char data: %s\n", string(textJsonBytes),
				),
			)
			if writeErr != nil {
				panic(writeErr.Error())
			}
		case xml.Comment:
			text := string(et)

			textJsonBytes, writeErr := json.Marshal(text)
			if writeErr != nil {
				panic(writeErr.Error())
			}

			_, writeErr = io.WriteString(
				w,
				fmt.Sprintf(
					"Comment: %s\n", string(textJsonBytes),
				),
			)
			if writeErr != nil {
				panic(writeErr.Error())
			}
		default:
			// Ignore
		}
	}
	return
}

// NOTE (mristin, 2023-06-12):
// These tests were written during the development of the Go generator. We keep them
// simply as development notes, if we ever need to recall how to do certain steps
// in XML de/serialization.

func TestTokensReadFromDecoder(t *testing.T) {
	text := `<?xml version="1.0" encoding="UTF-8"?><something xmlns="https://admin-shell.io/aas/3/0">
	some <!-- comment --> text
	<childElement>asdf</childElement>
	<emptyElement />
</something>`

	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	writer := new(strings.Builder)

	err := dispatchOnXMLTokensToWriter(decoder, writer)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	got := writer.String()

	// NOTE (mristin, 2023-06-12):
	// We have to be careful when we iterate over content of an element --
	// comments can be intertwined with text.

	expected := `Start element: space: https://admin-shell.io/aas/3/0, local: something
Char data: "\n\tsome "
Comment: " comment "
Char data: " text\n\t"
Start element: space: https://admin-shell.io/aas/3/0, local: childElement
Char data: "asdf"
End element: space: https://admin-shell.io/aas/3/0, local: childElement
Char data: "\n\t"
Start element: space: https://admin-shell.io/aas/3/0, local: emptyElement
End element: space: https://admin-shell.io/aas/3/0, local: emptyElement
Char data: "\n"
End element: space: https://admin-shell.io/aas/3/0, local: something
`

	if got != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, got)
		return
	}
}

func TestWriteWithEncoder(t *testing.T) {
	b := new(strings.Builder)
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")

	namespace := "https://hello.world.com/1/0"

	encoder.EncodeToken(
		xml.StartElement{
			Name: xml.Name{
				Space: namespace,
				Local: "Hello",
			},
		},
	)

	encoder.EncodeToken(
		xml.StartElement{
			Name: xml.Name{
				Local: "Nested",
			},
		},
	)

	encoder.EncodeToken(
		xml.CharData([]byte("some <> text")),
	)

	encoder.EncodeToken(
		xml.EndElement{
			Name: xml.Name{
				Local: "Nested",
			},
		},
	)

	encoder.EncodeToken(
		xml.EndElement{
			Name: xml.Name{
				Space: namespace,
				Local: "Hello",
			},
		},
	)

	encoder.Flush()

	got := b.String()
	expected := `<Hello xmlns="https://hello.world.com/1/0">
  <Nested>some &lt;&gt; text</Nested>
</Hello>`

	if expected != got {
		t.Fatalf("Expected VS got:\n%s\n%s", expected, got)
		return
	}
}

func TestAttrsWithNamespace(t *testing.T) {
	text := `<?xml version="1.0" encoding="UTF-8"?><something xmlns="https://admin-shell.io/aas/3/0">
	some <!-- comment --> text
	<childElement>asdf</childElement>
	<emptyElement />
</something>`

	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	writer := new(strings.Builder)

	for {
		token, tokenErr := decoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			} else {
				panic(tokenErr.Error())
			}
		}

		switch et := token.(type) {
		case xml.StartElement:
			writer.WriteString(
				fmt.Sprintf(
					"Start element: space: %s, local: %s\n",
					et.Name.Space, et.Name.Local,
				),
			)
			writer.WriteString(
				fmt.Sprintf(
					"  len(Attr): %d\n",
					len(et.Attr),
				),
			)
			for _, attr := range et.Attr {
				writer.WriteString(
					fmt.Sprintf(
						"  Attr: space: %s, local: %s\n",
						attr.Name.Space, attr.Name.Local,
					),
				)
			}
		default:
			// Ignore
		}
	}

	got := writer.String()

	expected := `Start element: space: https://admin-shell.io/aas/3/0, local: something
  len(Attr): 1
  Attr: space: , local: xmlns
Start element: space: https://admin-shell.io/aas/3/0, local: childElement
  len(Attr): 0
Start element: space: https://admin-shell.io/aas/3/0, local: emptyElement
  len(Attr): 0
`

	if got != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, got)
		return
	}

	return
}
