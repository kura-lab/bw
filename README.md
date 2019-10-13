# Terminal progress table like Bubble Wrap for Golang

# Demo

![bubble_wrap](https://user-images.githubusercontent.com/1156984/66710841-4becfb00-edbb-11e9-855f-d68a86e60af3.gif)

# Installation

```
go get github.com/kura-lab/bw
```

# Get start

```
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
