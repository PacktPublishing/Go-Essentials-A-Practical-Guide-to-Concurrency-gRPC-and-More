package assignment1

import "fmt"

// Create a map via a map literal that uses an int as a key and a string as the value.
// Create as many values as you’d like, but there must be a minimum of three key:value pairs.
// After, create a function that accepts the map and then use a for loop to delete each value out of the map.
// Call the delete function and then print out the map to confirm all values were deleted.

func CreateMap(){
	mappy := map[string]int{"one":1, "two": 2, "five": 5}
	fmt.Println(mappy)
	DeleteMap(mappy)
	fmt.Println(mappy)
}

func DeleteMap(mapArg map[string]int){
	for key, _ := range mapArg {
		delete(mapArg, key)
	}
}