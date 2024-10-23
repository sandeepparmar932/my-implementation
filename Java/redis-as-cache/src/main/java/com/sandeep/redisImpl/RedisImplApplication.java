package com.sandeep.redisImpl;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;

import java.util.*;
import java.util.stream.Collectors;

@SpringBootApplication
@EnableCaching
public class RedisImplApplication {

	public static void main(String[] args) {

		SpringApplication.run(RedisImplApplication.class, args);


	}

}
