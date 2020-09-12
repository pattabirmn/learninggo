package main

import "fmt"

func main() {
	var arr [4]int //fixed size. need to have all the values. use slices for dynamic sizing?
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	fmt.Println(arr) //prints [1 2 3 4]

	strArr := [3]string{"one", "two", "three"}
	fmt.Println(strArr)

	fmt.Println(strArr[2])
}
