package util

// Contains tells whether `entry` is an element in `slice`
func Contains(slice []string, entry string) bool {
	for _, value := range slice {
		if value == entry {
			return true
		}
	}
	return false
}
