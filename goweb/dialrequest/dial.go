/*
Dial reads whats going on at the url passed through the tcp
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {

	// net.Dial takes (networktype and address string) returns (connection and error)
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// ioutil.RealAll takes (connection) returns (bytestring and error)
	bs, _ := ioutil.ReadAll(conn)
	fmt.Printf("%s", bs)
}
