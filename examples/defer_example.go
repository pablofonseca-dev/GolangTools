package examples

import "fmt"

func EmulateDefer() {
	for i := 0; i <= 20; i++ {
		defer fmt.Println(i)
	}
}
