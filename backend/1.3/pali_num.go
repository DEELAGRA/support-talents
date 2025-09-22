package main

func pali_num(x int) bool {
	original := x
	reverse := 0
	flag := false
	for x > 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	if original == reverse {
		flag = true
	}

	return flag
}
