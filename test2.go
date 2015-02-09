package main

import "fmt"

const (
	Course     = "Computer Science"
	University = "Califonia State Univesity, Fresno"
)

type Student struct {
	name string
	age  int
	id   int
}

func main() {

	var c = Student{"Name", 22, 1235}

	StudentInformation(c)

}

func StudentInformation(student Student) {
	fmt.Println(student.name)
	fmt.Println(student.age)
	fmt.Println(student.id)

	fmt.Println(Course)
	fmt.Println(University)

	fmt.Println(ReturnFunction())
	name, _ := ReturnFunction()
	fmt.Println(name)
}

func ReturnFunction() (string, int) {
	return "Henna", 15
}
