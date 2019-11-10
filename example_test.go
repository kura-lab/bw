package bw_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fatih/color"
	"github.com/kura-lab/bw"
)

func Example(t *testing.T) {
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

	bw := bw.NewBubbleWrap(numbers)
	bw.Display()

	for _, i := range slice {
		time.Sleep(20 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}
func ExampleBubbleWrap_SetColumn(t *testing.T) {
	numbers := 220

	bw := bw.NewBubbleWrap(numbers).SetColumn(60)
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(10 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}

func ExampleBubbleWrap_ChangeBubbleShape(t *testing.T) {
	numbers := 220

	bw := bw.NewBubbleWrap(numbers).ChangeBubbleShape("o", "x")
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(10 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}

func ExampleBubbleWrap_ChangeBubbleColor(t *testing.T) {
	numbers := 220

	before := color.New(color.FgCyan)
	after := color.New(color.FgRed)
	bw := bw.NewBubbleWrap(numbers).ChangeBubbleColor(before, after)
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(10 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}

func ExampleBubbleWrap_Interrupt(t *testing.T) {
	numbers := 220

	bw := bw.NewBubbleWrap(numbers)
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

func ExampleBubbleWrap_Clear(t *testing.T) {
	numbers := 220

	bw := bw.NewBubbleWrap(numbers)
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
