package empTest

type Employee struct {
	Id     int
	Name   string
	Salary int
}

type EmpRepo interface {
	FindByID(int) (*Employee, error)
}

type EmpService struct {
	repo EmpRepo
}

func newEmpService(repo EmpRepo) *EmpService {
	return &EmpService{
		repo: repo,
	}
}

func (e *EmpService) FindByID(id int) (*Employee, error) {
	return e.repo.FindByID(id)
}
