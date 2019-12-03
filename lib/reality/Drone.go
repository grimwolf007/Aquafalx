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

type location struct {
	x int
	y int
	z int
}

type bearing struct {
	pitch int
	yaw   int
	roll  int
}

type drone struct {
	loca     location
	bear     bearing
	dtype    int
	pitch    int
	yaw      int
	roll     int
	payloads Payloads
	team     string
	maxSpeed int
}

func droneCreate(l location, b bearing, droneType int, _team string, _maxSpeed int) drone {
	newDrone := drone{loca: l, dtype: droneType, bear: b, team: _team, maxSpeed: _maxSpeed}
	return newDrone
}

//teamCheck : Checks if team exists
func teamCheck() {}

//DroneCreate : Creates a new Drone
func DroneCreate(dronetype int) {}

//PayloadMapCreate : Creates a payload map based on Drone type adn
