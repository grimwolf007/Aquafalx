package reality

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

//================================================================Constants

//DRONEISR : ISR Drone
const DRONEISR = 0

//DRONEA2A : AirToAirDrone
const DRONEA2A = 1

//DRONEA2G : AirToGround
const DRONEA2G = 2

//DRONEHybrid : Hybrid A2A-A2G Drone
const DRONEHybrid = 3

//DRONEGasConsumption : How much fuel is used per tick
const DRONEGasConsumption = 1

//================================================================Variables

var droneIDCount = 0

//================================================================Structs

// Cargo : maintains what are on the Drone at a given time
type Cargo struct {
	radar    int
	fueltank int
	fuel     int
	payloads map[string]int
}

// Bearing : Orientation of drone
type Bearing struct {
	pitch float64
	yaw   float64
	roll  float64
}

// Drone : structure to hold data for drones
type Drone struct {
	id       int
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

//================================================================Constructors

// Change : changes a drones cargo
func (c *Cargo) Change(r int, f int, p map[string]int) {
	c.radar = r
	c.fueltank = f
	c.fuel = c.fueltank * 100
	c.payloads = p
	return
}

// loadCargo : adds default cargo to drone on creation
func (d *Drone) loadCargo(droneType int) error {
	var r int
	var f int
	var p map[string]int

	switch d.dtype {
	case DRONEISR:
		r = 1
		f = 1
		p = nil
	case DRONEA2A:
		r = 1
		f = 1
		p = nil
	case DRONEA2G:
		r = 1
		f = 1
		p = nil
	case DRONEHybrid:
		r = 1
		f = 1
		p = nil
	default:
		err := errors.New("Unknown Drone Type")
		return err
	}
	d.cargo = Cargo{radar: r, fueltank: f, payloads: p}
	return nil
}

// BearingCreate : Constructor of the bearing struct
func BearingCreate(p float64, y float64, r float64) Bearing {
	b := Bearing{pitch: p, yaw: y, roll: r}
	return b
}

//DroneCreate : Creates a new Drone
func DroneCreate(l Location, b Bearing, droneType int, t *Team, _maxSpeed float64) *Drone {

	newDrone := Drone{id: t.DroneIDCount(), loca: l, dtype: droneType, bear: b, maxSpeed: _maxSpeed, team: t.Name()}
	t.DroneIDInc()
	t.AddDrone(&newDrone)
	i := t.DroneIDCount() - 1
	return t.Drones()[i]
}

//PayloadMapCreate : Creates a payload map based on Drone type [WIP]
func PayloadMapCreate() {}

//================================================================Getters

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

//String : displays information about a drone object
func (d Drone) String() string {
	str := "id: " + fmt.Sprint(d.id) + "\tteam: " + d.team + "\ttype: " + d.Type()
	str = str + "\tlocation:" + fmt.Sprint(d.loca.x) + "," + fmt.Sprint(d.loca.y) + "," + fmt.Sprint(d.loca.z)
	return str
}

//locaEqual : Checks if a drones location is equal to the location specified
func (d Drone) locaEqual(l Location) bool {
	if d.loca.x == l.x && d.loca.y == l.y && d.loca.z == l.z {
		return true
	}
	return false

}

//================================================================Setters

// FlyTo : Drone flies directly to location, returns true if successful
func (d *Drone) FlyTo(l Location, wg *sync.WaitGroup) bool {
	wg.Add(1)
	defer wg.Done()
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
		time.Sleep(1 * time.Second)
	}
	//free flight lock
	lock[0] = false
	if d.locaEqual(l) {
		return true
	}
	return false
}
