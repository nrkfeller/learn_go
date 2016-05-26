package main

import (
	"encoding/json"
	"fmt"
)

// Message must capitalize to export in JSON!
type Message struct {
	ID          int
	Body        string
	notExported int
}

func main() {
	m1 := Message{1, "hello world", 007}
	bs, _ := json.Marshal(m1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
