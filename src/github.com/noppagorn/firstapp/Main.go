package main

import "fmt"

// var i float32 = 42

// this can use
// var actorName string = "Elisabeth Sladen"
// var companion string = "Sarah Jane Smith"
// var doctorNumber int = 3
// var season int = 11

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	count int = 0
)

var i int = 27

// uppper case variable can export to be a global variable
func main() {
	var j float32
	j = 100
	var i int = 27 // easy assign
	var seasonNumber int = 11
	// i := 42. // init and assign
	fmt.Printf("%v, %T\n", j, j)

	fmt.Printf("%v, %T\n", i, i, seasonNumber)
	// var i int = 32
	fmt.Println("Hello Go 111")

	fmt.Println(i, j)

}
