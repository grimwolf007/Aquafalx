package reality

//Team : Contains all the infrastructure information for each team
type Team struct {
	Name        string
	hq          Base
	bases       []Base
	baseCounter []int //uses the constant name -> int relation for index of array
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
func TeamCreate(n string) Team {
	team := Team{Name: n}
	return team
}
