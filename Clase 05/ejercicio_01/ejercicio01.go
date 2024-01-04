package main

import "fmt"

type SalaryError struct {
}

func (s SalaryError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}

func checkSalary(salary int) error {
	var err error = nil
	if salary < 150000 {
		err = SalaryError{}
	}
	return err
}

func main() {
	var salary int
	fmt.Println("Insert your salary")
	fmt.Scanf("%d", &salary)
	err := checkSalary(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Must pay a tax.")
}
