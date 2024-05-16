package xmlization

import (
	"encoding/xml"
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
		t.Fatalf("Expected that the iterator reaches an EOF, but it has not.")
	}
}
