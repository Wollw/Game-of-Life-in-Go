/*
	A simulation of Conway's Game of Life in a terminal.
*/

package main

import (
	"time"
	"./life"
)

func main() {
	l := life.New(80,24)
	l.Init()
	for generation := 0; ; generation++ {
		l.Print()
		l.Update()
		time.Sleep(1000000000)
	}
}
