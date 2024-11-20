package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type employee struct {
	Name   string `json : "EmployeeName"`
	EmpId  string
	Salary int      `json:"-"`
	Domain []string `json:"technology,omitempty"`
}
type Region struct {
	RegionID   int       `json:"region_id"`
	RegionName string    `json:"region_name"`
	employee   *employee `json: "employee"`
}

func main() {
	region := Region{RegionID: 1,
		RegionName: "Europe"}
	jsonData, _ := json.Marshal(region)
	fmt.Println("Json -", reflect.TypeOf(jsonData))

	fmt.Println("Json encoding")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	empList := []employee{
		{"Sandeep", "cs123", 1000, []string{"java", "springboot"}},
		{"Aman", "cs456", 2000, []string{"Angular", "react"}},
		{"Naman", "cs111", 3000, nil},
	}
	fmt.Println("Data is ", empList)
	empjson, err := json.MarshalIndent(empList, "", "\t") //convert to json
	if err != nil {
		panic(err)
	}
	fmt.Println("Data is ", string(empjson))

}

func DecodeJson() {
	jsonData := []byte(`
	    {
	      "Name": "Sandeep",
                "EmpId": "cs123",
                "technology": [
                        "java",
                        "springboot"
                ]
        }
	 `)
	var empData employee

	if json.Valid(jsonData) {
		fmt.Println("Got valid data")
		json.Unmarshal(jsonData, &empData)
		fmt.Printf("%#v\n", empData)
	} else {
		fmt.Println("Got In-valid data")
	}

	var empMap map[string]interface{}

	if json.Valid(jsonData) {
		fmt.Println("Got valid data")
		json.Unmarshal(jsonData, &empMap)
		fmt.Printf("%#v\n", empMap)
	} else {
		fmt.Println("Got In-valid data")
	}

	for k, v := range empMap {
		fmt.Printf("Key - %v,Value - %v, Type - %T\n", k, v, v)
	}
}
