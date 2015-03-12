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
var freqCharSet = flag.String("freq-char-set", "Yjl", "")
var numItems = flag.Int("num-items", 10000, "")

func randRune(runeSet, freqRuneSet []rune) rune {
	chosen := rand.Intn(len(runeSet) * 2)
	if chosen < len(runeSet) {
		return runeSet[chosen]
	}
	chosen = (chosen - len(runeSet)) % len(freqRuneSet)
	return freqRuneSet[chosen]
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
	freqRuneSet := stringToRuneSlice(*freqCharSet)
	for i := 0; i < *numItems; i++ {
		length := rand.Intn(*maxLength-*minLength+1) + *minLength
		buffer := &bytes.Buffer{}
		for j := 0; j < length; j++ {
			buffer.WriteRune(randRune(runeSet, freqRuneSet))
		}
		fmt.Println(buffer.String())
	}
}
