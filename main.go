/*
	A simulation of Conway's Game of Life in a terminal.
*/
package main

import (
	"fmt"
	"time"
	"./life"
)

func main() {
	l := life.New(80,24)
	l.Init()
	for generation := 0; ; generation++ {
		l.Map(func(state bool) { 
			var str string
			if state { str = "\x1b[47m " } else { str = "\x1b[41m " }
			fmt.Printf("%s", str)
		})
		fmt.Printf("\x1b[0m")
		l.Update()
		time.Sleep(1000000000)
	}
}
