package main

func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss, tt := countChar(s), countChar(t)
	for k, v := range ss {
		v2, ok := tt[k]
		if !ok || v2 != v {
			return false
		}
	}
	return true
}

func countChar(s string) map[rune]int {
	ss := make(map[rune]int)
	for _, v := range s {
		n, ok := ss[v]
		if ok {
			ss[v] = n + 1
		} else {
			ss[v] = 1
		}
	}
	return ss
}
