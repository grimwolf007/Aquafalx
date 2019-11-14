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
	var name string = "test1"
	var ipaddr string = "192.168.0.1"
	var ipport int = 4000
	var baseBunker reality.Base = reality.BaseCreate(team, name, reality.BaseBUNKER, ipaddr, ipport)
	//sanity check
	println(reality.BaseBUNKER)
	println("baseBunker: %s ", reality.BaseTostring(baseBunker))
	println(team.Name)
}