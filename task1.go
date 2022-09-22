package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	job string
	salary int
	}

	func FunctionName( x Person) {
		fmt.Println("Name: ", x.name)
		}

func main() {
	var pers1 Person
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	

	FunctionName(  pers1)

}
