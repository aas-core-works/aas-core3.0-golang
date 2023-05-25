package reporting_test

import (
	aasreporting "github.com/aas-core-works/aas-core3.0-golang/reporting"
	"testing"
)

func TestToJSONPathOnEmpty(t *testing.T) {
	p := &aasreporting.Path{}
	got := aasreporting.ToJSONPath(p)
	expected := ""
	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestToJSONPathOnOne(t *testing.T) {
	p := &aasreporting.Path{}
	p.PrependName(&aasreporting.NameSegment{Name: "a"})

	got := aasreporting.ToJSONPath(p)
	expected := "a"
	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestToJSONPathOnTwo(t *testing.T) {
	p := &aasreporting.Path{}
	p.PrependName(&aasreporting.NameSegment{Name: "b"})
	p.PrependName(&aasreporting.NameSegment{Name: "a"})

	got := aasreporting.ToJSONPath(p)
	expected := "a.b"
	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestToJSONPathOnThree(t *testing.T) {
	p := &aasreporting.Path{}
	p.PrependName(&aasreporting.NameSegment{Name: "c"})
	p.PrependName(&aasreporting.NameSegment{Name: "b"})
	p.PrependName(&aasreporting.NameSegment{Name: "a"})

	got := aasreporting.ToJSONPath(p)
	expected := "a.b.c"
	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}
