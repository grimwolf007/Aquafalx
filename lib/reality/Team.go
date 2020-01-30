package reality

//Team : Contains all the infrastructure information for each team
type Team struct {
	name        string
	hq          Base
	bases       []Base
	baseCounter []int //uses the constant name -> int relation for index of array
}

//TeamCreate : creates a new team and assets required
func TeamCreate(n string) Team {
	t := Team{name: n}
	return t
}

// Name : returns the name of the Team
func (t Team) Name() string {
	return t.name
}
