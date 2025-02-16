package com.example.demo.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController  // Marks this class as a REST Controller
@RequestMapping("/api")  // Base URL for this controller
public class HelloController {

    @GetMapping("/hello")  // Maps HTTP GET request to /api/hello
    public String sayHello() {
        return "Hello, Spring Boot!";
    }
}
