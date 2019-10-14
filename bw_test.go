package bw

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

func TestSetColumn(t *testing.T) {
	numbers := 220

	bw := NewBubbleWrap(numbers).SetColumn(60)
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(10 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}

func TestChangeBubbleShape(t *testing.T) {
	numbers := 220

	bw := NewBubbleWrap(numbers).ChangeBubbleShape("o", "x")
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(10 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}
