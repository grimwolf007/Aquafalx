package reality

import (
	"bufio"
	"fmt"
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
	Location() string
	ChangeLocation() bool
}

//Location : Stuct that contains the x and y coord of infrastructure
type Location struct {
	x float64
	y float64
	z float64
}

//LocationCreate : Creates an instance of a location
func LocationCreate(x float64, y float64, z float64) Location {
	loca := Location{x: x, y: y, z: z}
	return loca
}

//Change : Change the location it is refering to
func (l *Location) Change(x float64, y float64, z float64) {
	l.x = x
	l.y = y
	l.z = z
}
func (l Location) String() string {
	str := "x: " + fmt.Sprint(l.x)
	str = str + ", y: " + fmt.Sprint(l.y)
	str = str + ", z: " + fmt.Sprint(l.z)
	return str
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
