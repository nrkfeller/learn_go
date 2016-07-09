// These examples demonstrate more intricate uses of the flag package.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
)

var ip = flag.Int("ipflag", 1234, "help message for flagname")

var flagvar int

func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
func main() {
	flag.Parse()
	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

	server := flag.String("server", "tcp://iot.eclipse.org:1883", "The MQTT server to connect to")
	room := flag.String("room", "gochat", "The chat room to enter. default 'gochat'")
	name := flag.String("name", "user"+strconv.Itoa(rand.Intn(1000)), "Username to be displayed")
	fmt.Println(*server, *room, *name)
}
