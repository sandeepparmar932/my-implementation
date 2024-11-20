package apiofemployee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	t.Run("Successful Case", func(t *testing.T) {
		emp := Employee{
			Ename:  "Sandeep",
			Salary: 1000,
			ID:     1,
		}
		val, err := emp.createEmployee(emp)
		assert.NoError(t, err)
		assert.Equal(t, emp, val)
	})

	t.Run("Failed Case", func(t *testing.T) {
		emp1 := Employee{
			Salary: 1000,
			ID:     1,
		}
		_, err1 := emp1.createEmployee(emp1)
		assert.Error(t, err1, "abc")
		assert.Equal(t, err1.Error(), "Please provide all data")
	})

}

func TestGetEmployee(t *testing.T) {
	t.Run("Successful Case", func(t *testing.T) {
		emp, err1 := Employee{}.getEmployee("Sandeep")
		assert.NotNil(t, emp)
		assert.NoError(t, err1)
		t.Log("Employee is", emp)
		assert.Equal(t, "Sandeep", emp.Ename, "The names should be same")

	})

	t.Run("Fail Case", func(t *testing.T) {
		_, err1 := Employee{}.getEmployee("Naman")
		assert.Error(t, err1)
		assert.Equal(t, "No Data found with this name", err1.Error())
	})

}

func TestUpdateEmployee(t *testing.T) {
	t.Run("Successful Case", func(t *testing.T) {
		emp := Employee{
			Ename:  "Sandeep",
			Salary: 9999,
			ID:     1,
		}
		val, err := emp.updateEmployee(emp, "Sandeep")
		assert.NoError(t, err)
		assert.Equal(t, emp, val)
		assert.Equal(t, 9999, emp.Salary)
	})

	t.Run("Failed Case", func(t *testing.T) {
		emp1 := Employee{
			Salary: 1000,
			ID:     1,
		}
		_, err1 := emp1.createEmployee(emp1)
		assert.Error(t, err1)
		assert.Equal(t, err1.Error(), "Please provide all data")
	})

}

func TestDeleteEmployee(t *testing.T) {
	t.Run("Successful Case", func(t *testing.T) {
		err := Employee{}.deleteEmployee("Sandeep")
		assert.NoError(t, err)
	})

	t.Run("Failed Case", func(t *testing.T) {
		err1 := Employee{}.deleteEmployee("Naman")
		assert.Error(t, err1)
		assert.Equal(t, err1.Error(), "Please provide all data")
	})

}
