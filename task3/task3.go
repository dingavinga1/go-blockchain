package main

import (
	"fmt"
	"strings"
)

type Student struct{
	rollnumber int
	name string
	address string
}

func (st *Student)print(){
	fmt.Println("student rollno\t", st.rollnumber)
	fmt.Println("student name\t", st.name)
	fmt.Println("student address\t", st.address)
}

func NewStudent(rollno int, name string, address string) *Student{
	s:=new(Student)
	s.rollnumber=rollno
	s.name=name
	s.address=address
	return s
}

type StudentList struct{
	list []*Student
}

func (ls *StudentList)CreateStudent(rollno int, name string, address string) *Student{
	st:=NewStudent(rollno, name, address)
	ls.list=append(ls.list, st)
	return st
}

func (ls *StudentList)print(){
	for index, element:=range ls.list{
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), index, strings.Repeat("=", 25))
		element.print()
	}
}

func main(){
	student:=new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")
	student.print()
}