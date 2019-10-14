package bw

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fatih/color"
)

type BubbleWrap struct {
	Numbers     int
	Column      int
	Bubbles     []bool
	before      string
	after       string
	beforeColor color.Color
	afterColor  color.Color
}

const cursorHide = "\033[?25l"

func NewBubbleWrap(numbers int) *BubbleWrap {
	bw := new(BubbleWrap)
	bw.printCopyRight()

	if numbers <= 0 {
		return nil
	} else {
		bw.Numbers = numbers
	}

	bw.Column = 50
	bw.before = "*"
	bw.after = "."

	bw.Bubbles = make([]bool, bw.Numbers)

	return bw
}

func (bw *BubbleWrap) SetColumn(column int) *BubbleWrap {
	bw.Column = column
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
	bw.Bubbles[i] = true
}

func (bw *BubbleWrap) Display() {
	// display bubbles
	var c, comp int
	for _, v := range bw.Bubbles {
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
		if c == bw.Column {
			fmt.Printf(cursorHide + "\n")
			c = 0
		} else if bw.Numbers < bw.Column && c == bw.Numbers {
			break
		}
	}

	// display progress
	fmt.Printf(cursorHide + "\n")
	var space int
	if bw.Numbers < bw.Column {
		space = bw.Numbers
	} else {
		space = bw.Column
	}
	for i := 0; i < space-4; i++ {
		fmt.Printf(cursorHide + " ")
	}
	persent := (float64(comp) / float64(bw.Numbers)) * 100.0
	trunc := math.Trunc(persent)
	fmt.Printf(cursorHide+"%3v%% \n", trunc)

	for i := 0; i < space-len(strconv.Itoa(comp))-len(strconv.Itoa(bw.Numbers))-3; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf(cursorHide+"%d / %d", comp, bw.Numbers)

	// output result
	if trunc == 100 {
		fmt.Printf("\n\nBubble Wrape finished.\n")
	}
}

func (bw *BubbleWrap) Clear() {
	var c, r int
	r = int(bw.Numbers / bw.Column)
	c = bw.Numbers + 2
	// move cursor to start point
	fmt.Printf(cursorHide+"\033[%dA", r+2)
	fmt.Printf(cursorHide+"\033[%dD", c)
}

func (bw *BubbleWrap) Redisplay() {
	bw.Clear()
	bw.Display()
}

func (bw *BubbleWrap) printCopyRight() {
	fmt.Printf("Bubble Wrap %v by kura.\n\n", bw.version())
}

func (bw *BubbleWrap) version() string {
	return "1.0.5"
}
