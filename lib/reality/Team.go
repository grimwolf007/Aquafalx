package reality

import "fmt"

//Team : Contains all the infrastructure information for each team
type Team struct {
	name         string
	hq           Base
	droneIDCount int
	bases        map[string]Base
	drones       map[int]*Drone
}

//TeamCreate : creates a new team and assets required
func TeamCreate(n string) Team {
	t := Team{name: n, bases: make(map[string]Base), drones: make(map[int]*Drone), droneIDCount: 0}
	return t
}

// Name : returns the name of the Team
func (t Team) Name() string {
	return t.name
}

// AddBase : adds a base to a team
func (t *Team) AddBase(h Base) {
	t.bases[h.Name()] = h
}

// HQ : returns the HQ of the Team
func (t Team) HQ() Base {
	return t.hq
}

// Bases : returns all bases on the team
func (t Team) Bases() map[string]Base {
	return t.bases
}

// SetHQ : Sets a base as the Head Quarters of a team
func (t *Team) SetHQ(h Base) {
	t.hq = h
	t.bases[h.Name()] = h

}

// AddDrone : adds a drone to a team
func (t *Team) AddDrone(h *Drone) {
	t.drones[h.id] = h
}

// Drones : returns all drones on the team
func (t Team) Drones() map[int]*Drone {
	return t.drones
}

// DroneIDCount : returns the next drone id number for the team
func (t Team) DroneIDCount() int {
	return t.droneIDCount
}

// DroneIDInc : increments the Drone ID number for the team and returns it
func (t *Team) DroneIDInc() int {
	t.droneIDCount++
	return t.droneIDCount
}

//String : returns the items of the team as a string
func (t Team) String() string {
	str := t.name

	str = str + "\n Bases:"
	str = str + "\n\t" + fmt.Sprint(t.Bases())

	str = str + "\n Drones:"
	str = str + "\n\t" + fmt.Sprint(t.Drones())

	return str
}
