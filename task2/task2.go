package main

import "fmt"


type employee struct{
	name string
	salary int
	position string
}

type company struct {
	companyName string 
	employees []employee
}

func printDetails(arg company){
	fmt.Printf("====%s====\n", arg.companyName)
	for index, element:=range arg.employees{
		fmt.Printf("Employee %d\n", index)
		fmt.Printf("---------\n")
		fmt.Printf("Name: %s\n", element.name)
		fmt.Printf("Salary: %d\n", element.salary)
		fmt.Printf("Position: %s\n\n", element.position)
	}
}

func main(){
	emplys:=  []employee{
		employee{
			"Abdullah Irfan",
			600000,
			"Senior Devops Engineer",
		},
		employee{
			"Aisha Irfan",
			700000,
			"Director Engineering",
		},
		employee{
			"Muhammad Huzaifa",
			50000,
			"Clerk",
		},
	}

	company1:= company{"Secpose", emplys}

	printDetails(company1)



}