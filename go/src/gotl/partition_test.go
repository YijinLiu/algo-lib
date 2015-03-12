package gotl

import . "testutils"

import (
	"testing"
)

func TestPartition(t *testing.T) {
	inArray := [...]int{6, 3, 7, 4, 1, 2, 5, 3, -3, 0, 5, 2, 3}
	actual := inArray[:]
	CheckIntEq(5, Partition(IntSlice(actual)), t)
	outArray := [...]int{1, 2, -3, 0, 2, 3, 5, 3, 7, 4, 5, 6, 3}
	expected := outArray[:]
	CheckIntSliceEq(expected, actual, t)
}

func TestNthElement(t *testing.T) {
	inArray := [...]int{8, 1, -3, -2, 7, 4, -1, 0, 5, 6, -4, 9, 4}
	var actual []int
	var expected []int

	// min
	actual = inArray[:]
	NthElement(IntSlice(actual), 0)
	CheckIntEq(-4, actual[0], t)
	outArray1 := [...]int{-4, -3, -2, -1, 0, 1, 4, 7, 5, 6, 4, 9, 8}
	expected = outArray1[:]
	CheckIntSliceEq(expected, actual, t)

	// max
	actual = inArray[:]
	NthElement(IntSlice(actual), 12)
	CheckIntEq(9, actual[12], t)
	outArray2 := [...]int{-4, -3, -2, -1, 0, 1, 4, 7, 5, 6, 4, 8, 9}
	expected = outArray2[:]
	CheckIntSliceEq(expected, actual, t)

	// middle
	actual = inArray[:]
	Median(IntSlice(actual))
	CheckIntEq(4, actual[len(actual)/2], t)
	outArray3 := [...]int{-4, -3, -2, -1, 0, 1, 4, 7, 5, 6, 4, 8, 9}
	expected = outArray3[:]
	CheckIntSliceEq(expected, actual, t)

	actual = inArray[:]
	NthElement(IntSlice(actual), 3)
	CheckIntEq(-1, actual[3], t)
	outArray4 := [...]int{-4, -3, -2, -1, 0, 1, 4, 7, 5, 6, 4, 8, 9}
	expected = outArray4[:]
	CheckIntSliceEq(expected, actual, t)

	actual = inArray[:]
	NthElement(IntSlice(actual), 10)
	CheckIntEq(7, actual[10], t)
	outArray5 := [...]int{-4, -3, -2, -1, 0, 1, 4, 4, 5, 6, 7, 8, 9}
	expected = outArray5[:]
	CheckIntSliceEq(expected, actual, t)
}
