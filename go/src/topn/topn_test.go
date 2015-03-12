package topn

import . "testutils"

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestFrequentItemMonitor(t *testing.T) {
	m := NewFrequentItemMonitor(10)
	inFile, err := os.Open("testdata/items.txt")
	if err != nil {
		t.Error("Failed to open input testdata!")
	}
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		item := StringItem(strings.TrimSpace(scanner.Text()))
		m.Observe(item)
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}
	actual := make([]string, len(m.TopItems))
	for i, v := range m.SortedTopItems() {
		actual[i] = v.item.Key()
	}

	expected := []string{}
	inFile, err = os.Open("testdata/top_items.txt")
	if err != nil {
		t.Error("Failed to open output testdata!")
	}
	scanner = bufio.NewScanner(inFile)
	for scanner.Scan() {
		expected = append(expected, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	CheckStringSliceEq(expected, actual, t)
}
