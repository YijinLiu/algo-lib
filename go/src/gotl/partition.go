package gotl

import (
	"log"
)

func Partition(c Container) int {
	size := c.Len()
	pivot := size - 1
	splitPos := -1
	for i := 0; i < size - 1; i++ {
		if c.Less(i, pivot) {
			splitPos++
			if i != splitPos {
				c.Swap(splitPos, i)
			}
		}
	}
	splitPos++
	if splitPos != pivot {
		c.Swap(splitPos, pivot)
	}
	return splitPos
}

// nth is 0-based
func NthElement(c Container, nth int) {
	if nth >= c.Len() {
		log.Fatalf("%d is beyond the container's size!", nth)
	}
	if nth == 0 && c.Len() == 1 {
		return
	}
	pivot := Partition(c)
	if pivot == nth {
		return
	}
	if pivot > nth {
		NthElement(c.SubContainer(0, pivot), nth)
	} else {
		NthElement(c.SubContainer(pivot + 1, c.Len()), nth - pivot - 1)
	}

}

func Median(c Container) {
	NthElement(c, c.Len() / 2)
}
