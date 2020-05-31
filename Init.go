package main

import (
	"Aquafalx/lib/reality"
	"bufio"
	"sync"
)

func main() {
	//make a flight wait group
	var flightWaitGroup sync.WaitGroup

	//for the random name gen
	scanner := reality.CreateBaseNameScanner("/home/john/go/src/Aquafalx/lib/reality/BaseNames.txt")

	//creates the board
	b := reality.BoardCreate("Test1")

	//creates a location and bearings for drones
	l := reality.LocationCreate(4, 4, 0)
	bear := reality.BearingCreate(0, 0, 0)

	//makes a team
	teams := startTeams([]string{"Red", "Blue"})

	//adds 2 bases
	addBases(teams[0], scanner, l)
	l.Change(8, 9, 0)
	addBases(teams[1], scanner, l)

	//makes 10 drones
	l.Change(4, 4, 1)
	count := 10
	drones := startDrones(count, bear, l, teams[0])

	basesTeam0 := teams[0].Bases()
	base := basesTeam0["Hickory"]

	// print base
	println(base.String())

	//changes the location of the Hickory base
	base.ChangeLocation(5, 5)
	println(base.String())

	// print drones
	for i := 0; i < len(drones); i++ {
		println(drones[i].String())
	}

	// add a drone to base
	println("Drone added?")
	println(base.AddDrone(drones[0]))
	println(base.String())

	// remove drone from base
	println("drone removed?")
	drone := base.RemoveDrone(0)
	if drone == nil {
		println("Drone not found")
	} else {
		println("Drone removed")
		drones = append(drones, drone)
	}
	println(base.String())
	for i := 0; i < len(drones); i++ {
		println(drones[i].String())
	}

	// fly drone
	flightWaitGroup.Add(1)
	l.Change(20, 20, 20)
	println("flying to location: [" + l.String() + "] \n" + drone.String())
	go drones[0].FlyTo(l, &flightWaitGroup)
	go drones[1].FlyTo(l, &flightWaitGroup)
	flightWaitGroup.Wait()
	// board status
	b.Status()
}

//startTeams : creates the teams for the scenario
func startTeams(teamNames []string) []reality.Team {
	teams := make([]reality.Team, 0)
	for i := 0; i < len(teamNames); i++ {
		newTeam := reality.TeamCreate(teamNames[i])
		teams = append(teams, newTeam)
	}
	return teams
}

//addBases : adds bases to a team
func addBases(t reality.Team, randomNameScanner *bufio.Scanner, l reality.Location) reality.Team {
	bases := make([]reality.Base, 0)
	bases = append(bases, reality.BaseCreate(t, reality.BaseAIRSTRIP, "127.0.0.1", 9999, randomNameScanner, l))
	for i := 0; i < len(bases); i++ {
		t.AddBase(bases[i])
	}
	t.SetHQ(bases[0])
	return t
}

//buildInfrastructure : creates the infrastructure for the scenario
func buildInfrastructure() {}

//startDrones : creates drones in bulk
func startDrones(c int, b reality.Bearing, l reality.Location, t reality.Team) []*reality.Drone {
	ds := make([]*reality.Drone, 0)
	for i := 0; i < c; i++ {
		newDrone := reality.DroneCreate(l, b, reality.DRONEISR, t.Name(), 20)
		ds = append(ds, &newDrone)
	}
	return ds
}
