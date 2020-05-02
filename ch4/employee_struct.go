package ch4

import (
	"fmt"
	"time"
)

type Employee struct {
	ID, Salary, ManagerID   int
	Name, Address, Position string
	DoB                     time.Time
}

func newEmployee(id, salary, managerid int, name, address, position string, t time.Time) *Employee {
	return &Employee{
		ID:        id,
		Salary:    salary,
		ManagerID: managerid,
		Name:      name,
		Address:   address,
		Position:  position,
		DoB:       t,
	}
}
func TestEmployeeStruct() {
	dilbert := newEmployee(2020, 6000, 1995, "name", "address", "position", time.Now())
	yuri := Employee{2020, 6000, 1995, "name", "address", "position", time.Now()}

	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior" + *position
	fmt.Println(dilbert)
	fmt.Println(yuri)
}
