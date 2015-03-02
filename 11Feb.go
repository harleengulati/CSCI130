package main

import "fmt"

type Harleen struct {
	number int
	phone  int
}

type MyPrinterType func(string)

func Greet(person Harleen, myWassa MyPrinterType) {
	myGreetingMas := CreateMessage(person.number, person.phone, 1234)
	myWassa(myGreetingMas)
}

func CreateMessage(age int, welcomeMas ...int) (myGreeting string) {
	switch "Blah" {
	case "Tim":
		myGreeting = ("Deffered Message =>")
	case "Blah":
		myGreeting = ("Added Message =>")
	}
	if welcomeMas[1] != 0 {
		myGreeting = myGreeting + " Message 1"
	} else {
		myGreeting = myGreeting + " Message 2"
	}
	return
}

func myPrintFunction(custom string) MyPrinterType {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func main() {
	var t = Harleen{21, 32}
	Greet(t, myPrintFunction("!!!"))
}
