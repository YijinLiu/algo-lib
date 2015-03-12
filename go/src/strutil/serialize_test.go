package strutil

import . "testutils"

import (
	"testing"
)

func TestInt32ToOrderedString(t *testing.T) {
	CheckStringEq("\x80\x01\x86\xa1", Int32ToOrderedString(100001), t)
	CheckStringEq("\x80\x01\x86\xa2", Int32ToOrderedString(100002), t)
	CheckStringEq("\u007f\xfey_", Int32ToOrderedString(-100001), t)
	CheckStringLt(Int32ToOrderedString(0), Int32ToOrderedString(1), t)
	CheckStringLt(Int32ToOrderedString(-1), Int32ToOrderedString(0), t)
	CheckStringLt(Int32ToOrderedString(-1), Int32ToOrderedString(1), t)
	CheckStringLt(Int32ToOrderedString(-100), Int32ToOrderedString(-10), t)
	CheckStringLt(Int32ToOrderedString(50), Int32ToOrderedString(1000), t)
}

func TestOrderedStringToInt32(t *testing.T) {
	CheckInt32Eq(15, OrderedStringToInt32(Int32ToOrderedString(15)), t)
	CheckInt32Eq(0, OrderedStringToInt32(Int32ToOrderedString(0)), t)
	CheckInt32Eq(-150, OrderedStringToInt32(Int32ToOrderedString(-150)), t)
}

func TestUint32ToOrderedString(t *testing.T) {
	CheckStringEq("\x00\x01\x86\xa1", Uint32ToOrderedString(100001), t)
	CheckStringEq("\x00\x01\x86\xa2", Uint32ToOrderedString(100002), t)
	CheckStringEq("\x80\x00\x00\x00", Uint32ToOrderedString(0x80000000), t)
	CheckStringLt(Uint32ToOrderedString(0), Uint32ToOrderedString(1), t)
	CheckStringLt(Uint32ToOrderedString(1), Uint32ToOrderedString(10), t)
	CheckStringLt(Uint32ToOrderedString(10), Uint32ToOrderedString(1000), t)
	CheckStringLt(Uint32ToOrderedString(1000), Uint32ToOrderedString(0xffffffff), t)
}

func TestOrderedStringToUint32(t *testing.T) {
	CheckUint32Eq(0, OrderedStringToUint32(Uint32ToOrderedString(0)), t)
	CheckUint32Eq(15, OrderedStringToUint32(Uint32ToOrderedString(15)), t)
	CheckUint32Eq(0xffffffff, OrderedStringToUint32(Uint32ToOrderedString(0xffffffff)), t)
}

func TestInt64ToOrderedString(t *testing.T) {
	CheckStringEq("\x80\x00\x00\x00\x00\x01\x86\xa1", Int64ToOrderedString(100001), t)
	CheckStringEq("\x80\x00\x00\x00\x00\x01\x86\xa2", Int64ToOrderedString(100002), t)
	CheckStringEq("\u007f\xff\xff\xff\xff\xfey_", Int64ToOrderedString(-100001), t)
	CheckStringLt(Int64ToOrderedString(0), Int64ToOrderedString(1), t)
	CheckStringLt(Int64ToOrderedString(-1), Int64ToOrderedString(0), t)
	CheckStringLt(Int64ToOrderedString(-1), Int64ToOrderedString(1), t)
	CheckStringLt(Int64ToOrderedString(-100), Int64ToOrderedString(-10), t)
	CheckStringLt(Int64ToOrderedString(50), Int64ToOrderedString(1000), t)
}

func TestOrderedStringToInt64(t *testing.T) {
	CheckInt64Eq(15, OrderedStringToInt64(Int64ToOrderedString(15)), t)
	CheckInt64Eq(0, OrderedStringToInt64(Int64ToOrderedString(0)), t)
	CheckInt64Eq(-150, OrderedStringToInt64(Int64ToOrderedString(-150)), t)
}

func TestUint64ToOrderedString(t *testing.T) {
	CheckStringEq("\x00\x00\x00\x00\x00\x01\x86\xa1", Uint64ToOrderedString(100001), t)
	CheckStringEq("\x00\x00\x00\x00\x00\x01\x86\xa2", Uint64ToOrderedString(100002), t)
	CheckStringEq("\x80\x00\x00\x00\x00\x00\x00\x00", Uint64ToOrderedString(0x8000000000000000), t)
	CheckStringLt(Uint64ToOrderedString(0), Uint64ToOrderedString(1), t)
	CheckStringLt(Uint64ToOrderedString(1), Uint64ToOrderedString(10), t)
	CheckStringLt(Uint64ToOrderedString(10), Uint64ToOrderedString(1000), t)
	CheckStringLt(Uint64ToOrderedString(1000), Uint64ToOrderedString(0xffffffffffffffff), t)
}

func TestOrderedStringToUint64(t *testing.T) {
	CheckUint64Eq(0, OrderedStringToUint64(Uint64ToOrderedString(0)), t)
	CheckUint64Eq(15, OrderedStringToUint64(Uint64ToOrderedString(15)), t)
	CheckUint64Eq(0xffffffff, OrderedStringToUint64(Uint64ToOrderedString(0xffffffff)), t)
}

func TestFloat32ToOrderedString(t *testing.T) {
	CheckStringEq("\x80\x00\x00\x00", Float32ToOrderedString(0.0), t)
	CheckStringEq("\xbf\xd9\x99\x9a", Float32ToOrderedString(1.7), t)
	CheckStringEq("\x3e\xd1\x99\x9a", Float32ToOrderedString(-10.9), t)
	CheckStringLt(Float32ToOrderedString(0.0), Float32ToOrderedString(9.8), t)
	CheckStringLt(Float32ToOrderedString(9.8), Float32ToOrderedString(1.3e10), t)
	CheckStringLt(Float32ToOrderedString(-1.0), Float32ToOrderedString(0.0), t)
	CheckStringLt(Float32ToOrderedString(-1.5e3), Float32ToOrderedString(-1.0), t)
}

func TestOrderedStringToFloat32(t *testing.T) {
	CheckFloat32Eq(0.0, OrderedStringToFloat32(Float32ToOrderedString(0.0)), t)
	CheckFloat32Eq(1.5, OrderedStringToFloat32(Float32ToOrderedString(1.5)), t)
	CheckFloat32Eq(-1.3e8, OrderedStringToFloat32(Float32ToOrderedString(-1.3e8)), t)
}

func TestFloat64ToOrderedString(t *testing.T) {
	CheckStringLt(Float64ToOrderedString(0.0), Float64ToOrderedString(9.8), t)
	CheckStringLt(Float64ToOrderedString(9.8), Float64ToOrderedString(1.3e10), t)
	CheckStringLt(Float64ToOrderedString(-1.0), Float64ToOrderedString(0.0), t)
	CheckStringLt(Float64ToOrderedString(-1.5e3), Float64ToOrderedString(-1.0), t)
}

func TestOrderedStringToFloat64(t *testing.T) {
	CheckFloat64Eq(0.0, OrderedStringToFloat64(Float64ToOrderedString(0.0)), t)
	CheckFloat64Eq(1.5, OrderedStringToFloat64(Float64ToOrderedString(1.5)), t)
	CheckFloat64Eq(-1.3e8, OrderedStringToFloat64(Float64ToOrderedString(-1.3e8)), t)
}
