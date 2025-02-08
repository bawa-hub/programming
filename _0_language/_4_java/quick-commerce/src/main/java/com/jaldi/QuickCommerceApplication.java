package com.jaldi;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@EnableJpaRepositories("com.jaldi.repository")
public class QuickCommerceApplication {

	public static void main(String[] args) {
		SpringApplication.run(QuickCommerceApplication.class, args);
	}

}

@RestController
@RequestMapping("/api")
class TestController {
    @GetMapping("/hello")
    public String hello() {
        return "Hello, Spring Boot is running!";
    }
}