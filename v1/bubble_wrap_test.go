package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	numbers := 220

	slice := make([]int, numbers)

	for i := range slice {
		slice[i] = i
	}
	n := len(slice)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	bw := NewBubbleWrap(numbers)
	bw.Display()

	for _, i := range slice {
		time.Sleep(20 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}
