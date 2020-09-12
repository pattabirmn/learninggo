package main

import "fmt"

func main() {
	var i int
	i = 42
	fmt.Println(i)

	// float32 float 64
	var fl float32 = 3.14
	fmt.Println(fl)

	//implicit initialization operator :=
	name := "John Doe"
	fmt.Println(name)

	var b bool = true
	fmt.Println(b)

	//------------------------------------------------------------------------------------------
	//Pointers

	// var fname *string               // pointer data type without memory initilized . will return <nil>. can't use until initilizing
	var fname *string = new(string) //pointer data type with initializing memory allocation

	*fname = "John" // using * to dereference.

	fmt.Println(fname)  //prints the pointer. i.e. address location
	fmt.Println(*fname) //deferenced. Prints the value stored in the given pointer

	//address of operator - &

	lname := "Doe"
	lnamePtr := &lname //references the location of lname

	fmt.Println(lnamePtr, *lnamePtr)

	lname = "Dane"
	fmt.Println(lnamePtr, *lnamePtr) //location doesn't change. only the data

	//-------------------------------------------------------------------------------------------
	//Constants

	const con = 3 // initialize at declartion. can't change value afterwards. with/without type
	fmt.Println(con)
	fmt.Println(con + 2)
	fmt.Println(con + 2.2)

	const con2 int = 4 //initialized with type restriction. can't add float32 now. explicit type conversion required
	fmt.Println(con2)
	fmt.Println(con2 + 2)
	//fmt.Println(con2 + 2.2)		//won't run - constant 2.2 truncated to integer - need to convert con2 to float32 explicitly
	fmt.Println(float32(con2) + 2.2)

}
