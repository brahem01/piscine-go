package piscine

func Permutations(input string) []string {
	if len(input) < 2 {
		return []string{input}
	}

	perms := Permutations(input[1:])
	result := make([]string, 0)

	for _, perm := range perms {
		for i := 0; i <= len(perm); i++ {
			newPerm := perm[:i] + string(input[0]) + perm[i:]
			result = append(result, newPerm)
		}
	}

	// Remove duplicates by converting the slice to a map and back to a slice
	uniquePerms := make(map[string]bool)
	for _, perm := range result {
		uniquePerms[perm] = true
	}

	result = make([]string, 0, len(uniquePerms))
	for perm := range uniquePerms {
		result = append(result, perm)
	}
	return result

}
