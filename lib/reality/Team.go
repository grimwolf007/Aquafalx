package reality

//Team : Contains all the infrastructure information for each team
type Team struct {
	name  string
	hq    Base
	bases map[string]Base
}

//TeamCreate : creates a new team and assets required
func TeamCreate(n string) Team {
	t := Team{name: n, bases: make(map[string]Base)}
	return t
}

// Name : returns the name of the Team
func (t Team) Name() string {
	return t.name
}

// HQ : returns the HQ of the Team
func (t Team) HQ() Base {
	return t.hq
}

// Bases : returns all bases on the team
func (t Team) Bases() map[string]Base {
	return t.bases
}

// SetHQ : Sets a base as the Head Quarters
func (t *Team) SetHQ(h Base) {
	t.hq = h
	t.bases[h.Name()] = h

}
