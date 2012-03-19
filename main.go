package main;

import (
	"exp/gui/x11"
	"image"
	"time"
	"./life"
)

func main() {
	l := life.New(600,800)
	l.Init()
	w,_ := x11.NewWindow()
	img := w.Screen()
	c1 := image.RGBAColor{10,10,10,255}
	c2 := image.RGBAColor{222,222,222,255}
	for true {
		w.FlushImage()
		col := l.GetColony()
		for i := 0; i < len(col); i++ {
		for j := 0; j < len(col[i]); j++ {
			if col[i][j] {
				img.Set(i, j, c2)
			} else {
				img.Set(i, j, c1)
			}
		}}
		l.Update()
		time.Sleep(25000000)
	}
}
