# 🏥 Clinic Management System API

## 📘 API Documentation

All API endpoints are documented in Postman.  
🔗 [View Live API Documentation](https://documenter.getpostman.com/view/33373776/2sB2xBCVFd)



## Description 
A backend API built using **Golang** and the **Gin** framework for managing clinic operations involving **Doctors**, **Receptionists**, and **Patients**. The project is structured  around clean architecture principles with a focus on:

- ✅ **Repository Pattern** for abstraction
- 🔐 **JWT authentication**
- 🧑‍⚕️ **Role-based access control**
- ⚡ **Redis** integration for session caching and rate limiting

---

## 🚀 Features

- Doctor & Receptionist registration and login
- Patient management (add, view, update)
- JWT-based authentication with Redis storage
- Role-based access to secure routes
- Redis-ready for rate limiting
- Clean repository structure for scalable codebase
- Cookie-based auth token with expiry and security

---

## 🧾 API Endpoints

| Method | Endpoint                                | Description                                                  | Access Role   |
|--------|------------------------------------------|--------------------------------------------------------------|---------------|
| POST   | `/auth/register`                         | Register as a doctor or receptionist (unique email required) | Public        |
| POST   | `/auth/login`                            | Login for both doctors and receptionists                     | Public        |
| POST   | `/auth/logout`                           | Logout user by clearing the token                            | Authenticated |
| GET    | `/doctor/patients`                       | Retrieve all patients                                        | Doctor        |
| POST   | `/receptionist/registerPatient`          | Add a patient (requires name and disease)                    | Receptionist  |
| PUT    | `/doctor/patients/:id`                   | Update a patient by ID (ID in URL param)                     | Doctor        |

---

## 🏛️ Project Structure
```
ClinicManagement/
├── cmd/
│   └── main/
│       └── main.go                       # entry point 
├── controllers/
│   ├── auth.go
│   ├── doctor.go
│   └── receptionist.go
├── initializers/
│   ├── database.go
│   ├── loadEnvVariables.go
│   └── redis.go
├── middlewares/
│   ├── isDoctor.go
│   ├── isReceptionist.go
│   └── ratelimiter.go
├── migrate/
│   └── migrate.go
├── models/
│   ├── employee.go
│   └── patient.go
├── repositories/                      
│   ├── employeeRepository.go
│   ├── employeeRepositoryImpl.go
│   ├── patientRepository.go
│   └── patientRepositoryImpl.go
├── routes/
│   ├── auth.go
│   ├── doctor.go
│   └── receptionist.go
├── tests/
│   └── unit/
│       ├── auth_test.go
│       ├── patient_test.go
│       └── testutils/
│           └── testdb.go
├── development.env
├── production.env
├── go.mod
├── go.sum
└── README.md

```
---

## 🔐 Authentication & Authorization

- Uses **JWT** to generate tokens with embedded role claims.
- Tokens are stored in **HttpOnly cookies** and in **Redis** with expiry.
- Role-based access ensures:
  - Doctors access only doctor routes.
  - Receptionists access only receptionist routes.
- Middleware extracts and validates token & role from Redis.

---

## 📦 Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** Gin
- **Database:** PostgreSQL (GORM ORM)
- **Authentication:** JWT
- **Cache / Rate Limiting:** Redis
- **Testing:** Go's `testing` package
- **Design Patterns:** Repository Pattern

---
## Configure   .env:
```
PORT=3000
DB_URL=postgres://username:password@localhost:5432/clinicdb
JWT_SECRET=your_jwt_secret
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
```