# Users Management System

A go microservice that can enables us to create, modify, fetch, and delete users

## Usage

### To Run the application
```bash
go run main.go serve
```

### To Run unit tests
```bash
go test ./...
```

### To Run the application with custom environment variables
```bash
go run main.go serve --env STAGE --host 0.0.0.0 --port 4000
```