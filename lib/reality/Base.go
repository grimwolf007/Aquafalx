package reality

import (
	"fmt"
	"math"
	"strconv"
)

const(
//BaseBUNKER : Where users "are"
	BaseBUNKER = iota

//BaseSAM : Ground based Anti-Air station
	BaseSAM

//BaseRADAR : Ground based Radar for unit detection
	BaseRADAR

//BaseAIRSTRIP : Used to house and launch aircraft
	BaseAIRSTRIP

//BaseFUELDEPOT : Used to store Fuel to transport to other bases
	BaseFUELDEPOT

//BaseAMMODEPOT : Used to store ammunitions to transport to other bases
	BaseAMMODEPOT

//BaseTRANSDEPOT : Used to store land vehicles for transporting resources and offense
	BaseTRANSDEPOT

//BasePORT : Used to store sea vehicles for transporting resources and offense
	BasePORT
)
// Base : Generic base type
type Base struct {
	name      string
	baseType  int
	ipAddress string
	ipPort    int
	loca      Location
	drones    []*Drone
}

//BaseCreate : creates a base
func BaseCreate(t *Team, baseType int, ipAddress string, ipPort int, l Location, name string) Base {

	var base Base
	if IPv4Check(ipAddress) {
		base := Base{name: name, baseType: baseType, ipAddress: ipAddress, ipPort: ipPort, loca: l}
		t.AddBase(base)
		return base
	}
	return base
}

// SetName : Sets the name of the base, returns true if changed
func (b *Base) SetName(n string) bool {
	if n != "" {
		b.name = n
		return true
	}
	return false
}

// Name : returns the name of the base
func (b Base) Name() string {
	return b.name
}

// Type : returns the type of the base
func (b Base) Type() string {
	switch b.baseType {
	case BaseBUNKER:
		return "Bunker"
	case BaseSAM:
		return "SAM"
	case BaseRADAR:
		return "RADAR"
	case BaseAIRSTRIP:
		return "Air Strip"
	case BaseFUELDEPOT:
		return "Fuel Depot"
	case BaseAMMODEPOT:
		return "Ammo Depot"
	case BaseTRANSDEPOT:
		return "Transport Depot"
	case BasePORT:
		return "Port"
	default:
	}
	return "MISSINGNO."
}

//IPAddress : returns the IP address of the base
func (b Base) IPAddress() string {
	return b.ipAddress
}

//IPPort : returns the IP port of the base main service
func (b Base) IPPort() int {
	return b.ipPort
}

//String : formats the values of a base in the form of a string
func (b Base) String() string {
	c := 0
	var str = b.name + ": " + "\n\tIP Address: \t" + b.ipAddress + "\n\tIP Port: \t" + strconv.Itoa(b.ipPort) + "\n\tType: \t" + b.Type() + "\n\tLocation: \t" + b.Location()
	str = str + "\n\tDrones:"
	for _, d := range b.drones {
		str = str + "\n\t\t" + fmt.Sprint(c) + " [" + d.String() + "]"
	}
	return str
}

//Location : returns location of base as a string
func (b Base) Location() string {
	str := "x: " + fmt.Sprint(b.loca.x) + "\ty: " + fmt.Sprint(b.loca.y)
	return str
}

//Loca : returns location of base as a location struct
func (b Base) Loca() Location {
	return b.loca
}

//ChangeLocation : changes the location of a base
func (b *Base) ChangeLocation(x float64, y float64) bool {
	loca := LocationCreate(x, y, 0)
	//if out of bounds return false
	b.loca = loca
	return true
}

func dist(b Base, d Drone) float64 {

	x1 := d.loca.x
	x2 := b.loca.x
	y1 := d.loca.y
	y2 := b.loca.y
	z1 := d.loca.z
	z2 := b.loca.z
	dist := math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2) + math.Pow((z1-z2), 2)
	math.Sqrt(dist)
	return dist
}

//AddDrone : Adds a drone to a base
func (b *Base) AddDrone(d *Drone) bool {

	dist := dist(*b, *d)
	//println(fmt.Sprint(dist))
	if dist < 4 {
		b.drones = append(b.drones, d)
		return true
	}
	return false
}

//Drones : displays the drones inside of the base
func (b *Base) Drones() string {
	var str string
	for i, d := range b.drones {
		str = str + "\n" + fmt.Sprint(i) + "\t" + d.String()
	}
	return str
}

//RemoveDrone : removes drone from base and places it just outside the base collision box
func (b *Base) RemoveDrone(i int) *Drone {
	if len(b.drones) > i {
		d := b.drones[i]
		copy(b.drones[i:], b.drones[i+1:])
		b.drones[len(b.drones)-1] = nil
		b.drones = b.drones[:len(b.drones)-1]
		d.loca.x = b.loca.x + 1
		d.loca.y = b.loca.y + 1
		d.loca.z = 0
		return d
	}
	return nil
}
