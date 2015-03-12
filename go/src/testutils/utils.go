package testutils

import (
	"bytes"
	"fmt"
	"path"
	"runtime"
	"testing"
)

func failTest(message string, t *testing.T) {
	if _, file, line, ok := runtime.Caller(2); ok {
		t.Errorf("[%s:%d] %s", path.Base(file), line, message)
	} else {
		t.Error(message)
	}
}

// For int
func CheckIntEq(expected, actual int, t *testing.T) {
	if expected != actual {
		failTest(fmt.Sprintf("Expected value is %d, got %d.", expected, actual), t)
	}
}

func CheckIntNe(expected, actual int, t *testing.T) {
	if expected == actual {
		failTest(fmt.Sprintf("Expected value != %d, got %d.", expected, actual), t)
	}
}

func CheckIntLt(expected, actual int, t *testing.T) {
	if expected >= actual {
		failTest(fmt.Sprintf("Expected value < %d, got %d.", expected, actual), t)
	}
}

func CheckIntGt(expected, actual int, t *testing.T) {
	if expected <= actual {
		failTest(fmt.Sprintf("Expected value > %d, got %d.", expected, actual), t)
	}
}

func CheckIntLe(expected, actual int, t *testing.T) {
	if expected > actual {
		failTest(fmt.Sprintf("Expected value <= %d, got %d.", expected, actual), t)
	}
}

func CheckIntGe(expected, actual int, t *testing.T) {
	if expected < actual {
		failTest(fmt.Sprintf("Expected value >= %d, got %d.", expected, actual), t)
	}
}

// For int32
func CheckInt32Eq(expected, actual int32, t *testing.T) {
	if expected != actual {
		failTest(fmt.Sprintf("Expected value is %d, got %d.", expected, actual), t)
	}
}

func CheckInt32Ne(expected, actual int32, t *testing.T) {
	if expected == actual {
		failTest(fmt.Sprintf("Expected value != %d, got %d.", expected, actual), t)
	}
}

func CheckInt32Lt(expected, actual int32, t *testing.T) {
	if expected >= actual {
		failTest(fmt.Sprintf("Expected value < %d, got %d.", expected, actual), t)
	}
}

func CheckInt32Gt(expected, actual int32, t *testing.T) {
	if expected <= actual {
		failTest(fmt.Sprintf("Expected value > %d, got %d.", expected, actual), t)
	}
}

func CheckInt32Le(expected, actual int32, t *testing.T) {
	if expected > actual {
		failTest(fmt.Sprintf("Expected value <= %d, got %d.", expected, actual), t)
	}
}

func CheckInt32Ge(expected, actual int32, t *testing.T) {
	if expected < actual {
		failTest(fmt.Sprintf("Expected value >= %d, got %d.", expected, actual), t)
	}
}

// For uint32
func CheckUint32Eq(expected, actual uint32, t *testing.T) {
	if expected != actual {
		failTest(fmt.Sprintf("Expected value is %d, got %d.", expected, actual), t)
	}
}

func CheckUint32Ne(expected, actual uint32, t *testing.T) {
	if expected == actual {
		failTest(fmt.Sprintf("Expected value != %d, got %d.", expected, actual), t)
	}
}

func CheckUint32Lt(expected, actual uint32, t *testing.T) {
	if expected >= actual {
		failTest(fmt.Sprintf("Expected value < %d, got %d.", expected, actual), t)
	}
}

func CheckUint32Gt(expected, actual uint32, t *testing.T) {
	if expected <= actual {
		failTest(fmt.Sprintf("Expected value > %d, got %d.", expected, actual), t)
	}
}

func CheckUint32Le(expected, actual uint32, t *testing.T) {
	if expected > actual {
		failTest(fmt.Sprintf("Expected value <= %d, got %d.", expected, actual), t)
	}
}

func CheckUint32Ge(expected, actual uint32, t *testing.T) {
	if expected < actual {
		failTest(fmt.Sprintf("Expected value >= %d, got %d.", expected, actual), t)
	}
}

// For int64
func CheckInt64Eq(expected, actual int64, t *testing.T) {
	if expected != actual {
		failTest(fmt.Sprintf("Expected value is %d, got %d.", expected, actual), t)
	}
}

func CheckInt64Ne(expected, actual int64, t *testing.T) {
	if expected == actual {
		failTest(fmt.Sprintf("Expected value != %d, got %d.", expected, actual), t)
	}
}

func CheckInt64Lt(expected, actual int64, t *testing.T) {
	if expected >= actual {
		failTest(fmt.Sprintf("Expected value < %d, got %d.", expected, actual), t)
	}
}

func CheckInt64Gt(expected, actual int64, t *testing.T) {
	if expected <= actual {
		failTest(fmt.Sprintf("Expected value > %d, got %d.", expected, actual), t)
	}
}

func CheckInt64Le(expected, actual int64, t *testing.T) {
	if expected > actual {
		failTest(fmt.Sprintf("Expected value <= %d, got %d.", expected, actual), t)
	}
}

