package main

import "fmt"

func main() {
	elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	name, ok := elements["Un"]
	fmt.Println(name, ok)
}
