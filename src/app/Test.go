package main

import "Aquafalx/lib/reality"

func main() {
	var team reality.Team = reality.TeamCreate("Red")
	var name string = "Bunker test 1"
	var ipaddr string = "192.168.0.1"
	var ipport int = 4000
	var baseBunker reality.Base = reality.BaseCreate(team, name, reality.BaseBUNKER, ipaddr, ipport)
	//sanity check
	println("\nTeam: \n", team.Name)
	println("Bunkers: \n")
	println(reality.BaseToString(baseBunker))

}
