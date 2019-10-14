package bw

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fatih/color"
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

func TestChangeBubbleColor(t *testing.T) {
	numbers := 220

	before := color.New(color.FgCyan)
	after := color.New(color.FgRed)
	bw := NewBubbleWrap(numbers).ChangeBubbleColor(before, after)
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(10 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}

func TestInterrupt(t *testing.T) {
	numbers := 220

	bw := NewBubbleWrap(numbers)
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(10 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
		if i == 120 {
			bw.Interrupt()
		}
	}
}

func TestCompleted(t *testing.T) {
	numbers := 220

	bw := NewBubbleWrap(numbers)
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(10 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
	bw.Display()
	bw.Clear()
	bw.Redisplay()
}
