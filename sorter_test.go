package main

import (
	"testing"
)

func TestCasualSort(t *testing.T) {
	s := sorterSuite{t, &SorterSettings{}}
	s.Expect(
		[]string{},
		[]string{},
	)
	s.Expect(
		[]string{"foo"},
		[]string{"foo"},
	)
	s.Expect(
		[]string{"3", "2", "zero"},
		[]string{"2", "3", "zero"},
	)
	s.Expect(
		[]string{"3", "2", "zero"},
		[]string{"2", "3", "zero"},
	)
	s.Expect(
		[]string{"foo", "FOO"},
		[]string{"FOO", "foo"},
	)
	s.Expect(
		[]string{"foo", "foo", "bar"},
		[]string{"bar", "foo", "foo"},
	)
}

func TestReverseSort(t *testing.T) {
	s := sorterSuite{t, &SorterSettings{Reverse: true}}
	s.Expect(
		[]string{"3", "2", "1", "one"},
		[]string{"one", "3", "2", "1"},
	)
}

func TestCaseInsensitiveSort(t *testing.T) {
	s := sorterSuite{t, &SorterSettings{CaseInsensitive: true}}
	s.Expect(
		[]string{"foo", "FOO"},
		[]string{"foo", "FOO"},
	)
}

func TestUniqueSort(t *testing.T) {
	uniqueSuite := sorterSuite{t, &SorterSettings{Unique: true}}
	uniqueSuite.Expect(
		[]string{"foo", "foo", "bar"},
		[]string{"bar", "foo"},
	)

	uniqueCaseInsensitiveSuite := sorterSuite{t, &SorterSettings{CaseInsensitive: true, Unique: true}}
	uniqueCaseInsensitiveSuite.Expect(
		[]string{"foo", "Foo", "bar"},
		[]string{"bar", "foo"},
	)
	uniqueCaseInsensitiveSuite.Expect(
		[]string{"Foo", "foo", "bar"},
		[]string{"bar", "Foo"},
	)
}

func TestColumnSort(t *testing.T) {
	firstColumnSuite := sorterSuite{t, &SorterSettings{Column: 1}}
	firstColumnSuite.Expect(
		[]string{"2 Alice", "1 Bob", "3 Evan"},
		[]string{"1 Bob", "2 Alice", "3 Evan"},
	)

	secondColumnSuite := sorterSuite{t, &SorterSettings{Column: 2}}
	secondColumnSuite.Expect(
		[]string{"1 Bob", "2 Alice", "3 Evan"},
		[]string{"2 Alice", "1 Bob", "3 Evan"},
	)
	secondColumnSuite.Expect(
		[]string{"Bob", "Alice", "Evan"},
		[]string{"Alice", "Bob", "Evan"},
	)

	secondColumnUniqueSuite := sorterSuite{t, &SorterSettings{Column: 2, Unique: true}}
	secondColumnUniqueSuite.Expect(
		[]string{"1 Bob", "2 Alice", "3 Evan", "4 Evan"},
		[]string{"2 Alice", "1 Bob", "3 Evan"},
	)
}

func TestNumericSort(t *testing.T) {
	nonNumericSuite := sorterSuite{t, &SorterSettings{Numeric: false}}
	nonNumericSuite.Expect(
		[]string{"12", "   24", "543.5", "-3", "2", "60.35"},
		[]string{"   24", "-3", "12", "2", "543.5", "60.35"},
	)

	numericSuite := sorterSuite{t, &SorterSettings{Numeric: true}}
	numericSuite.Expect(
		[]string{"12", "   24", "543.5", "-3", "2", "60.35"},
		[]string{"-3", "2", "12", "   24", "60.35", "543.5"},
	)
	numericSuite.Expect(
		[]string{"1", "d", "a", "12"},
		[]string{"a", "d", "1", "12"},
	)

	numericUniqueSuite := sorterSuite{t, &SorterSettings{Numeric: true, Unique: true}}
	numericUniqueSuite.Expect(
		[]string{"   12   ", "12", "100", "-100"},
		[]string{"-100", "   12   ", "100"},
	)
}

type sorterSuite struct {
	t *testing.T
	*SorterSettings
}

func (s sorterSuite) Expect(input []string, expected []string) {
	sorter := Sorter{input, s.SorterSettings}
	result := sorter.GetSorted()
	if !s.equal(result, expected) {
		s.t.Errorf("\nExpected: %v\nActual:   %v\n", expected, result)
	}
}

func (s sorterSuite) equal(first []string, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}
