package concatenation_of_array

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
