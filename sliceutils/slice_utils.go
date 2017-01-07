package sliceutils

func Contains(slice []string, str string) bool {
	if len(slice) == 0 {
		return false
	}
	for _, val := range slice {
		if val == str {
			return true
		}
	}
	return false
}
