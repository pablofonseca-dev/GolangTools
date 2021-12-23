package utils

import (
	"fmt"
	"time"
)

func EmulateIntegerOverflow() {
	var x uint8 = 0

	for i := 0; i < 257; i++ {
		x = uint8(i)
		if !(i >= 250) {
			time.Sleep(50 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
			if x == 0 {
				fmt.Printf("%d --> This is an integer overflow\n", x)
				return
			}
		}
		fmt.Println(x)
	}
}
