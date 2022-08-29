package array

// StringInArray try to find string in array
func StringInArray(s string, arr ...string) bool {
	for i := range arr {
		if arr[i] == s {
			return true
		}
	}
	return false
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
