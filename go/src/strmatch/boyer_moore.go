package strmatch

import (
	"logging"
)

type BoyerMooreMatch struct {
	P string

	// For definition, see section 2.2.2 in ASTS.
	badCharPositions map[byte][]int
	// For definition of L_tilt and l_tilt, see section 2.2.4 in ASTS.
	L_tilt []int
	l_tilt []int
}

func computeBadCharPositions(P string) map[byte][]int {
	result := make(map[byte][]int)
	for i := len(P) - 1; i >= 0; i-- {
		ch := P[i]
		if positions, found := result[ch]; found {
			positions = append(positions, i)
		} else {
			result[ch] = []int{i}
		}
	}
	return result
}

// For definition of Nj, see section 2.2.4 in ASTS.
func computeNj(P string) []int {
	n := len(P)
	Nj := make([]int, n)
	// [l, r] is the current Z box.
	l := n
	r := n - 1
	for i := n - 2; i >= 0; i-- {
		if i < l {
			j := i
			for j >= 0 && P[n-1-i+j] == P[j] {
				j--
			}
			if j < i {
				l = j + 1
				r = i
				Nj[i] = i - j
			}
		} else {
			belta := Nj[n-1-r+i]
			if belta <= i-l {
				Nj[i] = belta
			} else {
				j := l - 1
				for ; j >= 0 && P[n-belta-i+j] == P[j]; j-- {
				}
				l = j + 1
				r = i
				Nj[i] = i - j
			}
		}
	}
	return Nj
}

func NewBoyerMooreMatch(P string) *BoyerMooreMatch {
	n := len(P)
	if n < 2 {
		logging.Fatalf("Pattern string (%q) is too small!", P)
	}
	Nj := computeNj(P)
	L_tilt := make([]int, n)
	l_tilt := make([]int, n)
	longestPreSuffix := 0
	for i := 0; i < n-1; i++ {
		if Nj[i] > 0 {
			j := n - Nj[i]
			L_tilt[j] = i
		}
		if Nj[i] == i+1 {
			longestPreSuffix = i + 1
		}
		l_tilt[n-1-i] = longestPreSuffix
	}
	return &BoyerMooreMatch{P, computeBadCharPositions(P), L_tilt, l_tilt}
}

func (m *BoyerMooreMatch) FindFirst(T string) int {
	positions := m.findInternal(T, true)
	if len(positions) > 0 {
		return positions[0]
	}
	return -1
}

func (m *BoyerMooreMatch) FindAll(T string) []int {
	return m.findInternal(T, false)
}

func (m *BoyerMooreMatch) getShiftByBadChar(ch byte, mismatchPos int) int {
	if positions, found := m.badCharPositions[ch]; found {
		for _, pos := range positions {
			if pos < mismatchPos {
				return mismatchPos - pos
			}
		}
	}
	return mismatchPos + 1
}

func (m *BoyerMooreMatch) getShiftByGoodSuffix(lastMatchPos int) int {
	L_tilt := m.L_tilt[lastMatchPos]
	if L_tilt > 0 {
		return len(m.P) - 1 - L_tilt
	}
	return len(m.P) - 1 - m.l_tilt[lastMatchPos]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func (m *BoyerMooreMatch) findInternal(T string, returnFirst bool) []int {
	result := []int{}

	for k := len(m.P) - 1; k < len(T); {
		i := len(m.P) - 1
		j := k
		for i >= 0 && m.P[i] == T[j] {
			i--
			j--
		}
		if i < 0 {
			result = append(result, j+1)
			if returnFirst {
				return result
			}
			k += len(m.P) - 1 + m.l_tilt[1]
		} else if i == len(m.P)-1 {
			k++
		} else {
			k += max(m.getShiftByBadChar(T[j], i), m.getShiftByGoodSuffix(i+1))
		}
	}

	return result
}
