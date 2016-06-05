package main

import (
	"fmt"
	"log"
)

type rooterr struct {
	lat, long string
	err       error
}

func (n *rooterr) Error() string {
	return fmt.Sprintf("a root error has occured with : %v", n.err)
}

func main() {
	_, err := squareroot(-10)
	if err != nil {
		log.Println(err)
	}
}
func squareroot(f float64) (float64, error) {
	if f < 0 {
		myrooterr := fmt.Errorf("You got a rooterr due to : %v", f)

		// we can consider this an error, because this struct implements an error! amazing!!
		return 0, &rooterr{"lat : 54.24315", "long : 34.523465", myrooterr}
	}
	return 42, nil
}
