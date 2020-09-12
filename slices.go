package main

import "fmt"

func main() {

	// copy existing array into a slice
	// slice and array will be pointing to same location. changes to slice/array reflect in other type
	strArr := [3]string{"one", "two", "three"}
	strSlice := strArr[:]
	fmt.Println(strSlice)

	//now add more elements to the slice using append
	strSlice = append(strSlice, "four", "five")
	fmt.Println(strSlice)

	// initiating slice without explicitly copying an existing array. compiler takes care of it
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 6, 5)
	fmt.Println(slice)

	//copying slices/ slicing
	//[startingIndex:endingIndex] startingIndex - inclusive. endingIndex - exclusive
	s2 := slice[3:]  // slices "slice" from 3rd index till end (including 3rd element) - [4 5 6 5]
	s3 := slice[:5]  // slices "slice" from the beginning till the 5th index (excluding 6th element) [1 2 3 4 5]
	s4 := slice[1:5] // slices "slice" from 1st element till 5th element [2 3 4 5]
	s5 := slice[0:7] // copies everything

	fmt.Println(slice, s2, s3, s4, s5)

}
