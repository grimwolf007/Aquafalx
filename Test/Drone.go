package reality

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

//ISRDroneCreate : Creates a new ISR Drone
func ISRDroneCreate() {}

//AirToAirDroneCreate : Creates a new Air to Air Drone
func AirToAirDroneCreate() {}

//AirToGroundDroneCreate : Creates a new Air to Ground Drone
func AirToGroundDroneCreate() {}

//HybridDroneCreate : Creates a new Hybrid A2A-A2G Drone
func HybridDroneCreate() {}
