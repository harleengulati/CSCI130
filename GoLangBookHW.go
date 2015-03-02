package main

import (
	"fmt"
	"math"
	"time"

	"Web_Programming_CSUF/Web_Programming_CSUF/src/github.com/goestoeleven/math"
)
import "strconv"

func main() {
	//Temprature change
	FToC()
	//Feet to Meter
	FtoM()
	//Divisible by three check
	DivisibleByThree()
	//Check Fizz Buzz
	FizzBuzz()
	//Checking for smalest
	FSmall()
	//Boolean check for half value
	fmt.Println("Checking for 4 ", hbCheck(4))
	//Variadic function for getting highest in a list
	fmt.Println("Greatest of 1,2,12,12,54,151,21,15,121,2165,5000 => ", gList(1, 2, 12, 12, 54, 151, 21, 15, 121, 2165, 5000))

	functOdd := MOG()
	fmt.Println("MakeOddGenerator generates", functOdd())
	//Recursive Fibonacci
	fmt.Println("Fibonacci sequence for 10 => ", Recfib(10))

	//Swapping function call start
	x := 1
	y := 2
	swapVal(&x, &y)
	fmt.Println("X =", x, "Y =", y)
	//Swapping function Ends

	//Ch9
	myCircle := Circle{0, 0, 10}
	myRectangle := Rectangle{0, 0, 10, 10}
	fmt.Println("Perimeter of Circle is", myCircle.perimeter())
	fmt.Println("Perimeter of Rectangle is", myRectangle.perimeter())
	//Ch9 end

	//Sleep Function
	fmt.Println("Sleep Function Started")
	SleepProcess(1000)
	fmt.Println("Ending Sleep Function")

	minMaxFuncUsingPackage()
}

func FToC() {
	var far string
	fmt.Println("far plz ")
	fmt.Scanln(&far)
	fToInt, _ := strconv.ParseFloat(far, 64)
	fToInt = ((fToInt - 32.0) * (5.0 / 9.0))
	fmt.Println("Celsius :", fToInt)
}

func FtoM() {
	var feet string
	fmt.Println("Provide Feet to Convert")
	fmt.Scanln(&feet)
	feettoInt, _ := strconv.ParseFloat(feet, 64)
	feettoInt = (feettoInt * 0.3048)
	fmt.Println("Meters =>", feettoInt)
}

func DivisibleByThree() {
	for i := 3; i < 100; i = i + 3 {
		fmt.Println(i)
	}
}

func FizzBuzz() {
	fmt.Println("FizzBuzz Stated")
	myNum := 1.0
	for myNum <= 100 {
		if math.Mod(myNum, 3) == 0 {
			if math.Mod(myNum, 5) == 0 {
				fmt.Println("FizzBuzz")
			}
			fmt.Println("Fizz")
		} else if math.Mod(myNum, 5) == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(myNum)
		}
		myNum += 1
	}
}

func FSmall() {
	myArray := []int{12, 25, 326, 265, 15641654, 156313, 13216546310}
	minValue := myArray[0]
	for i := 1; i < len(myArray); i++ {
		if myArray[i] < minValue {
			minValue = myNumArray[i]
		}
	}
	fmt.Print("minValue =>")
	fmt.Println(minValue)
}

func hbCheck(myHalfBool float64) bool {
	myHalfBool = myHalfBool / 2
	boolValue := false
	if math.Mod(myHalfBool, 2) == 0 {
		boolValue = true
	}
	return boolValue
}

func gList(myArgList ...int) int {
	greatest := 0
	for _, v := range myArgList {
		if greatest < v {
			greatest = v
		}
	}
	return greatest
}

func MOG() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func Recfib(fibInt int) int {
	if fibInt == 0 || fibInt == 1 {
		return fibInt
	} else {
		return (Recfib(fibInt-1) + Recfib(fibInt-2))
	}
}

func swapVal(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

//Golang chapter 9

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	return ((r.y2 - r.y1) * (r.x2 - r.x1))
}
func (r *Rectangle) perimeter() float64 {
	return (2 * ((r.y2 - r.y1) + (r.x2 - r.x1)))
}

//End chapter 9

func SleepProcess(timeInMs time.Duration) {
	<-time.After(timeInMs * time.Second / 1000)
}

func minMaxFuncUsingPackage() {
	myIntArray := []float64{1, 2, -1, 20, 3, 4}
	avg := math.Average(myIntArray)
	fmt.Println(avg)
	fmt.Println(math.Min(myIntArray))
	fmt.Println(math.Max(myIntArray))
}
