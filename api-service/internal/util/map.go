package util

func ArrToMap[T any](slice []T, idFunc func(T) int64) map[int64]T {
	result := make(map[int64]T)

	for _, item := range slice {
		id := idFunc(item)
		result[id] = item
	}

	return result
}

func ArrToMapArr[T any](slice []T, idFunc func(T) int64) map[int64][]T {
	result := make(map[int64][]T)
	for _, item := range slice {
		id := idFunc(item)
		result[id] = append(result[id], item)
	}

	return result
}
