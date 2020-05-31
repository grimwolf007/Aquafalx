package reality

import (
	"fmt"
	"math"
	"sync"
	"time"
)

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

// Bearing : Orientation of drone
type Bearing struct {
	pitch float64
	yaw   float64
	roll  float64
}

// BearingCreate : Constructor of the bearing struct
func BearingCreate(p float64, y float64, r float64) Bearing {
	b := Bearing{pitch: p, yaw: y, roll: r}
	return b
}

// Drone : structure to hold data for drones
type Drone struct {
	id       string
	team     string
	loca     Location
	bear     Bearing
	dtype    int
	pitch    int
	yaw      int
	roll     int
	cargo    Cargo
	maxSpeed float64
	flyLock  [1]bool
}

//DroneCreate : Creates a new Drone
func DroneCreate(l Location, b Bearing, droneType int, _team string, _maxSpeed float64) Drone {
	//check if team exists
	//if location is nil use default for drone type
	//if bearing is nill use middle of map
	newDrone := Drone{id: fmt.Sprint(droneIDCount), loca: l, dtype: droneType, bear: b, team: _team, maxSpeed: _maxSpeed}
	droneIDCount++
	return newDrone
}

//PayloadMapCreate : Creates a payload map based on Drone type [WIP]
func PayloadMapCreate() {}

//Type : returns the type of drone
func (d Drone) Type() string {
	switch d.dtype {
	case DRONEA2A:
		return "Air to Air"
	case DRONEA2G:
		return "Air to Ground"
	case DRONEHybrid:
		return "Hybrid"
	case DRONEISR:
		return "Recon"
	default:
		return "unknown"
	}
}

// FlyTo : Drone flies directly to location, returns true if successful
func (d *Drone) FlyTo(l Location, wg *sync.WaitGroup) bool {
	lock := d.flyLock
	vec := make(map[string]float64)
	var mag float64
	//kill old FlyTo Instruction
	lock[0] = false
	time.Sleep(100 * time.Millisecond)
	lock[0] = true
	//loop till at the final location
	for !d.locaEqual(l) && lock[0] == true {
		//Find vector
		vec["x"] = float64(l.x - d.loca.x)
		vec["y"] = float64(l.y - d.loca.y)
		vec["z"] = float64(l.z - d.loca.z)
		//calculate vector magnitude
		mag = math.Sqrt(math.Pow(vec["x"], 2) + math.Pow(vec["y"], 2) + math.Pow(vec["z"], 2))
		//if vector magnitude < unit magnitude go to location
		if mag < d.maxSpeed {
			d.loca = l
		} else {
			//else vector/magnitude * speed magnitude
			vec["x"] = vec["x"] / mag * d.maxSpeed
			vec["y"] = vec["y"] / mag * d.maxSpeed
			vec["z"] = vec["z"] / mag * d.maxSpeed
			//unit vector + location
			d.loca.x = d.loca.x + vec["x"]
			d.loca.y = d.loca.y + vec["y"]
			d.loca.z = d.loca.z + vec["z"]
		}
		println(d.String())
	}

	//free flight lock
	lock[0] = false
	wg.Done()
	if d.locaEqual(l) {
		return true
	}
	return false
}

func (d Drone) locaEqual(l Location) bool {
	if d.loca.x == l.x && d.loca.y == l.y && d.loca.z == l.z {
		return true
	}
	return false

}

//String : displays information about a drone object
func (d Drone) String() string {
	str := "id: " + fmt.Sprint(d.id) + "\tteam: " + d.team + "\ttype: " + d.Type()
	str = str + "\tlocation:" + fmt.Sprint(d.loca.x) + "," + fmt.Sprint(d.loca.y) + "," + fmt.Sprint(d.loca.z)
	return str
}
