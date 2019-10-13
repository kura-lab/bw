# Terminal progress table like Bubble Wrap for Golang

# Get start

```
package main

import (
	"time"
	"github.com/kura-lab/bubble_wrap/v1"
)

func main() {
	numbers := 220

	bw := NewBubbleWrap(numbers)
	bw.Display()

	for i := 0; i < numbers; i++ {
		time.Sleep(20 * time.Millisecond)
		bw.Pop(i)
		bw.Redisplay()
	}
}
```
