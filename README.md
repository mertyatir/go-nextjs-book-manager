# go-nextjs-book-manager

## Table of Contents
- [Introduction](#introduction)
- [Setup Instructions](#setup-instructions)
- [Project Structure](#project-structure)
- [Endpoint Usage](#endpoint-usage)
- [Running Tests](#running-tests)


## Introduction
The Book Manager App is a full-fledged CRUD application designed to manage a small library of books. The frontend is built using React/Next.js and TypeScript, while the backend is developed using Golang with SQLite as the database. Key features include:

- Add, update, delete, and view books.
- Simple and intuitive user interface.
- Robust backend with efficient database management.

## Setup Instructions

### Prerequisites
- Go (version 1.22.5 or higher)
- Node.js (version 20.15.0 or higher)
- npm (version 10.7.0 or higher)
- Git

### Cloning the Repository

```sh
git clone https://github.com/mertyatir/go-nextjs-book-manager.git
cd go-nextjs-book-manager
```

### Backend Setup
#### Initializing Go Modules

```
cd backend
go mod tidy
```

#### Building the Application

```
go build
```

#### Running the Application
```
./book-manager
```

### Frontend Setup
#### Installing Dependencies
```
cd frontend
npm install
```
#### Running the Frontend
```
npm run dev
```

### Project Structure
A brief overview of the project structure:
```
go-nextjs-book-manager/
├── backend/              # Backend source code
│   ├── database/         # Database related code
│   │   └── database.go   # Database initialization and migration
│   ├── docs/             # Swagger related code
│   ├── handlers/         # HTTP handlers
│   │   ├── handlers.go   # Handlers for RESTful API
│   │   └── urlHandler.go # Handlers for URL Cleanup and Redirection Service
│   ├── models/           # Data models
│   │   └── book.go       # Book model
│   ├── tests/            # Unit tests
│   ├── go.mod            # Go module file
│   ├── go.sum            # Go checksum file
│   ├── main.go           # Entry point of the backend application
│   └── ...               # Other backend files
├── frontend/             # Frontend source code
│   ├── src/              # Next.js app source code
│   │   ├── app/          # Next.js app router
│   │   │   └── books/    # Book-related pages
│   │   │       └── [id]/ # Dynamic routes for individual book details
│   │   ├── components/   # React components
│   │   ├── context/      # Context providers
│   │   ├── services/     # Service functions for data fetching
│   │   └── types/        # TypeScript types
│   └── ...               # Other frontend files
├── screenshots/          # Screenshots of the application
└── README.md             # Project README file
```

### Endpoint Usage
You can access the list of available endpoints and their usage at:
 [Swagger Documentation](http://localhost:8000/swagger/index.html)

### Running Tests
#### Unit Tests
```
cd tests/
go test
```


