package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getName() string {
	return ft.name
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
}

func main() {
	test(fullTime{
		name:   "Aniket",
		salary: 200000,
	})

	test(contractor{
		name:         "Aniket",
		hourlyPay:    5000,
		hoursPerYear: 1200,
	})
}
