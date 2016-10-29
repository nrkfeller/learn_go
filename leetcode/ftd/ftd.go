package ftd

// func main() {
//
// 	s := "abcd"
// 	t := "abcde"
//
// 	fmt.Printf("%s", Ftd(s, t))
//
// }

func Ftd(a, b string) string {

	atotal := 0
	btotal := 0
	for _, k := range a {
		atotal += int(k)
	}
	for _, k := range b {
		btotal += int(k)
	}

	return string(btotal - atotal)
}
