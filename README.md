# 🏢 OrgManager – Employee Management System (GoLang CRUD API)

A simple and clean RESTful API built with Go to manage employees in an organization. This project is ideal for learning, showcasing your Go skills, or bootstrapping a small HR/employee management backend.

---

## 🚀 Features

- Add new employees
- View all employees
- Get a specific employee by ID
- Update employee details
- Delete employee
- Auto-generated SQLite database via GORM

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gorilla Mux (Routing)
- **ORM:** GORM
- **Database:** SQLite (auto-generated file)
- **Tools:** VSCode, Git, GitHub

---

## 📁 Folder Structure

```bash
orgmanager/
├── config/
│   └── db.go                 # DB connection using GORM
│
├── controllers/
│   └── employeeController.go # Logic for each CRUD endpoint
│
├── models/
│   └── employee.go           # Employee data model
│
├── routes/
│   └── routes.go             # HTTP route mappings using Gorilla Mux
│
├── main.go                   # Entry point – initializes DB and starts server
├── go.mod                    # Go module definition
├── go.sum                    # Go dependencies checksum
└── README.md                 # This file
```

---

## 📌 API Endpoints

| Method | Endpoint           | Description         |
|--------|--------------------|---------------------|
| GET    | `/employees`       | Get all employees   |
| GET    | `/employees/{id}`  | Get employee by ID  |
| POST   | `/employees`       | Create new employee |
| PUT    | `/employees/{id}`  | Update employee by ID |
| DELETE | `/employees/{id}`  | Delete employee by ID |

---
## 💡 How to Run Locally

### Clone the repository:
```bash
git clone https://github.com/VenkatRajendraGuntupalli/EmpManagement-golang-crud.git
cd EmpManagement-golang-crud
```
## Initialize Go modules:
```bash
go mod tidy
```
## Run the app:
```bash
go run main.go
```
## Test the API:
Use Postman or cURL
```bash
Server runs at: http://localhost:8080
```
