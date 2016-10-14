package astar

// Check checks in the string is good
func Check(input string) bool {
	if len(input)%2 == 0 {
		return false
	}
	for i := 0; i < len(input); i++ {
		if i%2 == 0 && string(input[i]) != "a" {
			return false
		}
		if i%2 == 1 && string(input[i]) != "*" {
			return false
		}
	}
	return true
}
