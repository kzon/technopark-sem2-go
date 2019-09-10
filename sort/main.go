package main

import (
	"flag"
	"fmt"
)

var flagReverse = flag.Bool("r", false, "sort descending")
var flagCaseInsensitive = flag.Bool("f", false, "case insensitive")
var flagUnique = flag.Bool("u", false, "output only unique values")
var flagOutputFile = flag.String("o", "", "output file path")
var flagColumn = flag.Int("k", 0, "column number to sort by column")
var flagNumeric = flag.Bool("n", false, "numeric sort (ignoring blanks in the beginning)")

func main() {
	flag.Parse()

	lines, err := readLines()
	if err != nil {
		fmt.Println("Error reading input\n", err)
		return
	}
	sortedLines := sortLines(lines)
	err = outputLines(sortedLines, *flagOutputFile)
	if err != nil {
		fmt.Println("Error writing output\n", err)
	}
}

func sortLines(lines []string) []string {
	sorter := Sorter{
		lines,
		&SorterSettings{
			Reverse:         *flagReverse,
			CaseInsensitive: *flagCaseInsensitive,
			Unique:          *flagUnique,
			Column:          *flagColumn,
			Numeric:         *flagNumeric,
		},
	}
	return sorter.GetSorted()
}
