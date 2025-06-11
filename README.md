# 📰 News Platform (Golang)

A server-side news publishing platform built with Go, Gin, and GORM. Supports API access to articles, categories, tags, and authors, along with seed data and test cases.

---

## 🚀 Features

- ✅ Fetch list of **categories**
- ✅ Fetch list of **articles** with filters:
  - category ID
  - author ID
  - tag name
- ✅ Get **article details** by ID
- ✅ Get **author details** by ID
- ✅ Seed database with sample data
- ✅ Unit tests using `net/http/httptest` and `testify`

---

## 🛠️ Tech Stack

- **Go (Golang)**
- **Gin** – Web framework
- **GORM** – ORM for SQLite
- **SQLite** – Lightweight DB
- **Testify** – Test assertions

---

## 📦 Installation

### 1. Clone the Repo

```bash
git clone https://github.com/your-username/news-platform.git
cd news-platform
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run main.go
```

Runs on: `http://localhost:8080`

---

## 📚 API Endpoints

| Method | Endpoint        | Description                |
| ------ | --------------- | -------------------------- |
| GET    | `/categories`   | Get list of all categories |
| GET    | `/articles`     | Get list of all articles   |
| GET    | `/articles/:id` | Get single article by ID   |
| GET    | `/authors/:id`  | Get author details by ID   |

#### Example filters for `/articles`:

```bash
/articles?categoryId=1&authorId=2&tag=GoLang
```

---

## 🧪 Running Tests

### Install test dependencies

```bash
go get github.com/stretchr/testify/assert@v1.10.0
```

### Run tests

```bash
go test ./tests
```

---

## 🧬 Project Structure

```
.
├── config         # DB setup
├── controllers    # Route handlers
├── models         # DB models
├── routes         # Route definitions
├── utils          # Seed & migration
├── tests          # Test cases
└── main.go        # App entry point
```

---

## 🧠 Seed Data

On first run, the app automatically seeds:

- 3 Categories: News, Tech, Lifestyle
- 3 Tags: AI, GoLang, Video
- 2 Authors
- 3 Articles

No need to insert anything manually!

---

## 🗂 Sample Article Schema

```json
{
  "title": "Go Tutorial",
  "sub_title": "Learning Go",
  "image_url": "https://...",
  "article_type": "text",
  "description": "<p>GoLang is great</p>",
  "media_url": null,
  "publish_date": "2025-06-11",
  "author": { ... },
  "category": { ... },
  "tags": [ ... ]
}
```

---

## 📌 To Do (Optional Enhancements)

- [ ] CMS UI (React or Admin Panel)
- [ ] Pagination support
- [ ] Authentication for admin routes (JWT or session)

---

## 👨‍💻 Author

Made by [Jaison Rego](https://github.com/JaisonRego)  
Feel free to contribute or raise an issue 🙌

---
