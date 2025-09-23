package main

func valid(s string) bool {

	pairs := []rune{}

	simbol := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, val := range s {
		if val == '(' || val == '[' || val == '{' {
			pairs = append(pairs, val)
		} else {
			if len(pairs) == 0 {
				return false
			}
			last := pairs[len(pairs)-1]
			pairs = pairs[:len(pairs)-1]
			if last != simbol[val] {
				return false
			}
		}

	}
	return len(pairs) == 0
}
