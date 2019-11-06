package reality

//BASES
const baseBUNKER = 0
const baseSAM = 1
const baseRADAR = 2
const baseAIRstrip = 3
const baseFUELdepot = 4
const baseAMMOdepot = 5
const baseTRANSdepot = 6
const basePORT = 7

//Team : Contains all the infrastructure information for each team
type Team struct {
	hq    Base
	bases []Base
}

type resource struct {
	name string
	num  int
}

//Base : Contains the information to connect base to physical computer
type Base struct {
	name      string
	baseType  int
	ipAddr    string
	ipPort    int
	resources []resource
}

//TeamCreate : creates a new team and assets required
func TeamCreate(name string) Team {
	team := Team{}
	return team
}