func CheckInt64Ge(expected, actual int64, t *testing.T) {
	if expected < actual {
		failTest(fmt.Sprintf("Expected value >= %d, got %d.", expected, actual), t)
	}
}

// For uint64
func CheckUint64Eq(expected, actual uint64, t *testing.T) {
	if expected != actual {
		failTest(fmt.Sprintf("Expected value is %d, got %d.", expected, actual), t)
	}
}

func CheckUint64Ne(expected, actual uint64, t *testing.T) {
	if expected == actual {
		failTest(fmt.Sprintf("Expected value != %d, got %d.", expected, actual), t)
	}
}

func CheckUint64Lt(expected, actual uint64, t *testing.T) {
	if expected >= actual {
		failTest(fmt.Sprintf("Expected value < %d, got %d.", expected, actual), t)
	}
}

func CheckUint64Gt(expected, actual uint64, t *testing.T) {
	if expected <= actual {
		failTest(fmt.Sprintf("Expected value > %d, got %d.", expected, actual), t)
	}
}

func CheckUint64Le(expected, actual uint64, t *testing.T) {
	if expected > actual {
		failTest(fmt.Sprintf("Expected value <= %d, got %d.", expected, actual), t)
	}
}

func CheckUint64Ge(expected, actual uint64, t *testing.T) {
	if expected < actual {
		failTest(fmt.Sprintf("Expected value >= %d, got %d.", expected, actual), t)
	}
}

// For float32
func CheckFloat32Eq(expected, actual float32, t *testing.T) {
	if expected != actual {
		failTest(fmt.Sprintf("Expected value is %f, got %f.", expected, actual), t)
	}
}

func CheckFloat32Ne(expected, actual float32, t *testing.T) {
	if expected == actual {
		failTest(fmt.Sprintf("Expected value != %f, got %f.", expected, actual), t)
	}
}

func CheckFloat32Lt(expected, actual float32, t *testing.T) {
	if expected >= actual {
		failTest(fmt.Sprintf("Expected value < %f, got %f.", expected, actual), t)
	}
}

func CheckFloat32Gt(expected, actual float32, t *testing.T) {
	if expected <= actual {
		failTest(fmt.Sprintf("Expected value > %f, got %f.", expected, actual), t)
	}
}

func CheckFloat32Le(expected, actual float32, t *testing.T) {
	if expected > actual {
		failTest(fmt.Sprintf("Expected value <= %f, got %f.", expected, actual), t)
	}
}

func CheckFloat32Ge(expected, actual float32, t *testing.T) {
	if expected < actual {
		failTest(fmt.Sprintf("Expected value >= %f, got %f.", expected, actual), t)
	}
}

// For float64
func CheckFloat64Eq(expected, actual float64, t *testing.T) {
	if expected != actual {
		failTest(fmt.Sprintf("Expected value is %f, got %f.", expected, actual), t)
	}
}

func CheckFloat64Ne(expected, actual float64, t *testing.T) {
	if expected == actual {
		failTest(fmt.Sprintf("Expected value != %f, got %f.", expected, actual), t)
	}
}

func CheckFloat64Lt(expected, actual float64, t *testing.T) {
	if expected >= actual {
		failTest(fmt.Sprintf("Expected value < %f, got %f.", expected, actual), t)
	}
}

func CheckFloat64Gt(expected, actual float64, t *testing.T) {
	if expected <= actual {
		failTest(fmt.Sprintf("Expected value > %f, got %f.", expected, actual), t)
	}
}

func CheckFloat64Le(expected, actual float64, t *testing.T) {
	if expected > actual {
		failTest(fmt.Sprintf("Expected value <= %f, got %f.", expected, actual), t)
	}
}

func CheckFloat64Ge(expected, actual float64, t *testing.T) {
	if expected < actual {
		failTest(fmt.Sprintf("Expected value >= %f, got %f.", expected, actual), t)
	}
}

// For string
func CheckStringEq(expected, actual string, t *testing.T) {
	if expected != actual {
		failTest(fmt.Sprintf("Expected value is %q, got %q.", expected, actual), t)
	}
}

func CheckStringNe(expected, actual string, t *testing.T) {
	if expected == actual {
		failTest(fmt.Sprintf("Expected value != %q, got %q.", expected, actual), t)
	}
}

func CheckStringLt(expected, actual string, t *testing.T) {
	if expected >= actual {
		failTest(fmt.Sprintf("Expected value < %q, got %q.", expected, actual), t)
	}
}

func CheckStringGt(expected, actual string, t *testing.T) {
	if expected <= actual {
		failTest(fmt.Sprintf("Expected value > %q, got %q.", expected, actual), t)
	}
}

func CheckStringLe(expected, actual string, t *testing.T) {
	if expected > actual {
		failTest(fmt.Sprintf("Expected value <= %q, got %q.", expected, actual), t)
	}
}

