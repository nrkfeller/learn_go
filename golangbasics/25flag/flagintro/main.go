// These examples demonstrate more intricate uses of the flag package.
package main

import (
	"flag"
	"fmt"
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
}
