package main

import "fmt"
import "github.com/goestoeleven/math"

type Massage string

// TYPES CAN CONTAIN METHODS
func (s Massage) mePrint() string {
	return "METHOD: Printing Message"
}

func main() {

	nSlice := []int{15, 100, 1115}

	sONumbers := math.Sum(nSlice)

	fmt.Println(sONumbers)

	mSlice := []int{1, 5, 15, 20, 25, 30}
	mOSlice := []int{50, 70, 90}

	mSlice = append(mSlice, mOSlice...)

	for i, cEntry := range mSlice {
		fmt.Println(i, " - ", cEntry)
	}

	fmt.Println("")

	mSlice = append(mSlice[:4], mSlice[5:]...)

	for i, cEntry := range mSlice {
		fmt.Println(i, " - ", cEntry)
	}

	var massage Massage = "Test For Work"
	fmt.Println(massage.mePrint())

}
