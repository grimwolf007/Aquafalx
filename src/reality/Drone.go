package reality

//DRONEISR : ISR Drone
const DRONEISR = 0

//DRONEA2A : AirToAirDrone
const DRONEA2A = 1

//DRONEA2G : AirToGround
const DRONEA2G = 2

//DRONEHybrid : Hybrid A2A-A2G Drone
const DRONEHybrid = 3

// Payloads : maintains what payloads are on the Drone at a given time
type Payloads struct {
	radar    int
	fueltank map[string]int
	payloads map[string]int
}

// Drone : maintains the orientation and location of a drone
type Drone struct {
	X        int
	Y        int
	Z        int
	pitch    int
	yaw      int
	roll     int
	payloads Payloads
	team     string
	maxSpeed int
}

func droneCreate(x int, y int, z int, _pitch int, _yaw int, _roll int, _team string, _maxSpeed int) Drone {
	drone := Drone{X: x, Y: y, Z: z, pitch: _pitch, yaw: _yaw, roll: _roll, team: _team, maxSpeed: _maxSpeed}
	return drone
}

//teamCheck : Checks if team exists
func teamCheck() {}

//DroneCreate : Creates a new Drone
func DroneCreate(dronetype int) {}

//PayloadMapCreate : Creates a payload map based on Drone type adn
