package main

import (
	"Aquafalx/lib/reality"
	"time"
)

func main() {
	b := BoardCreate("Test1")
	b.status()

}

//Board : Contains each team
type Board struct {
	name      string
	TimeStart time.Time
	TimeEnd   time.Time
	Teams     map[string]reality.Team
}

//BoardCreate : creates a new gameboard
func BoardCreate(n string) Board {
	b := Board{name: n}
	b.TimeStart = time.Now()
	return b
}

func (b Board) status() {
	time.Sleep(2 * time.Second)
	t := time.Now()
	elapsed := t.Sub(b.TimeStart)
	println(elapsed)
}
