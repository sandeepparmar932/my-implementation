#About Project
It project is created to implement redis in spring-boot project. 

Redis is an open source in-memory data store that can be used as a database, cache, or message broker. It's often used for caching web pages and reducing the load on servers.
It is a NoSQL database.

For this project using jedis as redis-client over default for Spring Boot client Lettuce.

Choose Jedis if:

Your application requires a simple, synchronous API and blocking calls.
You are working on a small project or proof of concept where ease of use is a priority.
Choose Lettuce if:

You need high concurrency or are developing a reactive application.
Your application will benefit from non-blocking I/O and you want to leverage advanced Redis features.
You are working in a multi-threaded environment and need a thread-safe solution.