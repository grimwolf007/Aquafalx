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

// AddBase : adds a new base to the team roster
func (t *Team) AddBase(h Base) bool {
	if _, ok := t.bases[h.Name()]; ok {
		println("%s is already on $s", h.Name(), t.Name())
		return false
	}
	t.bases[h.Name()] = h
	return true
}
