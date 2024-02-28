package piscine

func permutation(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	slice := []string{}
	perms:= permutation(s[1:])
	for _, perm := range perms {
		for i:= 0; i<= len(perm); i++ {
			newPerm:= perm[:i] + string(s[0]) + perm[i:]
			slice = append(slice, newPerm)
		}
		
	}
	return slice
}