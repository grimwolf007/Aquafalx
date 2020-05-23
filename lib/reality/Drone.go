package reality

import "fmt"

//DRONEISR : ISR Drone
const DRONEISR = 0

//DRONEA2A : AirToAirDrone
const DRONEA2A = 1

//DRONEA2G : AirToGround
const DRONEA2G = 2

//DRONEHybrid : Hybrid A2A-A2G Drone
const DRONEHybrid = 3

var droneIDCount = 0

// Cargo : maintains what are on the Drone at a given time
type Cargo struct {
	radar    int
	fueltank int
	payloads map[string]int
}

// CargoCreate : maintains drone cargo
func CargoCreate(r int, f int, p map[string]int) Cargo {
	c := Cargo{radar: r, fueltank: f, payloads: p}
	return c
}

// DroneLocation : maintains drone coordinates
type DroneLocation struct {
	x int
	y int
	z int
}

// DroneLocationCreate : maintains drone coordinates
func DroneLocationCreate(_x int, _y int, _z int) DroneLocation {
	d := DroneLocation{x: _x, y: _y, z: _z}
	return d
}

// DroneBearing : Orientation of drone
type DroneBearing struct {
	pitch int
	yaw   int
	roll  int
}

// DroneBearingCreate : Constructor of the bearing struct
func DroneBearingCreate(p int, y int, r int) DroneBearing {
	b := DroneBearing{pitch: p, yaw: y, roll: r}
	return b
}

// Drone : structure to hold data for drones
type Drone struct {
	id       string
	team     string
	loca     DroneLocation
	bear     DroneBearing
	dtype    int
	pitch    int
	yaw      int
	roll     int
	cargo    Cargo
	maxSpeed int
}

//DroneCreate : Creates a new Drone
func DroneCreate(l DroneLocation, b DroneBearing, droneType int, _team string, _maxSpeed int) Drone {
	//check if team exists
	//if location is nil use default for drone type
	//if bearing is nill use middle of map
	newDrone := Drone{id: fmt.Sprint(droneIDCount), loca: l, dtype: droneType, bear: b, team: _team, maxSpeed: _maxSpeed}
	droneIDCount++
	return newDrone
}

//PayloadMapCreate : Creates a payload map based on Drone type

//String : displays information about a drone object
func (d Drone) String() string {
	str := "id: " + fmt.Sprint(d.id) + "\tteam: " + d.team
	return str
}
