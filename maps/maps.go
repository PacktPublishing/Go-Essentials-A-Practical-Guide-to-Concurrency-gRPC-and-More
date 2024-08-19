package maps

import "fmt"

func LearnMaps(){
	mappy := make(map[string]int) //int is the key, string is value
	// mappy[5] = "five"
	// mappy[1] = "one"
	// fmt.Println(mappy)

	var ok bool
	mappy["two"] = 0
	delete(mappy, "two")
	_, ok = mappy["two"]
	fmt.Println(ok)
}