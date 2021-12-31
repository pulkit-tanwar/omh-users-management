# Users Management System

A go microservice that can enables us to create, modify, fetch, and delete users

## Usage

### To Run the application
```bash
  $ go run main.go serve
```

### To Run unit tests
```bash
  $ go test ./...
```

### To Run the application with custom environment variables
```bash
go run main.go serve --env STAGE --host 0.0.0.0 --port 4000
```

## Development
```bash
  $ make dep           # install dependencies
  $ make test          # run unit tests
  $ make cover         # run code coverage report service (http://localhost:3001)
  $ make run           # run the service
```

