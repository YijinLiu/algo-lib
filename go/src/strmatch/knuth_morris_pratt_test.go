package strmatch

import . "testutils"

import (
	"testing"
)

func TestKMPSimple(t *testing.T) {
	m := NewKnuthMorrisPrattMatch("abc")
	CheckIntSliceEq([]int{0, 0, 0}, m.sp_tilt, t)
	CheckIntEq(0, m.FindFirst("abc"), t)
	CheckIntSliceEq([]int{0, 3}, m.FindAll("abcabc"), t)
}

func TestKMPComplicated(t *testing.T) {
	m := NewKnuthMorrisPrattMatch("abcxabcde")
	CheckIntSliceEq([]int{0, 0, 0, 0, 0, 0, 3, 0, 0}, m.sp_tilt, t)
	CheckIntEq(6, m.FindFirst("xyabcxabcxabcdefg"), t)
}
