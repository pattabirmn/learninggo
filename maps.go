package main

import "fmt"

func main() {
	strMap := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(strMap)

	strMap["four"] = 4
	strMap["five"] = 5
	fmt.Println(strMap)
	fmt.Println(strMap["one"])

	delete(strMap, "four")
	fmt.Println(strMap)

}
