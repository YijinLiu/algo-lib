package strutil

/*
#cgo CFLAGS: -std=c99

void uint32_to_ordered_bytes(unsigned int u32, char* bytes);
unsigned int ordered_bytes_to_uint32(const char* bytes);

void int32_to_ordered_bytes(int i32, char* bytes);
int ordered_bytes_to_int32(const char* bytes);

void uint64_to_ordered_bytes(unsigned long long u64, char* bytes);
unsigned long long ordered_bytes_to_uint64(const char* bytes);

void int64_to_ordered_bytes(long long i64, char* bytes);
long long ordered_bytes_to_int64(const char* bytes);

void float32_to_ordered_bytes(float f32, char* bytes);
float ordered_bytes_to_float32(const char* bytes);

void float64_to_ordered_bytes(double f64, char* bytes);
double ordered_bytes_to_float64(const char* bytes);
*/
import "C"

import (
	"unsafe"

	"logging"
)

func stringToCharPointer(str string) *C.char {
	bytes := []byte(str)
	return (*C.char)(unsafe.Pointer(&bytes[0]))
}

func Uint32ToOrderedString(u32 uint32) string {
	var serialized [4]byte
	C.uint32_to_ordered_bytes(C.uint(u32), (*C.char)(unsafe.Pointer(&serialized[0])))
	return string(serialized[:])
}

func OrderedStringToUint32(str string) uint32 {
	if len(str) != 4 {
		logging.Fatalf("Invalid string: %q.", str)
	}
	return uint32(C.ordered_bytes_to_uint32(stringToCharPointer(str)))
}

func Int32ToOrderedString(i32 int32) string {
	var serialized [4]byte
	C.int32_to_ordered_bytes(C.int(i32), (*C.char)(unsafe.Pointer(&serialized[0])))
	return string(serialized[:])
}

func OrderedStringToInt32(str string) int32 {
	if len(str) != 4 {
		logging.Fatalf("Invalid string: %q.", str)
	}
	return int32(C.ordered_bytes_to_int32(stringToCharPointer(str)))
}

func Uint64ToOrderedString(u64 uint64) string {
	var serialized [8]byte
	C.uint64_to_ordered_bytes(C.ulonglong(u64), (*C.char)(unsafe.Pointer(&serialized[0])))
	return string(serialized[:])
}

func OrderedStringToUint64(str string) uint64 {
	if len(str) != 8 {
		logging.Fatalf("Invalid string: %q.", str)
	}
	return uint64(C.ordered_bytes_to_uint64(stringToCharPointer(str)))
}

func Int64ToOrderedString(i64 int64) string {
	var serialized [8]byte
	C.int64_to_ordered_bytes(C.longlong(i64), (*C.char)(unsafe.Pointer(&serialized[0])))
	return string(serialized[:])
}

func OrderedStringToInt64(str string) int64 {
	if len(str) != 8 {
		logging.Fatalf("Invalid string: %q.", str)
	}
	return int64(C.ordered_bytes_to_int64(stringToCharPointer(str)))
}

func Float32ToOrderedString(f32 float32) string {
	var serialized [4]byte
	C.float32_to_ordered_bytes(C.float(f32), (*C.char)(unsafe.Pointer(&serialized[0])))
	return string(serialized[:])
}

func OrderedStringToFloat32(str string) float32 {
	if len(str) != 4 {
		logging.Fatalf("Invalid string: %q.", str)
	}
	return float32(C.ordered_bytes_to_float32(stringToCharPointer(str)))
}

func Float64ToOrderedString(f64 float64) string {
	var serialized [8]byte
	C.float64_to_ordered_bytes(C.double(f64), (*C.char)(unsafe.Pointer(&serialized[0])))
	return string(serialized[:])
}

func OrderedStringToFloat64(str string) float64 {
	if len(str) != 8 {
		logging.Fatalf("Invalid string: %q.", str)
	}
	return float64(C.ordered_bytes_to_float64(stringToCharPointer(str)))
}
