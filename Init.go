package main

import (
	"Aquafalx/lib/reality"
	"time"
)

func main() {
	count := 10
	b := BoardCreate("Test1")
	l := reality.DroneLocationCreate(0, 0, 0)
	teams := []reality.Team{reality.TeamCreate("Red")}
	bear := reality.DroneBearingCreate(0, 0, 0)
	drones := startDrones(count, bear, l, teams[0])
	for i := 0; i < len(drones); i++ {
		println(drones[i].String())
	}

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
	t := time.Now()
	elapsed := t.Sub(b.TimeStart)
	println(elapsed)
}

func startDrones(c int, b reality.DroneBearing, l reality.DroneLocation, t reality.Team) []reality.Drone {
	ds := make([]reality.Drone, 0)
	for i := 0; i < c; i++ {
		newDrone := reality.DroneCreate(l, b, reality.DRONEISR, t.Name(), 20)
		ds = append(ds, newDrone)
	}
	return ds
}
