# Medication API

This project is a Medication API built with Go, MongoDB, and Docker. It provides endpoints to manage medications, including creating, reading, updating, and deleting medication records.

## Prerequisites

Before getting started, make sure you have the following prerequisites installed on your system:

- [Go](https://golang.org/dl/): The Go programming language.
- [Docker](https://www.docker.com/get-started): Docker is required if you wish to run the application in a container.

## Installation

Follow the steps below to install the project in your development environment:

1. **Clone the repository:**

   ```
   git clone https://github.com/MarcosMorelli/medications-api.git
   ```

2. **Navigate to the project directory:**

   ```
   cd medications-api
   ```

3. **Build the application using Docker Compose:**

   ```
   docker compose up
   ```

## Running the Application

After installation, you can run the MedicationsAPI Go application with the following command (if you want to run it directly with Golang):

```
docker compose up -d mongo

go run main.go
```

The application will be accessible at `http://localhost:8080`.

## Testing the Application

The MedicationsAPI Go application offers REST endpoints for creating, listing, updating, and deleting users. You can use tools like [curl](https://curl.se/) or [Postman](https://www.postman.com/) to test the endpoints. Here are some `curl` command examples for testing the endpoints:

- **Create a medication:**

  ```
  curl -X POST http://localhost:8080/v1/medications -H "Content-Type: application/json" -d '{"name": "Paracetamol", "dosage": "500 mg", "form": "Tablet"}'
  ```

- **Update a medication:**

  ```
  curl -X PUT http://localhost:8080/v1/medications/{id} -H "Content-Type: application/json" -d '{"name": "Paracetamol", "dosage": "500 mg", "form": "Tablet"}'
  ```

- **Delete a medication:**

  ```
  curl -X DELETE http://localhost:8080/v1/medications/{id}
  ```

Remember to adjust the commands according to your needs and requirements.

## Features

- **NoSQL database**: [MongoDB](https://www.mongodb.com) object data modeling using [Go Mongo Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- ~~**Validation**: request data validation.~~ TBD
- **Health Check**: service and ~~db~~ (TBD) health check
- **Logging**: using [golang native sLog](https://pkg.go.dev/log/slog)
- **Testing**: unit and ~~integration~~ TBD
- ~~**Error handling**: centralized error handling mechanism~~ TBD
- ~~**API documentation**: with Swagger~~ TBD
- ~~**Automatic dependencies update**: with [Dependabot](https://docs.github.com/pt/code-security/dependabot)~~ TBD
- **Environment variables**: using [golang native os](https://pkg.go.dev/os)
- ~~**Security**: set security HTTP headers~~ TBD
- ~~**Santizing**: sanitize request data against xss and query injection~~ TBD
- ~~**CORS**: Cross-Origin Resource-Sharing enabled~~ TBD
- ~~**Compression**: gzip compression~~ TBD
- **CI**: continuous integration with [Github Actions](https://docs.github.com/en/actions)
- **Docker support**
- ~~**Code coverage**: TBD