package bw

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fatih/color"
)

type BubbleWrap struct {
	numbers     int
	column      int
	bubbles     []bool
	before      string
	after       string
	beforeColor color.Color
	afterColor  color.Color
	interrupted bool
	completed   bool
}

const cursorHide = "\033[?25l"
const cursorShow = "\033[?25h"

func NewBubbleWrap(numbers int) *BubbleWrap {
	bw := new(BubbleWrap)
	bw.printCopyRight()

	if numbers <= 0 {
		return nil
	} else {
		bw.numbers = numbers
	}

	bw.column = 50
	bw.before = "*"
	bw.after = "."

	bw.bubbles = make([]bool, bw.numbers)

	return bw
}

func (bw *BubbleWrap) SetColumn(column int) *BubbleWrap {
	bw.column = column
	return bw
}

func (bw *BubbleWrap) ChangeBubbleShape(before string, after string) *BubbleWrap {
	bw.before = before
	bw.after = after
	return bw
}

func (bw *BubbleWrap) ChangeBubbleColor(before *color.Color, after *color.Color) *BubbleWrap {
	bw.beforeColor = *before
	bw.afterColor = *after
	return bw
}

func (bw *BubbleWrap) Pop(i int) {
	if bw.interrupted || bw.completed {
		return
	}
	bw.bubbles[i] = true
}

func (bw *BubbleWrap) Display() {
	if bw.interrupted || bw.completed {
		return
	}

	// display bubbles
	var c, comp int
	for _, v := range bw.bubbles {
		if v {
			if &bw.beforeColor != nil {
				bw.beforeColor.Printf(cursorHide+"%v", bw.before)
			} else {
				fmt.Printf(cursorHide+"%v", bw.before)
			}
			comp++
		} else {
			if &bw.afterColor != nil {
				bw.afterColor.Printf(cursorHide+"%v", bw.after)
			} else {
				fmt.Printf(cursorHide+"%v", bw.after)
			}
		}

		c++
		if c == bw.column {
			fmt.Printf(cursorHide + "\n")
			c = 0
		} else if bw.numbers < bw.column && c == bw.numbers {
			break
		}
	}

	// display progress
	fmt.Printf(cursorHide + "\n")
	var space int
	if bw.numbers < bw.column {
		space = bw.numbers
	} else {
		space = bw.column
	}
	for i := 0; i < space-4; i++ {
		fmt.Printf(cursorHide + " ")
	}
	persent := (float64(comp) / float64(bw.numbers)) * 100.0
	trunc := math.Trunc(persent)
	fmt.Printf(cursorHide+"%3v%% \n", trunc)

	for i := 0; i < space-len(strconv.Itoa(comp))-len(strconv.Itoa(bw.numbers))-3; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf(cursorHide+"%d / %d", comp, bw.numbers)

	// output result
	if trunc == 100 {
		bw.completed = true
		fmt.Printf(cursorShow + "\n\nBubble Wrape finished.\n")
	}
}

func (bw *BubbleWrap) Clear() {
	if bw.interrupted || bw.completed {
		return
	}

	var c, r int
	r = int(bw.numbers / bw.column)
	c = bw.numbers + 2
	// move cursor to start point
	fmt.Printf(cursorHide+"\033[%dA", r+2)
	fmt.Printf(cursorHide+"\033[%dD", c)
}

func (bw *BubbleWrap) Redisplay() {
	bw.Clear()
	bw.Display()
}

func (bw *BubbleWrap) Interrupt() {
	if bw.interrupted || bw.completed {
		return
	}
	bw.interrupted = true
	fmt.Printf(cursorShow + "\n\nBubble Wrape interrupted.\n")
}

func (bw *BubbleWrap) printCopyRight() {
	fmt.Printf("Bubble Wrap %v by kura.\n\n", bw.version())
}

func (bw *BubbleWrap) version() string {
	return "1.0.7"
}
