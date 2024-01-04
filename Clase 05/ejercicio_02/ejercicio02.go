package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
}

func (s SalaryError) Error() string {
	return "Error: salary is less than 10000"
}

func checkSalary(salary int) error {
	var err error = nil
	if salary < 100000 {
		err = SalaryError{}
	}
	return err
}

func main() {
	var salary int
	fmt.Println("Insert your salary:")
	fmt.Scanf("%d", &salary)
	err := checkSalary(salary)
	if errors.Is(err, SalaryError{}) {
		fmt.Println(err)
		return
	}
	fmt.Println("Must pay tax")
}
