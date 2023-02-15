package algorithms

func addToArrayForm(num []int, k int) []int {

	for i := len(num) - 1; i >= 0; i-- {
		num[i] += k % 10
		k /= 10
		if num[i] >= 10 {
			k++
			num[i] -= 10
		}
	}

	for k > 0 {
		num = append([]int{k % 10}, num...)
		k /= 10
	}

	return num
}