func CheckStringGe(expected, actual string, t *testing.T) {
	if expected < actual {
		failTest(fmt.Sprintf("Expected value >= %q, got %q.", expected, actual), t)
	}
}

// For int slice
func compareIntSlice(first, second []int) int {
	for i := 0; i < len(first); i++ {
		if i >= len(second) {
			return 1
		}
		if first[i] > second[i] {
			return 1
		}
		if first[i] < second[i] {
			return -1
		}
	}
	if len(first) == len(second) {
		return 0
	}
	return -1
}

func intSliceToString(ia []int) string {
	if len(ia) == 0 {
		return "[]"
	}
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "[%d", ia[0])
	for i := 1; i < len(ia); i++ {
		fmt.Fprintf(&buffer, ", %d", ia[i])
	}
	buffer.WriteString("]")
	return buffer.String()
}

func CheckIntSliceEq(expected, actual []int, t *testing.T) {
	if compareIntSlice(expected, actual) != 0 {
		failTest(fmt.Sprintf("Expected value is %s, got %s.", intSliceToString(expected), intSliceToString(actual)), t)
	}
}

// For int64 slice
func compareInt64Slice(first, second []int64) int {
	for i := 0; i < len(first); i++ {
		if i >= len(second) {
			return 1
		}
		if first[i] > second[i] {
			return 1
		}
		if first[i] < second[i] {
			return -1
		}
	}
	if len(first) == len(second) {
		return 0
	}
	return -1
}

func int64SliceToString(ia []int64) string {
	if len(ia) == 0 {
		return "[]"
	}
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "[%d", ia[0])
	for i := 1; i < len(ia); i++ {
		fmt.Fprintf(&buffer, ", %d", ia[i])
	}
	buffer.WriteString("]")
	return buffer.String()
}

func CheckInt64SliceEq(expected, actual []int64, t *testing.T) {
	if compareInt64Slice(expected, actual) != 0 {
		failTest(fmt.Sprintf("Expected value is %s, got %s.", int64SliceToString(expected), int64SliceToString(actual)), t)
	}
}

// For float32 slice
func compareFloat32Slice(first, second []float32) int {
	for i := 0; i < len(first); i++ {
		if i >= len(second) {
			return 1
		}
		if first[i] > second[i] {
			return 1
		}
		if first[i] < second[i] {
			return -1
		}
	}
	if len(first) == len(second) {
		return 0
	}
	return -1
}

func float32SliceToString(fa []float32) string {
	if len(fa) == 0 {
		return "[]"
	}
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "[%f", fa[0])
	for i := 1; i < len(fa); i++ {
		fmt.Fprintf(&buffer, ", %f", fa[i])
	}
	buffer.WriteString("]")
	return buffer.String()
}

func CheckFloat32SliceEq(expected, actual []float32, t *testing.T) {
	if compareFloat32Slice(expected, actual) != 0 {
		failTest(fmt.Sprintf("Expected value is %s, got %s.", float32SliceToString(expected), float32SliceToString(actual)), t)
	}
}

// For float64 slice
func compareFloat64Slice(first, second []float64) int {
	for i := 0; i < len(first); i++ {
		if i >= len(second) {
			return 1
		}
		if first[i] > second[i] {
			return 1
		}
		if first[i] < second[i] {
			return -1
		}
	}
	if len(first) == len(second) {
		return 0
	}
	return -1
}

func float64SliceToString(fa []float64) string {
	if len(fa) == 0 {
		return "[]"
	}
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "[%f", fa[0])
	for i := 1; i < len(fa); i++ {
		fmt.Fprintf(&buffer, ", %f", fa[i])
	}
	buffer.WriteString("]")
	return buffer.String()
}

func CheckFloat64SliceEq(expected, actual []float64, t *testing.T) {
	if compareFloat64Slice(expected, actual) != 0 {
		failTest(fmt.Sprintf("Expected value is %s, got %s.", float64SliceToString(expected), float64SliceToString(actual)), t)
	}
}

// For string slice
func compareStringSlice(first, second []string) int {
	for i := 0; i < len(first); i++ {
		if i >= len(second) {
			return 1
		}
		if first[i] > second[i] {
			return 1
		}
		if first[i] < second[i] {
			return -1
		}
	}
	if len(first) == len(second) {
		return 0
	}
	return -1
}

func stringSliceToString(sa []string) string {
	if len(sa) == 0 {
		return "[]"
	}
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "[%q", sa[0])
	for i := 1; i < len(sa); i++ {
		fmt.Fprintf(&buffer, ", %q", sa[i])
	}
	buffer.WriteString("]")
	return buffer.String()
}

func CheckStringSliceEq(expected, actual []string, t *testing.T) {
	if compareStringSlice(expected, actual) != 0 {
		failTest(fmt.Sprintf("Expected value is %s, got %s.", stringSliceToString(expected), stringSliceToString(actual)), t)
	}
}
