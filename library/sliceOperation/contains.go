package sliceOperation

func Contains(data []string, s string) bool {
	for _, v := range data {
		if s == v {
			return true
		}
	}
	return false
}
