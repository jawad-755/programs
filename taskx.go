package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
)

type Student struct {
	rollNo   int
	name     string
	address  string
	subjects []string
}

func NewStudent(rollNo int, name string, address string, subjects []string) *Student {
	s := new(Student)
	s.rollNo = rollNo
	s.name = name
	s.address = address
	s.subjects = subjects
	return s
}

// func NewStudent(rollNo int, name string, address string) *Student {
// 	s := new(Student)
// 	s.rollNo = rollNo
// 	s.name = name
// 	s.address = address
// 	return s
// }

type StudentList struct {
	students []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, subjects []string) *Student {
	s := NewStudent(rollno, name, address, subjects)
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
		fmt.Println("Subjects the student is studying is: ", s.subjects)
		counter++
	}
}

func (ls StudentList) String() string {
	var buffer bytes.Buffer
	for _, s := range ls.students {
		buffer.Reset()

		buffer.WriteString(fmt.Sprintf("Roll No: %d, Name: %s, Address: %s, Subjects: %s", s.rollNo, s.name, s.address, s.subjects))
		sum := sha256.Sum256([]byte(buffer.String()))
		fmt.Println(buffer.String())

		fmt.Printf("%x\n", sum)
		fmt.Println()
	}

	return buffer.String()
}

func main() {
	ls := new(StudentList)
	ls.CreateStudent(1, "Jawad", "Pakistan",[]string{"MATH", "PHY", "CHEM"})
	ls.CreateStudent(2, "Jude", "UK",[]string{"PF", "OOP", "COAL"})
	ls.CreateStudent(3, "Cristiano", "Portugal",[]string{"SST", "CRYPTO", "INFO"})

	ls.String()
	//ls.PrintStudentDetails()

}