package gotl

type Container interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
	SubContainer(start, end int) Container
}

// IntSlice
type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func (s IntSlice) SubContainer(start, end int) Container { return s[start:end] }


// Int64Slice
type Int64Slice []int64

func (s Int64Slice) Len() int { return len(s) }

func (s Int64Slice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Int64Slice) Less(i, j int) bool { return s[i] < s[j] }

func (s Int64Slice) SubContainer(start, end int) Container { return s[start:end] }


// Float32Slice
type Float32Slice []float32

func (s Float32Slice) Len() int { return len(s) }

func (s Float32Slice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Float32Slice) Less(i, j int) bool { return s[i] < s[j] }

func (s Float32Slice) SubContainer(start, end int) Container { return s[start:end] }


// Float32Slice
type Float64Slice []float64

func (s Float64Slice) Len() int { return len(s) }

func (s Float64Slice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Float64Slice) Less(i, j int) bool { return s[i] < s[j] }

func (s Float64Slice) SubContainer(start, end int) Container { return s[start:end] }


// StringSlice
type StringSlice []string

func (s StringSlice) Len() int { return len(s) }

func (s StringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }

func (s StringSlice) SubContainer(start, end int) Container { return s[start:end] }
