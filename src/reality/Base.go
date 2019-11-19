package reality

import (
	"fmt"
	"net"
	"strconv"
)

//BaseBUNKER : Where users are
const BaseBUNKER = 0

//BaseSAM : Ground based Anti-Air station
const BaseSAM = 1

//BaseRADAR : Ground based Radar for unit detection
const BaseRADAR = 2

//BaseAIRstrip : Used to house and launch aircraft
const BaseAIRstrip = 3

//BaseFUELdepot : Used to store Fuel to transport to other bases
const BaseFUELdepot = 4

//BaseAMMOdepot : Used to store ammunitions to transport to other bases
const BaseAMMOdepot = 5

//BaseTRANSdepot : Used to store land vehicles for transporting resources and offense
const BaseTRANSdepot = 6

//BasePORT : Used to store sea vehicles for transporting resources and offense
const BasePORT = 7

//BaseCreate : creates a base
func BaseCreate(t Team, n string, b int, i string, p int) Base {
	// name comes from random word list
	// ipAddr comes from list generated by initial cidr notation
	// ipPort is generated by initial random list separated by type
	if IPv4Check(i) {
		base := Base{name: n, baseType: b, ipAddr: i, ipPort: p}
		return base
	}
	return Base{}
}

//BaseTostring : formats the values of a base in the form of a string
func BaseTostring(b Base) string {
	var str = b.name + ": " + "\n\tIP Address: \t" + b.ipAddr + "\n\tIP Port: \t" + strconv.Itoa(b.ipPort)
	return str
}

//IPv4Check : Checks if a string is a valid IPv4 address
func IPv4Check(ipAddress string) bool {
	testInput := net.ParseIP(ipAddress)
	if testInput.To4() == nil {
		fmt.Printf("%v is not a valid IPv4 address\n", testInput)
		return false
	}
	fmt.Printf("%v is a valid IPv4 address\n", testInput)
	return true
}