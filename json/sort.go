package json

import "encoding/json"

// SortJson sorts the received json lexicographically and returns the bytes
func SortJson(jsonBytes []byte) ([]byte, error) {
	var obj interface{}
	err := json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		return nil, err
	}
	// Marshalling sorts keys lexicographically
	return json.Marshal(obj)
}
