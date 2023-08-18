package main

import (
	"fmt"
)

func main() {
	list := make([]int, 0)
	list = append(list, 1,2,3,4)

	func(list []int) {
		list = append(list[:1], list[2:]...)
		fmt.Println(" func: ",list)
		//list = make([]int, 0)
		for i := 0; i < 3; i++ {
			list[i] = 0
		}
	}(list)
	fmt.Println(list)

	//reflect.SliceHeader{}
}