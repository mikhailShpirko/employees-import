# Employees Import
A practice task to implement employees import from CSV using Golang and applying Clean Architecture


## Launch instructions

### Docker Compose

Execute the following command to run the project to play around with it:

```
docker compose -f deploy/EmployeesImport.DockerCompose.yaml -p employees_import up -d --build
```

After that you will be able to test the API using the attached postman collection with base URL http://localhost:9898

Execute the following command to shut down the project:

```
docker compose -f deploy/EmployeesImport.DockerCompose.yaml -p employees_import down
```

### Local

Create .env file in src folder from the example below

```
DATABASE_CONNECTION_STRING=postgres://DB_USERNAME:DB_PASSWORD@DB_HOST:DB_PORT/DB_NAME
PORT=3000
```

In terminal navigate to sc folder and run

```
go run main.go
```

Given that you indicate same port as in example .env file, the API will be available at base URL http://127.0.0.1:3000. You will be able to test the API using the attached postman collection by adjusting baseUrl accordingly.

## Testing notes

Use `dataset.csv` file for testing importing employees via file upload - in postman collection navigate to `Import Employees from CSV` and selected the file