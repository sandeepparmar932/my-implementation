/*
package com.sandeep.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;



@RequestMapping(path = "Employee")
public class EmployeeController {

	@Autowired
	EmployeeService emp;
	
	@GetMapping("/findAllInOrder")
	public List<Employee> findAllInOrder(@PathVariable String field) {
		return emp.findAllInOrder(field);
	}

	@GetMapping("/findDataByPaging")
	public List<Employee> findDataByPaging(@PathVariable int offset,@PathVariable int page) {
		return emp.findDataByPaging(offset,page);
	}
	
}
*/
