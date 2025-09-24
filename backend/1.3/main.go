package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 6, 3, 2, 5}
	target := 8
	fmt.Println("Индексы суммы двух чисел:", two_sum(arr, target))

	x := 1221
	fmt.Println("Число полиндром?", pali_num(x))

	s := "MCMXCIV"
	fmt.Println(romanToInt(s))

	par := "([)}"
	fmt.Println(valid(par))

	oneplus := []int{4,3,2,1}
	fmt.Println(plusOne(oneplus))
}
