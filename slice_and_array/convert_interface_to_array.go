package slice_and_array

func ConvertInterfaceToArray[T any](input interface{}) []T {
	result, ok := input.([]T)
	if !ok {
		return nil
	}
	return result
}
