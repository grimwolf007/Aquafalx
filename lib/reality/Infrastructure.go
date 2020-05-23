package reality

import (
	"net"
)

//Infrastructure : Contains the information to connect infrastructure to physical computer
type Infrastructure interface {
	Name() string
	Type() string
	IpAddress() string
	IpPort() int
}

//IPv4Check : Checks if a string is a valid IPv4 address
func IPv4Check(ipAddress string) bool {
	testInput := net.ParseIP(ipAddress)
	if testInput.To4() == nil {
		// 	fmt.Printf("%v is not a valid IPv4 address\n", testInput)
		return false
	}
	// fmt.Printf("%v is a valid IPv4 address\n", testInput)
	return true
}
