package strmatch

import (
	"logging"
)

type KnuthMorrisPrattMatch struct {
	P string

	// For definition of sp_tilt, see section 2.3.1 of book ATST.
	sp_tilt []int
}

// For definition of Zj, see section 1.3 of book ATST.
func computeZj(P string) []int {
	n := len(P)
	Zj := make([]int, n)
	// [l, r] is the current Z box.
	l := 0
	r := -1
	for i := 1; i < n; i++ {
		if i > r {
			j := i
			for j < n && P[j-i] == P[j] {
				j++
			}
			if i < j {
				l = i
				r = j - 1
				Zj[i] = j - i
			}
		} else {
			beta := Zj[i-l]
			if beta <= r-i {
				Zj[i] = beta
			} else {
				j := r + 1
				for j < n && P[beta+j-r] == P[j] {
					j++
				}
				l = i
				r = j - 1
				Zj[i] = j - i
			}
		}
	}
	return Zj
}

func NewKnuthMorrisPrattMatch(P string) *KnuthMorrisPrattMatch {
	n := len(P)
	if n < 2 {
		logging.Fatalf("Pattern string (%q) is too small!", P)
	}
	Zj := computeZj(P)
	sp_tilt := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if Zj[i] > 0 {
			j := i + Zj[i] - 1
			sp_tilt[j] = Zj[i]
		}
	}
	return &KnuthMorrisPrattMatch{P, sp_tilt}
}

func (m *KnuthMorrisPrattMatch) FindFirst(T string) int {
	positions := m.findInternal(T, true)
	if len(positions) > 0 {
		return positions[0]
	}
	return -1
}

func (m *KnuthMorrisPrattMatch) FindAll(T string) []int {
	return m.findInternal(T, false)
}

func (m *KnuthMorrisPrattMatch) findInternal(T string, returnFirst bool) []int {
	result := []int{}

	n := len(m.P)
	c := 0
	p := 0
	for c+n-p <= len(T) {
		for p < n && m.P[p] == T[c] {
			c++
			p++
		}
		if p == n {
			result = append(result, c-n)
			if returnFirst {
				return result
			}
		}

		if p == 0 {
			c++
		} else {
			p = m.sp_tilt[p-1]
		}
	}

	return result
}
