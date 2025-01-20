package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPopStack(t *testing.T) {
    s := Stack{}
    s.Push(5)
    s.Push(3)
    s.Push(8)
    s.Push(2)

    assert.Equal(t, 2, s.Pop())
    assert.Equal(t, 8, s.Pop())
    assert.Equal(t, 3, s.Pop())
    assert.Equal(t, 5, s.Pop())
}

func TestGetMinStack(t *testing.T) {
    s := Stack{}
    s.Push(5)
    s.Push(3)
    s.Push(2)
    s.Push(8)
    s.Push(6)

    assert.Equal(t, 2, s.GetMin())
}

func TestGetMinStackWhenMinValueIsPoppedOnceTime(t *testing.T) {
    s := Stack{}
    s.Push(5)
    s.Push(8)
    s.Push(6)
    s.Push(2)
    s.Push(3)
    s.Pop()
    s.Pop()
    s.Pop()

    assert.Equal(t, 5, s.GetMin())
}
