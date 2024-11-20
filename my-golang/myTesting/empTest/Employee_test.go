package empTest

import (
	"errors"
	"testing"

	"github.com/go-playground/assert"
)

type MockEmpRepo struct {
	emp     *Employee
	mockErr error
}

func (m *MockEmpRepo) FindByID(id int) (*Employee, error) {
	if m.mockErr != nil {
		return nil, m.mockErr
	}
	return m.emp, nil
}

func TestFindByID(t *testing.T) {
	mockEmp := &Employee{
		Id:     1,
		Name:   "Sandeep",
		Salary: 1000,
	}

	mockRepo := &MockEmpRepo{
		emp:     mockEmp,
		mockErr: nil, // No error in this case
	}

	empService := newEmpService(mockRepo)
	emp, err := empService.FindByID(1)

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	assert.Equal(t, emp, mockEmp)
}

func TestFindByID_Error(t *testing.T) {
	mockEmp := &Employee{
		Id:     1,
		Name:   "Sandeep",
		Salary: 1000,
	}

	mockRepo := &MockEmpRepo{
		emp:     mockEmp,
		mockErr: errors.New("employee not found"),
	}

	empService := newEmpService(mockRepo)
	emp, err := empService.FindByID(1)

	if err == nil || err.Error() != "user not found" {
		t.Fatalf("expected 'user not found' error, but got %v", err)
	}
	assert.Equal(t, emp, mockEmp)

}
