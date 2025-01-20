package main

type Stack struct {
    data []int
    min  []int
}

func (s *Stack) Push(x int) {
    s.data = append(s.data, x)
    if len(s.min) == 0 || x <= s.GetMin() {
        s.min = append(s.min, x)
    }
}

func (s *Stack) Pop() int {
    if len(s.data) == 0 {
        return -1
    }
    top := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    if top == s.GetMin() {
        s.min = s.min[:len(s.min)-1]
    }

    return top
}

func (s *Stack) GetMin() int {
    if len(s.min) == 0 {
        panic("stack is empty")
    }
    return s.min[len(s.min)-1]
}
