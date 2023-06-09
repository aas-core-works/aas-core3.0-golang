package verification

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	leapYears := []*big.Int{
		big.NewInt(2016), big.NewInt(1600), big.NewInt(2000),
		big.NewInt(-1), big.NewInt(-5),
	}
	notLeapYears := []*big.Int{
		big.NewInt(1700), big.NewInt(1800), big.NewInt(-4),
	}

	for _, year := range leapYears {
		if !isLeapYear(year) {
			t.Fatalf(
				"Expected year %d to be a leap year, but it was reported as not.",
				year,
			)
		}
	}

	for _, year := range notLeapYears {
		if isLeapYear(year) {
			t.Fatalf(
				"Expected year %d not to be a leap year, but it was reported as one.",
				year,
			)
		}
	}
}

// NOTE (mristin, 2023-06-08):
// We explicitly test the implementation in Go here merely for easier debugging.
// There are many more cases in the generated testdata.

func TestIsXsByte(t *testing.T) {
	value := "1"

	if !MatchesXsByte(value) {
		t.Fatalf(
			"Expected the value %v to match a xs:byte, but it was reported as not.",
			value,
		)
	}

	if !IsXsByte(value) {
		t.Fatalf(
			"Expected the value %v to be a xs:byte, but it was reported as not.",
			value,
		)
	}
}

func TestIsXsDateOnAVeryLargeYear(t *testing.T) {
	// NOTE (mristin, 2023-06-08):
	// We handle years as 64-bit integers, so this year will overflow.
	value := "12345678901234567890123456789012345678901234567890-12-12"

	if !MatchesXsDate(value) {
		t.Fatalf(
			"Expected the value %v to match a xs:date, but it was reported as not.",
			value,
		)
	}

	if !IsXsDate(value) {
		t.Fatalf(
			"Expected the value %v to be accepted as xs:date, but it was not.",
			value,
		)
	}
}

func TestIsXsUnsignedByte(t *testing.T) {
	value := "+1"

	if !MatchesXsUnsignedByte(value) {
		t.Fatalf(
			"Expected the value %v to match a xs:unsignedByte, "+
				"but it was reported as not.",
			value,
		)
	}

	_, err := strconv.ParseUint(value, 10, 64)
	fmt.Printf("err: %v", err)

	if !IsXsUnsignedByte(value) {
		t.Fatalf(
			"Expected the value %v to be a xs:unsignedByte, "+
				"but it was reported as not.",
			value,
		)
	}
}
