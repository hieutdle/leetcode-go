package subtract_the_product_and_sum_of_digits_of_an_integer

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1

	for n > 0 {
		sum += n % 10
		product *= n % 10
		n = n / 10
	}
	
	return product - sum
}
