package utils 

// IsValInSlice python a in []
func IsValInSlice(e int, slice []int) bool {
	for _, s := range slice {
		if s == e {
			return true
		}
	}
	return false
}
