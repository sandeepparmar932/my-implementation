package apiofemployee

import "errors"

type Employee struct {
	ID     int
	Ename  string
	Salary int
}

var myData []Employee

func (Employee) createEmployee(emp Employee) (Employee, error) {
	if emp.Ename == "" || emp.Salary == 0 {
		return Employee{}, errors.New("Please provide all data")
	}
	myData = append(myData, emp)
	return emp, nil

}

func (Employee) getEmployee(name string) (Employee, error) {
	if name == "" {
		return Employee{}, errors.New("PLEASE PROVIDE DATA")
	}

	for _, employee := range myData {
		if employee.Ename == name {
			return employee, nil
		}
	}

	return Employee{}, errors.New("No Data found with this name")
}

func (Employee) updateEmployee(emp Employee, name string) (Employee, error) {
	if name == "" {
		return Employee{}, errors.New("PLEASE PROVIDE DATA")
	}

	for i, employee := range myData {
		if employee.Ename == name {
			myData = append(myData[:i], myData[i+1:]...)
			myData = append(myData, emp)
			return emp, nil
		}
	}

	return Employee{}, errors.New("No Data found with this name")
}

func (Employee) deleteEmployee(name string) error {
	if name == "" {
		return errors.New("PLEASE PROVIDE DATA")
	}

	for i, employee := range myData {
		if employee.Ename == name {
			myData = append(myData[:i], myData[i+1:]...)
			return nil
		}
	}

	return errors.New("No Data found with this name")
}
