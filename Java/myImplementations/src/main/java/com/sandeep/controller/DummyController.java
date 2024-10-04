package com.sandeep.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.LinkedHashSet;

@RestController
@ResponseBody
public class DummyController {

    @GetMapping("/get")
    public String getMethod() {
        return "This is dummy Response";
    }

    @PostMapping("/post")
    public ResponseEntity<Object> postMethod(@RequestBody Object obj) {

        return  new ResponseEntity<>(obj, HttpStatus.OK);

    }
}
