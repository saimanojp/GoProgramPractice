package main

import "fmt"

func main() {
	var basicSalary, hra, da, grossSalary float64

	fmt.Print("Enter Basic Salary of Employee")
	fmt.Scanln(&basicSalary)

	if basicSalary <= 10000 {
		hra = (basicSalary * 8) / 100
		da = (basicSalary * 10) / 100

	} else if basicSalary <= 20000 {
		hra = (basicSalary * 16) / 100
		da = (basicSalary * 20) / 100
	} else {
		hra = (basicSalary * 24) / 100
		da = (basicSalary * 30) / 100
	}

	grossSalary = basicSalary + hra + da
	fmt.Println("Gross Salary of Employee is ", grossSalary)

}
