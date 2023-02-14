package algorithms

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	carry := byte('0')
	ans := ""
	for i := 0; i < len(b); i++ {
		switch byte(a[len(a)-1-i]) + byte(b[len(b)-1-i]) + carry {
		case '0' + '0' + '0':
			ans = "0" + ans
			carry = '0'
		case '1' + '0' + '0':
			ans = "1" + ans
			carry = '0'
		case '1' + '1' + '0':
			ans = "0" + ans
			carry = '1'
		case '1' + '1' + '1':
			ans = "1" + ans
			carry = '1'
		}
	}

	for i := len(b); i < len(a); i++ {
		switch byte(a[len(a)-i]) + carry {
		case '0' + '0':
			ans = "0" + ans
			carry = '0'
		case '1' + '0':
			ans = "1" + ans
			carry = '0'
		case '1' + '1':
			ans = "0" + ans
			carry = '1'
		}
	}
	if carry == '1' {
		ans = "1" + ans
	}

	return ans
}
