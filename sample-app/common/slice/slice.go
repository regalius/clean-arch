package slice

// IndexOfInt64 get index of element in slice (int64)
func IndexOfInt64(slice []int64, element int64) int {
	for index, el := range slice {
		if el == element {
			return index
		}
	}
	return -1
}

// ResliceInt64 remove element at given index and re-order
func ResliceInt64(slice []int64, index int) []int64 {
	return append(slice[:index], slice[index+1:]...)
}
