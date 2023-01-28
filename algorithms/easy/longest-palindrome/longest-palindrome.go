package algortihms

func longestPalindrome(s string) int {

	// Map to store the count of each character
	m := map[rune]int{}
	ans := 0
	hasOdd := 0

	for _, c := range s {
		m[c]++
	}

	for _, v := range m {
		if v%2 == 0 {
			ans += v
		} else {
			ans += v - 1
			hasOdd = 1
		}
	}

	return ans + hasOdd
}
