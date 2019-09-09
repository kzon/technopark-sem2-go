package main

import (
	"sort"
	"strconv"
	"strings"
)

type Sorter struct {
	data []string
	*SorterSettings
}

type SorterSettings struct {
	Reverse         bool
	CaseInsensitive bool
	Unique          bool
	Column          int
	Numeric         bool
}

func (s Sorter) GetSorted() []string {
	sort.Sort(s)
	if s.Unique {
		s.filterUnique()
	}
	return s.data
}

func (s Sorter) Len() int { return len(s.data) }

func (s Sorter) Less(i, j int) bool {
	value1 := s.getUniqueValue(s.data[i])
	value2 := s.getUniqueValue(s.data[j])
	if s.Reverse {
		return s.less(value2, value1)
	}
	return s.less(value1, value2)
}

func (s Sorter) Swap(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] }

func (s *Sorter) filterUnique() {
	isUnique := make(map[string]bool, len(s.data))
	var uniqueValues []string
	for _, value := range s.data {
		preparedValue := s.getUniqueValue(value)
		if _, ok := isUnique[preparedValue]; !ok {
			isUnique[preparedValue] = true
			uniqueValues = append(uniqueValues, value)
		}
	}
	s.data = uniqueValues
}

func (s Sorter) getUniqueValue(value string) string {
	if s.CaseInsensitive {
		value = strings.ToLower(value)
	}
	if s.Column != 0 {
		columns := strings.Split(value, " ")
		if s.Column <= len(columns) {
			value = columns[s.Column-1]
		}
	}
	if s.Numeric {
		value = strings.TrimSpace(value)
	}
	return value
}

func (s Sorter) less(value1, value2 string) bool {
	if !s.Numeric {
		return value1 < value2
	}
	numericValue1, err1 := strconv.ParseFloat(value1, 64)
	numericValue2, err2 := strconv.ParseFloat(value2, 64)
	if err1 == nil && err2 == nil {
		return numericValue1 < numericValue2
	}
	if err1 != nil && err2 != nil {
		return value1 < value2
	}
	// If one of values is non-numeric, it is always less then numeric one.
	return err1 != nil
}
