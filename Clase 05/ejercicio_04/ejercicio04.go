package main

import (
	"fmt"
)

func checkSalary(salary int) error {
	var err error = nil
	if salary < 150000 {
		err = fmt.Errorf("Error: the minimum taxable amount is $150,000 and the salary entered is: $%d", salary)
	}
	return err
}

func main() {
	var salary int
	fmt.Println("Insert your salary:")
	fmt.Scanf("%d", &salary)
	err := checkSalary(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Must pay tax")
}
