package jef

func ReverseStringSlice(ss []string) {
	for i, j := 0, len(ss)-1; j > i; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
}

func ReverseIntSlice(si []int) {
	for i, j := 0, len(si)-1; j > i; i, j = i+1, j-1 {
		si[i], si[j] = si[j], si[i]
	}
}
