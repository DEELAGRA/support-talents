package main

type Stack struct {
	item []int
}

func (s *Stack) Push(item int) {
	s.item = append(s.item, item)
}

func (s *Stack) Pop() int {

	if len(s.item) == 0 {
		return 0
	}

	lastIndex := len(s.item) - 1
	item := s.item[lastIndex]
	s.item = s.item[:lastIndex]
	return item
}
