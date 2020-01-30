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
	// Add Hq to team
	println("Creating Base HQ")
	var name string = "Red HQ"
	var ipaddr string = "127.0.0.1"
	var ipport int = 4000
	var basehq reality.Base = reality.BaseCreate(team, name, reality.BaseBUNKER, ipaddr, ipport)
	println("Adding to Team")
	team.SetHQ(basehq)

	if team.HQ().Name() != "Red HQ" {
		t.Errorf("Team HQ is not Red HQ, instead '%s'", basehq.Name())
	}
	// Is HQ in the team roster?
	println("on roster?")
	if team.Bases()["Red HQ"] != basehq {
		t.Errorf("Red HQ is not on the team roster")
	}

	// Add a base to a team

	// Add a drone to the team

	// Remove a drone from a team

	// Display status of drones

	// Display status of buildings
}
