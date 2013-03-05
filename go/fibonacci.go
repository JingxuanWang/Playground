package main

import (
	"fmt"
)

func main() {
	var pre int = 0
	var cur int = 1
	for i := 0; i < 100; i++ {
		cur = cur + pre
		pre = cur - pre
	}
	fmt.Println(cur);
}
