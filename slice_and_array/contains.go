package slice_and_array

// depreciated as the standard lib has this method now
func contains[T comparable](list []T, element T) bool {
	for _, curr := range list {
		if curr == element {
			return true
		}
	}
	return false
}
