package main

import (
	"fmt"
	"sort"
)

func main() {

	ms := []MetroStation{
		{"angrignon", 4},
		{"dls", 3},
		{"berri", 1},
		{"snowdon", 5},
		{"mt royal", 2},
	}

	sort.Sort(byImportance(ms))

	fmt.Println(ms)

}

// MetroStation like sub
type MetroStation struct {
	name       string
	importance int
}

type byImportance []MetroStation

// Methods required by the sort package
func (l byImportance) Len() int           { return len(l) }
func (l byImportance) Less(i, j int) bool { return l[i].importance < l[j].importance }
func (l byImportance) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
