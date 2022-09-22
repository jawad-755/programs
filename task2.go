package main

import "fmt"

type employee struct {
	name string

	salary int

	position string
}

type company struct {
	companyName string

	employees []employee
}


func main() {
	var emp1 employee
	emp1.name = "Hege"
	emp1.salary = 45000
	emp1.position = "scratcher"

	var emp2 employee
	emp2.name = "jack"
	emp2.salary = 65000
	emp2.position = "itcher"

	var emp3 employee
	emp3.name = "frazzer"
	emp3.salary = 450000000
	emp3.position = "hitter"

	var arr1 = []employee{emp1, emp2, emp3}

	var khujlico company
	khujlico.companyName = "KhujliCo"
	khujlico.employees = arr1
	fmt.Println(khujlico.employees)
}