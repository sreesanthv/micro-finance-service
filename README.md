# Account Management Service (Go)
This project is a Golang-based service for managing user accounts and transactions. It provides a clean abstraction layer using interfaces for repository and service layers, making it suitable for extension and integration with various storage backends and APIs.

## Features
- Create, retrieve, update, and delete user accounts
- Fetch accounts by ID or Email
- Save transaction data associated with an account
- Clean separation of concerns via repository and service interfaces

## Data Model

```go
type Account struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Pan         string `json:"pan"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Sex         string `json:"sex"`
	Nationality string `json:"nationality"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
```

## Getting Started
## Start Service
- docker-compose up
- Open pgAdmin - http://localhost:5050/login?next=%2F (Username: admin@admin.com, Password: root)
- Please refer config.yaml for application credentials.
- Run init.sql in DB
- go run main.go serve
- Please refers to POSTMAN collection micro-finance-service.postman_collection.json
