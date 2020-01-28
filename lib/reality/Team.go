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

//TeamCreate : creates a new team and assets required
func TeamCreate(n string) Team {
	team := Team{Name: n}
	return team
}
