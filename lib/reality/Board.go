package reality

import (
	"time"
)

//Board : Contains each team
type Board struct {
	name      string
	TimeStart time.Time
	TimeEnd   time.Time
	Teams     map[string]Team
}

//BoardCreate : creates a new gameboard
func BoardCreate(n string) Board {
	b := Board{name: n}
	b.TimeStart = time.Now()
	return b
}

//Status : returns the time since the board was created
func (b Board) Status() string {
	t := time.Now()
	elapsed := t.Sub(b.TimeStart)
	return elapsed.String()
}

//teamCheck : checks if a team exists
func (b Board) teamCheck(t Team) bool {
	_, found := b.Teams[t.name]
	return found
}
