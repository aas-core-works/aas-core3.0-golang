package xmlization

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

func TestIteratorEmpty(t *testing.T) {
	text := ""
	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	iterator := newIterator(decoder)
	iterator.Start()

	if !iterator.IsEOF() {
		t.Fatalf(
			"Expected that the iterator reaches an EOF, " +
				"but its current token is: %T %v",
				iterator.Current(), iterator.Current(),
		)
	}
}

func TestIteratorStartTextEnd(t *testing.T) {
	text := "<something>some text</something>"
	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	iterator := newIterator(decoder)
	iterator.Start()

	buf := bytes.NewBufferString("")

	for {
		fmt.Fprintf(buf, "%T\n", iterator.Current())

		err := iterator.Next()
		if err != nil {
			t.Fatal(err)
			return
		}

		if iterator.IsEOF() {
			break
		}
	}

	expected := `xml.StartElement
xml.CharData
xml.EndElement
`

	if buf.String() != expected {
		t.Fatalf("Expected `%v`, got `%v`", expected, buf.String())
	}
}

func TestIteratorCommentInText(t *testing.T) {
	text := "<something>some <!-- comment in between --> text</something>"
	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	iterator := newIterator(decoder)
	iterator.Start()

	buf := bytes.NewBufferString("")

	for {
		fmt.Fprintf(buf, "%T\n", iterator.Current())

		err := iterator.Next()
		if err != nil {
			t.Fatal(err)
			return
		}

		if iterator.IsEOF() {
			break
		}
	}

	expected := `xml.StartElement
xml.CharData
xml.CharData
xml.EndElement
`

	if buf.String() != expected {
		t.Fatalf("Expected `%v`, got `%v`", expected, buf.String())
	}
}

func TestIteratorSkipWhitespaceEmpty(t *testing.T) {
	text := ""
	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	iterator := newIterator(decoder)
	iterator.Start()
	iterator.SkipWhitespaceIfAny()

	if !iterator.IsEOF() {
		t.Fatalf(
			"Expected that the iterator reaches an EOF, " +
				"but its current token is: %T %v",
				iterator.Current(), iterator.Current(),
		)
	}
}

func TestIteratorSkipWhitespaceNested(t *testing.T) {
	text := `
<something>
	<nested />
</something>`
	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	iterator := newIterator(decoder)
	iterator.Start()

	buf := bytes.NewBufferString("")

	for {
		err := iterator.SkipWhitespaceIfAny()
		if err != nil {
			t.Fatal(err)
			return
		}

		fmt.Fprintf(buf, "%T\n", iterator.Current())

		err = iterator.Next()
		if err != nil {
			t.Fatal(err)
			return
		}

		if iterator.IsEOF() {
			break
		}
	}

	expected := `xml.StartElement
xml.StartElement
xml.EndElement
xml.EndElement
`

	if buf.String() != expected {
		t.Fatalf("Expected `%v`, got `%v`", expected, buf.String())
	}
}

func TestReadAndMergeTextWithCommentInBetween(t *testing.T) {
	text := "<something>some <!-- comment in between --> text</something>"
	reader := strings.NewReader(text)
	decoder := xml.NewDecoder(reader)

	iterator := newIterator(decoder)
	iterator.Start()
	err := iterator.Next()
	if err != nil {
		t.Fatal(err)
		return
	}

	text, err = readAndMergeText(iterator)
	if err != nil {
		t.Fatal(err)
		return
	}

	expected := "some  text"

	if expected != text {
		t.Fatalf("Expected `%v`, got `%v`", expected, text)
		return
	}
}