package piscine

func isvalidebase(base string) bool {
	if len(base) < 2 {
		return false
	}
	check := make(map[rune]bool)
	for _, v := range base {
		if check[v] {
			return false
		}
		if v == '-' || v == '+' {
			return false
		}
		check[v] = true
	}
	return true
}

func AtoiBase(s string, base string) int {
	res := 0
	index := 0
	if !isvalidebase(base) {
		return 0
	}

	for i := range s {
		for j := range base {
			if s[i] == base[j] {
				index = j
				break
			}
		}
		res = res * len(base) + index
	}
	return res
}

