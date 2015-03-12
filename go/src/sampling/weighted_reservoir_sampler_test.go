package sampling

import . "testutils"

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestWeightedReservoirSampler(t *testing.T) {
	s := NewWeightedReservoirSamplerWithRand(10, rand.New(rand.NewSource(1)))
	inFile, err := os.Open("testdata/items.txt")
	if err != nil {
		t.Error("Failed to open input testdata!")
	}
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		items := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		CheckIntEq(2, len(items), t)
		w, err := strconv.ParseFloat(items[1], 64)
		if err != nil {
			t.Errorf("Invalid weight for '%s': %s!", items[0], items[1])
		}
		s.AddWithWeight(items[0], w)
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}
	actual := make([]string, len(s.Samples()))
	for i, v := range s.Samples() {
		actual[i] = v.(string)
	}

	expected := []string{}
	inFile, err = os.Open("testdata/sampled_items.txt")
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
