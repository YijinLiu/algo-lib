package sampling

import . "testutils"

import (
	"hash/fnv"
	"io"
	"math/rand"
	"testing"
	"unicode/utf8"

	"logging"
)

func TestReservoirSampler(t *testing.T) {
	sampler := NewReservoirSamplerWithRand(10, rand.New(rand.NewSource(1)))
	for i := 0; i < 50; i++ {
		sampler.Add("a")
	}
	for i := 0; i < 30; i++ {
		sampler.Add("b")
	}
	for i := 0; i < 10; i++ {
		sampler.Add("c")
	}
	for i := 0; i < 10; i++ {
		sampler.Add("d")
	}

	actual := make([]string, len(sampler.Samples()))
	for index, item := range sampler.Samples() {
		actual[index] = item.(string)
	}

	expected := []string{"a", "a", "b", "d", "a", "b", "a", "c", "a", "b"}
	CheckStringSliceEq(expected, actual, t)
}

func TestHashBasedReservoirSampler(t *testing.T) {
	hashGen := fnv.New64a()
	sampler := NewHashBasedReservoirSampler(10)
	for i := 0; i < 500; i++ {
		hashGen.Reset()
		if _, err := io.WriteString(hashGen, "a"); err != nil {
			logging.Fatalf("Failed to write key: %s!", err)
		}
		sampler.AddWithHash("a", hashGen.Sum64())
	}
	for i := 0; i < 300; i++ {
		hashGen.Reset()
		if _, err := io.WriteString(hashGen, "b"); err != nil {
			logging.Fatalf("Failed to write key: %s!", err)
		}
		sampler.AddWithHash("b", hashGen.Sum64())
	}
	for i := 0; i < 100; i++ {
		hashGen.Reset()
		if _, err := io.WriteString(hashGen, "c"); err != nil {
			logging.Fatalf("Failed to write key: %s!", err)
		}
		sampler.AddWithHash("c", hashGen.Sum64())
	}
	for i := 0; i < 100; i++ {
		hashGen.Reset()
		if _, err := io.WriteString(hashGen, "d"); err != nil {
			logging.Fatalf("Failed to write key: %s!", err)
		}
		sampler.AddWithHash("d", hashGen.Sum64())
	}

	actual := make([]string, len(sampler.Samples()))
	for index, item := range sampler.Samples() {
		actual[index] = item.(string)
	}

	expected := []string{"c", "d", "b", "a", "b", "a", "a", "a", "a", "d"}
	CheckStringSliceEq(expected, actual, t)
}

func randRune(runeSet []rune, randGen *rand.Rand) (r rune) {
	chosen := randGen.Intn(len(runeSet))
	return runeSet[chosen]
}

func stringToRuneSlice(s string) []rune {
	runeSlice := make([]rune, utf8.RuneCountInString(s))
	for i, r := range s {
		runeSlice[i] = r
	}
	return runeSlice
}

func TestFastForwardReservoirSampler(t *testing.T) {
	randGen := rand.New(rand.NewSource(1))
	sampler := NewFastForwardReservoirSamplerWithRand(10, randGen)
	runeSet := stringToRuneSlice("abcd")
	var i int64 = 0
	for ; i < 1000; i++ {
		skips := sampler.GetSkips()
		if skips < 0 {
			break
		}
		i += skips
		if i >= 10000 {
			break
		}
		sampler.Add(string(randRune(runeSet, randGen)))
	}

	actual := make([]string, len(sampler.Samples()))
	for index, item := range sampler.Samples() {
		actual[index] = item.(string)
	}

	expected := []string{"c", "b", "a", "a", "d", "b", "c", "a", "b", "b"}
	CheckStringSliceEq(expected, actual, t)
}
