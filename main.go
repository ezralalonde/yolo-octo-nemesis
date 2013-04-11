package main

import (
	"dpa/dpa"
	"fmt"
)

func main() {
	var lines int
	if _, err := fmt.Scanf("%d", &lines); err != nil {
		println("oops")
		return
	}
	for ii := 0; ii < lines; ii++ {
		var line int
		fmt.Scanf("%d", &line)
		fmt.Printf(dpa.Calc(line), line)
	}
}
