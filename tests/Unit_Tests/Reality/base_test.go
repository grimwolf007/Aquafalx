package tests

import (
	"Aquafalx/lib/reality"
	"testing"
)

//https://golang.org/pkg/testing/

func TestBase(t *testing.T) {
	// Create bunker to be tested
	println("Creating Bunker")
	var team reality.Team = reality.TeamCreate("Red")
	var name string = "Bunker test 1"
	var ipaddr string = "192.168.0.1"
	var ipport int = 4000
	var base reality.Base = reality.BaseCreate(team, name, reality.BaseBUNKER, ipaddr, ipport)

	// Does it have the correct data?
	println(reality.BaseToString(base))

	if base.Name() != "Bunker test 1" {
		t.Errorf("Base name is not Bunker test 1, instead '%s'", base.Name())
	}

	if base.BaseType() != "Bunker" {
		t.Errorf("Base Type is not Bunker, instead '%s'", base.BaseType())
	}
	if base.IPAddress() != "192.168.0.1" {
		t.Errorf("Base Address is not 192.168.0.1, instead '%s'", base.IPAddress())
	}
	if base.IPPort() != 4000 {
		t.Errorf("Base Port is not 4000, instead '%d'", base.IPPort())
	}
	// Edit base Name
	println("Changing Name")
	base.SetName("testing")
	if base.Name() != "testing" {
		t.Errorf("Base name is not testing, instead '%s'", base.Name())
	}
	// Is it in the team roster?

	// Is it placed somewhere?

	// Does it take damage?
}
