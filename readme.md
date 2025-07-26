1. docker-compose up -d
2. run sql script to create tables - /database/sql/init.sql
3. go mod tidy
4. make a copy of .env from example.env


System flow
- main -> router -> controller -> service -> repo -> model -> database

API documentation - Swagger
http://localhost:8080/swagger/index.html

Sample postman collection is included
