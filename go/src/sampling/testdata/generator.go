package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"unicode/utf8"
)

var minLength = flag.Int("min-len", 2, "")
var maxLength = flag.Int("max-len", 3, "")
var charSet = flag.String("char-set", "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
var favorCharSet = flag.String("favor-char-set", "Yjl", "")
var numItems = flag.Int("num-items", 10000, "")

func isIn(r rune, runeSet []rune) bool {
	for _, v := range runeSet {
		if v == r {
			return true
		}
	}
	return false
}

func randRune(runeSet, favorRuneSet []rune) (r rune, w float64) {
	chosen := rand.Intn(len(runeSet))
	r = runeSet[chosen]
	w = rand.Float64() / 1000
	if isIn(r, favorRuneSet) {
		w += 0.999
	}
	return
}

func stringToRuneSlice(s string) []rune {
	runeSlice := make([]rune, utf8.RuneCountInString(s))
	for i, r := range s {
		runeSlice[i] = r
	}
	return runeSlice
}

func main() {
	runeSet := stringToRuneSlice(*charSet)
	favorRuneSet := stringToRuneSlice(*favorCharSet)
	for i := 0; i < *numItems; i++ {
		length := rand.Intn(*maxLength - *minLength + 1) + *minLength
		buffer := &bytes.Buffer{}
		weight := 0.0
		for j := 0; j < length; j++ {
			r, w := randRune(runeSet, favorRuneSet)
			buffer.WriteRune(r)
			weight += w
		}
		fmt.Printf("%s %f\n", buffer.String(), weight / float64(length))
	}
}
