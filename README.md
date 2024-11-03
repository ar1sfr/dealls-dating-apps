# Dealls Dating Apps

This is a simple RESTful API service for sign-up and login using Go and MongoDB.

## Project Structure

dealls-dating-apps/
├── handlers/
│   ├── auth.go
│   └── auth_test.go
├── models/
│   └── user.go
├── utils/
│   └── db.go
├── main.go
└── README.md

- **handlers/**: Contains the HTTP handlers for sign-up and login.
- **models/**: Contains the data models used in the service.
- **utils/**: Contains utility functions, including the MongoDB connection setup.
- **main.go**: The entry point of the application.
- **README.md**: This file.

## Instructions to Run the Service

### Prerequisites

- Go 1.18 or later
- MongoDB

### Steps

1. Clone the repository:

```sh
git clone https://github.com/ar1sfr/dealls-dating-apps

cd dealls-dating-apps
```

2. Install dependencies:

```sh
go mod tidy
```

3. Start MongoDB server.

4. Run the application:

```sh
go run main.go
```

The server will start on port 9000.

API Endpoints

GET /: Check the server status.

POST /signup: Sign up a new user.

POST /login: Log in an existing user.

Running Tests

To run the tests, use the following command:

```sh
go test ./handlers
```