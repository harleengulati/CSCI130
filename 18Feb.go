package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i = i + 5 {
		fmt.Println(i)
	}
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}

		if i > 10 {
			fmt.Println("break")
			break
		}
	}

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for j, sPointer := range mySlice {
		fmt.Println(j, " - ", sPointer)
	}
	fmt.Print("[5:]")
	fmt.Println(mySlice[1:])

	//string slicing
	slice := "Harleen Slice !!!"
	fmt.Println(slice)
	fmt.Println(slice[3:])

}
