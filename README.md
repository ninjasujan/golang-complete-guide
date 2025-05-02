# 📘 GoLang Backend Learning Syllabus (2 Weeks)

A focused 2-week hands-on plan to learn GoLang core syntax, backend programming principles, and build production-grade applications. This plan is designed for developers (like me) transitioning from JavaScript/TypeScript to GoLang for backend systems.

---

## 📅 Week 1 – GoLang Core Syntax & Internals

### ✅ **Day 1: Getting Started with GoLang**

- [ ] Install GoLang and setup `GOPATH`, `GOROOT`
- [ ] Write your first `Hello World` app
- [ ] Understand how Go files run (`main` package & func)
- [ ] Explore:
  - [ ] Variables (`var`, `:=`, `const`)
  - [ ] Data types (`int`, `float`, `bool`, `string`)
  - [ ] Type inference & zero values
- [ ] Control structures:
  - [ ] `if`, `else`, `switch`
  - [ ] `for` loops (only loop in Go)

---

### ✅ **Day 2: Types, Collections, Pointers**

- [ ] Understand:
  - [ ] Value types vs Reference types
  - [ ] Pointers (`*`, `&`, `new()`)
- [ ] Arrays and Slices:
  - [ ] Slice vs Array
  - [ ] Slice operations: append, copy, slicing
- [ ] Maps:
  - [ ] Create, access, delete keys
  - [ ] Check for key existence
- [ ] Structs:
  - [ ] Define and use custom types
  - [ ] Nested structs
  - [ ] Struct field tags (intro)
- [ ] Error handling:
  - [ ] Idiomatic Go pattern: `if err != nil`
  - [ ] Create custom errors using `errors.New()` or `fmt.Errorf()`

---

### ✅ **Day 3: Functions, Interfaces, Packages**

- [ ] Functions:
  - [ ] Multiple return values
  - [ ] Named return values
  - [ ] Variadic functions
- [ ] Methods vs Functions:
  - [ ] Define receiver functions
  - [ ] Pointer vs Value receivers
- [ ] Interfaces:
  - [ ] Define and implement interfaces
  - [ ] Duck typing in Go
  - [ ] Interface composition
- [ ] Composition over inheritance (no classes)
- [ ] Packages & Modules:
  - [ ] `go mod init`, project structure
  - [ ] Internal vs exported functions (`camelCase` vs `Capitalized`)
  - [ ] Create and import custom packages

---

### 🔁 **Day 4: Review + Practice**

- [ ] Review code from Day 1–3
- [ ] Practice problems:
  - [ ] Structs, slices, loops, interfaces
- [ ] Read & explain:
  - [ ] A standard library file (`strings`, `math`, etc.)
- [ ] Build a mini CLI calculator app or string utility
- [ ] Run: `go run`, `go build`, `go install`

---

### ✅ **Day 5: Concurrency & Goroutines**

- [ ] What are Goroutines?
  - [ ] Launching using `go func()`
- [ ] Channels:
  - [ ] Send/Receive values
  - [ ] Buffered vs Unbuffered channels
  - [ ] Closing channels
- [ ] `select` statement:
  - [ ] Multiple channel waiting
  - [ ] Timeout patterns
- [ ] Compare Go concurrency model vs JS async model

---

### ✅ **Day 6: Error Flow, Tooling, Testing**

- [ ] `defer` usage:
  - [ ] Cleaning up (closing files, etc.)
- [ ] `panic` and `recover`
  - [ ] Controlled exception-style handling
- [ ] Testing in Go:
  - [ ] Write unit tests using `testing.T`
  - [ ] Test tables pattern
- [ ] Tooling:
  - [ ] `go fmt`, `go vet`, `go mod tidy`
  - [ ] Documentation: `go doc`, `godoc`

---

### ✅ **Day 7: Apply & Build**

- [ ] Build a CLI-based app (choose one):
  - [ ] Todo app
  - [ ] File organizer
  - [ ] URL shortener (non-API version)
- [ ] Use:
  - [ ] Structs
  - [ ] Interfaces
  - [ ] File I/O
  - [ ] Goroutines (where possible)

---

## 📅 Week 2 – Real-World Backend Concepts

### ✅ **Day 8: File I/O and JSON**

- [ ] File operations:
  - [ ] Open, read, write, delete files (`os`, `ioutil`, `bufio`)
- [ ] JSON:
  - [ ] Encode/Decode structs
  - [ ] Use struct tags (`json:"field_name"`)
- [ ] Build a JSON file-based DB (mock local database)

---

### ✅ **Day 9: HTTP Servers & Middleware**

- [ ] Build an HTTP server using `net/http`
- [ ] Define routes with `http.HandleFunc`
- [ ] Create and use middleware:
  - [ ] Logging
  - [ ] Basic Authentication
- [ ] Use `Context` for request lifecycle
- [ ] Serve static files (e.g., images or JS files)

---

### ✅ **Day 10: Databases (PostgreSQL/MySQL)**

- [ ] Connect to a database using `database/sql`
- [ ] Run basic queries:
  - [ ] Create, Read, Update, Delete (CRUD)
- [ ] Use `sql.DB`, prepared statements
- [ ] Error checking and SQL injection safety

---

### ✅ **Day 11: MongoDB / ORM**

- [ ] Use MongoDB Go Driver (optional)
- [ ] Understand Go ORMs:
  - [ ] Use GORM or SQLX
  - [ ] Auto migrations and model definitions
  - [ ] Handling relationships

---

### ✅ **Day 12: Config, Env, Logging**

- [ ] Use `os.Getenv`, `.env` file
  - [ ] Libraries like `godotenv`
- [ ] Implement custom configuration loader
- [ ] Logger setup:
  - [ ] Standard log, logrus, or zap
  - [ ] Structured logs
- [ ] Dependency injection: Understand pattern in Go

---

### ✅ **Day 13: JWT Auth & Validators**

- [ ] JWT Authentication:
  - [ ] Generate and validate tokens
- [ ] Use middleware to protect routes
- [ ] Struct Tags and Validation:
  - [ ] Use `validator` library for payloads

---

### ✅ **Day 14: Final Project**

- [ ] Build a complete mini REST API:
  - [ ] Suggested: Blog, Todo, Notes API
- [ ] Features:
  - [ ] Routing and Controllers
  - [ ] Authentication
  - [ ] CRUD operations with DB
  - [ ] Error handling and logging
  - [ ] Environment config
- [ ] Deploy locally or test with Postman

---

## ✅ Completion Milestones

- [ ] Go syntax and types mastered
- [ ] Interfaces, structs, and methods applied
- [ ] Goroutines and channels understood and used
- [ ] File I/O and JSON encoding practiced
- [ ] Built a functional HTTP REST API
- [ ] Worked with a real database
- [ ] Authenticated endpoints with JWT
- [ ] Fully prepared to adopt Gin/Fiber in Week 3+

---

## 👨‍💻 Author

**Sujan** – Full Stack Developer | Currently building real-world projects using GoLang backend 🚀
