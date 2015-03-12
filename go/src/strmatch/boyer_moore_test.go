package strmatch

import . "testutils"

import (
	"testing"
)

func TestBMSimple(t *testing.T) {
	m := NewBoyerMooreMatch("abc")
	CheckIntEq(3, m.FindFirst("deaabc"), t)
	CheckIntSliceEq([]int{2, 5}, m.FindAll("fgabcabc"), t)
}

func TestBMBadCharRule(t *testing.T) {
	m := NewBoyerMooreMatch("abc")
	CheckIntSliceEq([]int{3, 7}, m.FindAll("abdabceabc"), t)
}

func TestBMGoodPrefixRule(t *testing.T) {
	m := NewBoyerMooreMatch("abceabc")
	CheckIntSliceEq([]int{3}, m.FindAll("abdabceabc"), t)
}

func TestBMComplicated(t *testing.T) {
	m := NewBoyerMooreMatch("qcabdabdab")
	CheckIntSliceEq([]int{0, 0, 0, 0, 0, 6, 0, 0, 3, 0}, m.L_tilt, t)
	CheckIntSliceEq([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, m.l_tilt, t)
	CheckIntSliceEq([]int{}, m.FindAll("prstabstubabvqxrst"), t)
	CheckIntSliceEq([]int{8}, m.FindAll("abdabdabqcabdabdab"), t)
	// TODO: Add more tests.
}
