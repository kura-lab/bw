# Terminal progress table like Bubble Wrap for Golang

# Demo

![bw](https://user-images.githubusercontent.com/1156984/66727039-b82d3480-ee77-11e9-8fba-4b7fa82dc50e.gif)

# Installation

```
go get github.com/kura-lab/bw
```

# Get start

```go
package main

import (
	"time"

	"github.com/kura-lab/bw"
)

func main() {
	numbers := 220

	bw := bw.NewBubbleWrap(numbers)
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(20 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}
```

# Settings

```go
numbers := 220

// set numbers of a column
bw := NewBubbleWrap(numbers).SetColumn(60)

// change bubble's shapes
bw := NewBubbleWrap(numbers).ChangeBubbleShape("o", "x")

// change bubble's colors
before := color.New(color.FgCyan)
after := color.New(color.FgRed)
bw := NewBubbleWrap(numbers).ChangeBubbleColor(before, after)
```
