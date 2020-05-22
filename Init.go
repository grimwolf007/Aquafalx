package main

import (
	"Aquafalx/lib/reality"
	"time"
)

func main() {
	b := BoardCreate("Test1")
	b.status()
	l := reality.DroneLocationCreate(0, 0, 0)
	teams := []reality.Team{reality.TeamCreate("Red")}
	bear := reality.DroneBearingCreate(0, 0, 0)
	drones := startDrones(10, bear, l, teams[0])
	print(drones)
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

func startDrones(c int, b reality.DroneBearing, l reality.DroneLocation, t reality.Team) []reality.Drone {
	ds := make([]reality.Drone, c)
	for i := 0; i < c; i++ {
		ds = append(ds, reality.DroneCreate(l, b, reality.DRONEISR, t.Name(), 20))
	}
	return ds
}
