package main

import (
	"fmt"
	"reality"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	var team reality.Team = reality.TeamCreate("Red")
	var name string = "Bunker test 1"
	var ipaddr string = "192.168.0.1"
	var ipport int = 4000
	var baseBunker reality.Base = reality.BaseCreate(team, name, reality.BaseBUNKER, ipaddr, ipport)
	//sanity check
	println("\nTeam: \n", team.Name)
	println("Bunkers: \n")
	println(reality.BaseTostring(baseBunker))

}
