package main

func plusOne(digits []int) []int {
	last := digits[len(digits)-1]
	if last == 9 {
		digits = digits[:len(digits)-1]
		digits = append(digits, 1, 0)
		return digits
	} else {
		digits[len(digits)-1] += 1
		return digits
	}
}
