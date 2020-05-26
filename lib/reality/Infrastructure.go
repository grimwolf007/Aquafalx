package reality

import (
	"bufio"
	"net"
	"os"
)

var scanner bufio.Scanner

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

//GetBaseName : gives default name to infrastructure
func GetBaseName(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

//CreateBaseNameScanner : creates a scanner for the base random name file
func CreateBaseNameScanner(path string) *bufio.Scanner {
	_file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(_file)
	scanner.Split(bufio.ScanLines)
	return scanner
}
