package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Student struct{
	rollnumber int
	name string
	address string
	subjects []string
}

func (st *Student)print(){
	fmt.Println("student rollno\t", st.rollnumber)
	fmt.Println("student name\t", st.name)
	fmt.Println("student address\t", st.address)
}

func (st *Student)getHash() string{
	toHash:=fmt.Sprintf("%d%s%s", st.rollnumber, st.name, st.address)
	for _, element:=range st.subjects{
		toHash+=element
	}

	hash := sha256.Sum256([]byte(toHash))
	sum := fmt.Sprintf("%x", hash)

	return sum
}

func NewStudent(rollno int, name string, address string, subs []string) *Student{
	s:=new(Student)
	s.rollnumber=rollno
	s.name=name
	s.address=address
	s.subjects=subs
	return s
}

type StudentList struct{
	list []*Student
}

func (ls *StudentList)CreateStudent(rollno int, name string, address string, subjects []string) *Student{
	st:=NewStudent(rollno, name, address, subjects)
	ls.list=append(ls.list, st)
	return st
}

func (ls *StudentList)print(){
	for index, element:=range ls.list{
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), index, strings.Repeat("=", 25))
		element.print()
		fmt.Println("Hash: ", element.getHash())
	}
}

func main(){
	student:=new(StudentList)

	subArr:=[]string{"Chemistry", "Biology"}
	student.CreateStudent(24, "Asim", "AAAAAA", subArr)
	student.CreateStudent(25, "Naveed", "BBBBBB", subArr)
	student.print()
}