package strconvx

import (
	"strconv"
)

func MustStrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StrToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StrToUInt64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func MustStrToInt64(s string) int64 {
	num, _ := StrToInt64(s)
	return num
}

func MustStrToUInt(s string) uint {
	num, _ := strconv.ParseUint(s, 10, 64)
	return uint(num)
}

func MustStrToUInt64(s string) uint64 {
	num, _ := strconv.ParseUint(s, 10, 64)
	return num
}

func IntToStr(val int) string {
	return strconv.Itoa(val)
}

func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

func UIntToStr(val uint) string {
	return strconv.FormatUint(uint64(val), 10)
}

func UInt64ToStr(val uint64) string {
	return strconv.FormatUint(val, 10)
}
