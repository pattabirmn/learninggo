package main

import "fmt"

func main() {
	type user struct {
		ID   int
		name string
	}
	var u1 user
	u1.ID = 1
	u1.name = "John Doe"

	fmt.Println(u1)

	u2 := user{ID: 3, name: "Jane Doe"}
	u3 := user{ID: 2,
		name: "Harry Potter", //use comma at the last param as well, when using multi line initilizers
	}

	fmt.Println(u2, u3)
}
