package main

import (
	"fmt"
	"strconv"
	"strings"
)


type Student struct {
	rollNo int
	name string
	address string
	subject []string 
}
func NewStudent(rollNo int, name string, address string, subject1 []string) *Student {
	s := new(Student)
	s.rollNo = rollNo
	s.name = name
	s.address = address	
	s.subject = subject1

	return s
}

type StudentList struct {
	students []*Student
}
func (ls *StudentList) CreateStudent(rollno int, name string, address string, subject1 []string) *Student {
	s := NewStudent(rollno, name, address, subject1)
	ls.students = append(ls.students, s)
	return s
}
// Write a function to print the details of all the students in student list

func (ls *StudentList) PrintStudentDetails() {
	counter := 1
	for _, s := range ls.students {

		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 20), counter, strings.Repeat("=", 20))
		fmt.Println("Roll No: ", s.rollNo)
		fmt.Println("Name: ", s.name)
		fmt.Println("Address: ", s.address)
		fmt.Println("subjects: ", s.subject)

		counter++
	}
}




func (ls *StudentList) stringmaker(xx string) {
	counter := 1
	for _, s := range ls.students {

		// xx+=s.rollNo
		(s.rollNo),e := strconv.Atoi(xx)
		xx+=s.name
		xx+=s.address
	    x+=s.subject
		counter++
	}
}











func main() {
	




	ls := new(StudentList)

	
	ls.CreateStudent(1, "Jawad", "Pakistan",[]string{"MATH", "PHY", "CHEM"})
	ls.CreateStudent(2, "Jude", "UK",[]string{"PF", "OOP", "COAL"})
	ls.CreateStudent(3, "Cristiano", "Portugal",[]string{"SST", "CRYPTO", "INFO"})

	ls.PrintStudentDetails()
}