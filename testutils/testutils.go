package testutils

//IntSliceAreEqual foo
func IntSliceAreEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, j := 0, 0; i < len(a); i, j = i+1, j+1 {
		if a[i] != b[j] {
			return false
		}
	}

	return true
}

//StringSliceAreEqual foo
func StringSliceAreEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, j := 0, 0; i < len(a); i, j = i+1, j+1 {
		if a[i] != b[j] {
			return false
		}
	}

	return true
}
