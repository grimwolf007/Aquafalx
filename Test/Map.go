package airspace

const defaultx = 100
const defaulty = 100
const defaultz = 100

//FlightMap : The map of the airspace
type FlightMap struct {
	X int
	Y int
	Z int
}

func mapCreate(x int, y int, z int) FlightMap {
	airSpace := FlightMap{X: x, Y: y, Z: z}
	return airSpace
}

func mapDefault() FlightMap {
	return mapCreate(defaultx, defaulty, defaultz)
}
