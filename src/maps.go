package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//map literals
// ommited type for each element in map as Vertex is just a type name
var locations = map[string]Vertex{
	"London": {
		51.5074,
		0.1277,
	},
	"Manchester": {
		53.4808,
		2.2426,
	},
}

func main() {
	// Simple map with make
	m := make(map[string]Vertex)
	m["London"] = Vertex{
		51.5074,
		0.1277,
	}
	fmt.Println(m)

	fmt.Println(locations)
	// Update Map
	m["Birmingham"] = Vertex{
		1,
		2,
	}
	fmt.Println("Updated map: ", m)

	// Test that key is in map
	key := "New York"
	elem, ok := m[key]

	fmt.Printf("The value: %v Present in map? %v\n", elem, ok)

	//Delete an element from map
	delete(m, "Birmingham")
	fmt.Println(m)
}
