package utils

func Filter(nums []int, callback func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}
