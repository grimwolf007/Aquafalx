package main

import (
	"Aquafalx/lib/reality"
	"bufio"
	"fmt"
	"sync"
)

func main() {
	//make a flight wait group
	println(info("Setting up variables"))
	println(info("Setting up flight wait group"))
	var flightWaitGroup sync.WaitGroup

	//for the random name gen
	randomBaseNameFilePath := "/home/john/go/src/Aquafalx/lib/reality/BaseNames.txt"
	println(info("Using " + fmt.Sprint(randomBaseNameFilePath) + "to generate random base names"))
	scanner := reality.CreateBaseNameScanner(randomBaseNameFilePath)

	//creates the board
	println(info("Creating board Test1"))
	b := reality.BoardCreate("Test1")

	//creates a location and bearings for drones
	println(info("Creating location pointer at 4,4,0"))
	l := reality.LocationCreate(4, 4, 0)
	println(info("Creating Bearing at 0,0,0"))
	bear := reality.BearingCreate(0, 0, 0)

	//makes a team
	println(info("Making Red and Blue Teams"))
	teams := startTeams([]string{"Red", "Blue"})

	//adds 2 bases
	//Red team
	team := &teams[0]
	println(info("adding a base to " + team.Name()))
	redbases := addBases(team, scanner, l)

	println(info("Changing location pointer to 8,9,0"))
	l.Change(8, 9, 0)
	//Blue team
	team = &teams[1]
	println(info("adding a base to " + team.Name()))
	bluebases := addBases(team, scanner, l)

	//makes 10 drones per team
	team = &teams[0]
	l.Change(4, 4, 1)
	println(info("putting 10 drones near the location: " + l.String()))
	count := 10
	drones := startDrones(count, bear, l, team)

	basesTeam0 := teams[0].Bases()
	base := basesTeam0["hi"]

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

//addBases : adds an airstrip to a team for testing purposes
func addBases(t *reality.Team, randomNameScanner *bufio.Scanner, l reality.Location) []reality.Base {
	bases := make([]reality.Base, 0)
	bases = append(bases, reality.BaseCreate(t, reality.BaseAIRSTRIP, "127.0.0.1", 9999, l, "", randomNameScanner))
	for i := 0; i < len(bases); i++ {
		t.AddBase(bases[i])
	}
	t.SetHQ(bases[0])
	return bases
}

//buildInfrastructure : creates the infrastructure for the scenario [WIP]
func buildInfrastructure() {}

//startDrones : creates drones in bulk
func startDrones(c int, b reality.Bearing, l reality.Location, t *reality.Team) []*reality.Drone {
	ds := make([]*reality.Drone, 0)
	for i := 0; i < c; i++ {
		newDrone := reality.DroneCreate(l, b, reality.DRONEISR, t.Name(), 20)
		ds = append(ds, &newDrone)
	}
	return ds
}

//https://gist.github.com/ik5/d8ecde700972d4378d87
var (
	info    = teal
	warn    = yellow
	fata    = red
	black   = color("\033[1;30m%s\033[0m")
	red     = color("\033[1;31m%s\033[0m")
	green   = color("\033[1;32m%s\033[0m")
	yellow  = color("\033[1;33m%s\033[0m")
	purple  = color("\033[1;34m%s\033[0m")
	magenta = color("\033[1;35m%s\033[0m")
	teal    = color("\033[1;36m%s\033[0m")
	white   = color("\033[1;37m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
