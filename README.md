# Terminal progress table like Bubble Wrap for Golang

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
