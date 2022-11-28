package utils

import (
	"fmt"
	"strings"
)

// StringInArray try to find string in array
func StringInArray(s string, arr ...string) bool {
	for i := range arr {
		if arr[i] == s {
			return true
		}
	}
	return false
}

func Uint64ToStringArray(arr []uint64) []string {
	var res []string
	for i := range arr {
		res = append(res, fmt.Sprint(arr[i]))
	}
	return res

}
func StringToUint64Array(arr []string) []uint64 {
	var res []uint64
	for i := range arr {
		res = append(res, StringToUint64(arr[i]))
	}
	return res
}

func IntersectionInt64(a []int64, b []int64) bool {
	low, high := a, b
	if len(a) > len(b) {
		low = b
		high = a
	}

	found := false
mainLoop:
	for _, l := range low {
		for _, h := range high {
			if l == h {
				found = true
				break mainLoop
			}
		}

	}
	return found

}

func IntersectionString(a []string, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func UniqueString(a []string) []string {
	var res = make([]string, 0)
	var found = make(map[string]bool)
	for i := range a {
		if !found[a[i]] {
			res = append(res, a[i])
			found[a[i]] = true
		}
	}
	return res
}

func UniqueInt64(a []int64) []int64 {
	var res = make([]int64, 0)
	var found = make(map[int64]bool)
	for i := range a {
		if !found[a[i]] {
			res = append(res, a[i])
			found[a[i]] = true
		}
	}
	return res
}

func CheckFullStringIntersection(a []string, b []string) bool { // a eshterake b =a
	a = UniqueString(a)
	b = UniqueString(b)
	inter := IntersectionString(a, b)
	return len(inter) == len(a)
}

func MergeUint(a []uint, b []uint) []uint {
	for _, val := range b {
		a = append(a, val)
	}

	return a
}

func DiffUInt(a []uint, b []uint) []uint {
	diff := make([]uint, 0)
	for _, val := range a {
		if !ContainsUInt(b, val) {
			diff = append(diff, val)
		}
	}

	return diff
}

func ContainsUInt(s []uint, val uint) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}

	return false
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func UintsToString(numbers []uint, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), delim), "[]")
}

func SliceToMapS(s []string) map[string]bool {
	setMap := make(map[string]bool)
	for _, s := range s {
		setMap[s] = true
	}

	return setMap
}
