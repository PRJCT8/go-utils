package _map

func ConvertInterfaceToMap[K comparable, V any](input interface{}) map[K]V {
	result, ok := input.(map[K]V)
	if !ok {
		return nil
	}
	return result
}
