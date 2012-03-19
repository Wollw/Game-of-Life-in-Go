package life

import (
	"rand"
	"time"
)

type lifeInterface interface {
	Init()
	Update()
	Map(func())
	GetColony() [][] bool
}

type lifeType struct {
	colony[][] bool
}

func (this *lifeType) GetColony() [][] bool {
	return this.colony
}

func (this *lifeType) neighborCount(x, y int) int {
	nc := 0
	for i := x-1; i <= x+1; i++ {
	for j := y-1; j <= y+1; j++ {
		if (i >= 0 && i < len(this.colony)) {
		if (j >= 0 && j < len(this.colony[i])) {
		if (i != x || j != y) {
			if (this.colony[i][j]) {
				nc++
			}
		}}}
	}}
	return nc
}

func (this *lifeType) Init() {
	rand.Seed(time.Nanoseconds());
	for i := 0; i < len(this.colony); i++ {
	for j := 0; j < len(this.colony[i]); j++ {
		this.colony[i][j] = rand.Int() % 2 == 0
	}}
}

func (this *lifeType) Update() {

	next := make([][]bool, len(this.colony))
	for i := 0; i < len(this.colony); i++ {
		next[i] = make([]bool, len(this.colony[i]))
	}

	for x := 0; x < len(this.colony); x++ {
	for y := 0; y < len(this.colony[x]); y++ {
		nc := this.neighborCount(x,y)
		if (this.colony[x][y]) {
			switch (nc) {
				case 2, 3:
					next[x][y] = true
					break
				default:
					next[x][y] = false
			}
		} else {
			switch (nc) {
				case 3:
					next[x][y] = true
					break
				default:
					next[x][y] = false
			}
		}
	}}
	this.colony = next
}

func (this *lifeType) Map(fn func(state bool)) {
	for _,v := range this.colony {
	for _,v := range v {
		fn(v)
	}}
}

func (this *lifeType) Reduce(fn func()) {
	fn()
}

func New(x,y int) *lifeType {
	l := new(lifeType)
	l.colony = make([][]bool, y)
	for i := 0; i < len(l.colony); i++ {
		l.colony[i] = make([]bool, x)
	}
	return l
}
