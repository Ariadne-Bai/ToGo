package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m) // note that go can directly print out containers

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs) // _ is blank identifier   // return a false value if key does not exist
}
