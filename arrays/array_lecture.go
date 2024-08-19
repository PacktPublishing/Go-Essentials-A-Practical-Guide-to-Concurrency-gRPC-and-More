package arrays

import "fmt"

func Learn(){
	var array [3]int

	array[0] = 0
	array[1] = 1
	array[2] = 2

	//literal := [2]int{3,4}

	// fmt.Println(array)
	// fmt.Println(literal)

	for _, val := range array {
		fmt.Println(val)
	}

}