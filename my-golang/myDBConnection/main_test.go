package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		emp := Employee{
			Ename:  "Sandeep",
			Salary: 10000,
		}
		payload, err := json.Marshal(emp)
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, "/employee", bytes.NewBuffer(payload))
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateEmployee)
		handler.ServeHTTP(rr, request)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		var createdEmp Employee
		json.NewDecoder(rr.Body).Decode(&createdEmp)
		if createdEmp.ID <= 0 {
			t.Errorf("expected a CourseId to be generated, got empty string")
		}

		assert.Equal(t, emp.Salary, createdEmp.Salary)
		assert.Equal(t, emp.Ename, createdEmp.Ename)

	})
}
