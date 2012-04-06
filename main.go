package main

import (
	"./life"
	"time"
	"fmt"
)

func main() {
	l := life.New(80, 24)
	l.Init()
	for true {
		col := l.GetColony()
		for i := 0; i < len(col); i++ {
			fmt.Println("")
			for j := 0; j < len(col[i]); j++ {
				if col[i][j] {
					fmt.Print("X")
				} else {
					fmt.Print(" ")
				}
			}
		}
		l.Update()
		time.Sleep(250000000)
	}
}
