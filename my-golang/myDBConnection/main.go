package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID     int
	Ename  string
	Salary int
}

func dbOpen() *sql.DB {
	db, err := sql.Open("mysql", "root:jmahouse@2017@tcp(localhost:3306)/goTest")
	if err != nil {
		fmt.Println("Error validating sql.open")
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error validating db ping")

	}
	fmt.Println("Connected to mysql")
	return db
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", getEmployeeByID).Methods("GET")
	r.HandleFunc("/employee/{id}", updateEmployee).Methods("PUT")

	//listing to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	db := dbOpen()
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Db -", db)
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send data"})
		return
	}
	defer r.Body.Close()
	emp := &Employee{}
	json.NewDecoder(r.Body).Decode(&emp)
	if emp.Ename == "" || emp.Salary == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "PROVIDE MANDATORY DATA"})
	}
	row, err := db.Exec("Insert into employee(ename ,salary ) values (?,?)", emp.Ename, emp.Salary)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error is ", err)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Error While executing query"})
		return
	}
	id, err := row.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error is ", err)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Unable to fetch last inserted ID"})
		return
	}
	result := db.QueryRow("SELECT id, ename, salary FROM employee WHERE id = ?", id)
	insertedEmp := &Employee{}
	err = result.Scan(&insertedEmp.ID, &insertedEmp.Ename, &insertedEmp.Salary)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error is ", err)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Error While fetching data"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(insertedEmp)
	defer db.Close()
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if r.Body == nil || params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send data"})
		return
	}
	defer r.Body.Close()
	db := dbOpen()
	defer db.Close()
	emp := &Employee{}
	json.NewDecoder(r.Body).Decode(&emp)
	err := db.QueryRow("Update Employee set  ename=$1, salary = $2 where id=$3  RETURNING id, ename, salary ", emp.Ename, emp.Salary, params["id"]).Scan(&emp.Ename, &emp.Salary, &emp.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Error While executing query"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM employee WHERE id = ?"
	params := mux.Vars(r)
	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send data"})
		return
	}

	employee := &Employee{}
	db := dbOpen()
	defer db.Close()
	row := db.QueryRow(query, params["id"])
	err := row.Scan(&employee.ID, &employee.Ename, &employee.Salary)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "employee with ID not found"})
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(employee)
}

// func (Employee) TableName() string {
// 	return "Employee" // this will use "Employee" as the table name instead of "employees"
// }

// // using goRM
// func main() {
// 	dbString := "root:jmahouse@2017@tcp(localhost:3306)/goTest"

// 	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})
// 	if err != nil {
// 		//fmt.Println(err)
// 		panic(err)
// 	}

// 	db.AutoMigrate(&Employee{}) // will update or create the table acc. to struct
// 	//create
// 	db.Create(&emp)
// 	//read
// 	var employee Employee
// 	db.First(&employee, 5)
// 	fmt.Println("Employee based on id - ", employee.Ename, employee.Salary)
// 	employee = Employee{}
// 	db.First(&employee, "ename=?", "Manish")
// 	fmt.Println("Employee based on salary - ", employee.Ename, employee.Salary)

// 	// Update
// 	db.Model(&employee).Update("Salary", 60000)
// 	db.Model(&employee).Updates(Employee{Salary: 99000})
// 	fmt.Println("Employee Updated salary - ", employee.Ename, employee.Salary)
// 	db.Model(&employee).Updates(map[string]interface{}{"EName": "Lalit"})
// 	fmt.Println("Updated employee  - ", employee.Ename, employee.Salary)

// 	// Delete
// 	db.Delete(&employee, 1)

// }
