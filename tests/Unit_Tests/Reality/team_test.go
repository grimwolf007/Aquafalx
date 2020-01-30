//https://golang.org/pkg/testing/
package tests

import (
	"Aquafalx/lib/reality"
	"testing"
)

func TestTeam(t *testing.T) {
	// Create a team
	println("Creating Team")
	var team reality.Team = reality.TeamCreate("Red")
	if team.Name() != "Red" {
		t.Errorf("Team name is not Red, instead '%s'", team.Name())
	}
	// Add a base to a team

	// Add a drone to the team

	// Remove a drone from a team

	// Display status of drones

	// Display status of buildings
}
