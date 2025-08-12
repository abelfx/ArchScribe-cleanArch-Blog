# G6 Blog Starter Project

## Overview

The **G6 Blog Starter Project** is a backend API for a blog platform that enables users to create, read, update, and delete blog posts, manage user profiles, and perform advanced search and filtering operations. The platform supports user authentication and authorization, different user roles (Admin and User), and AI integration for content suggestions or enhancements.

---

## Features

- **User Management**
  - User Registration with email verification
  - Login with JWT-based authentication (access & refresh tokens)
  - Password reset (Forgot Password)
  - User logout and token invalidation
  - Role-based access control (User & Admin)
  - User promotion/demotion by Admin

- **Blog Management**
  - Create, read, update, and delete blog posts
  - Tagging, filtering, and search by title or author name
  - Pagination and sorting of blog posts
  - Popularity tracking: views, likes, dislikes, comments
  - Duplicate interaction prevention (likes/dislikes)
  
- **AI Integration**
  - Generate blog content or suggestions based on user keywords/topics
  
- **Profile Management**
  - Update user profile details including bio, profile picture, and contact info
  
---

## Architecture & Technologies

- **Language:** Go (Golang)
- **Architecture:** Clean Architecture with separation of concerns (Domain, Usecases, Infrastructure, Delivery)
- **Database:** MongoDB
- **Authentication:** JWT Tokens (access and refresh tokens)
- **AI Integration:** Mistral AI for content suggestion/generation
- **Web Framework:** Gin Gonic for HTTP routing and middleware
- **Configuration:** `.env` files with `godotenv`
- **Concurrency:** Uses Go’s goroutines for scalable request handling

---

## API Endpoints

### User Management

| Endpoint             | Method | Description                             |
|----------------------|--------|-------------------------------------|
| `/api/register`      | POST   | Register a new user                   |
| `/api/login`         | POST   | Login and receive tokens              |
| `/api/logout`        | POST   | Logout and invalidate tokens          |
| `/api/forgot-password`| POST  | Request password reset link           |
| `/api/reset-password` | POST  | Reset password using token            |
| `/api/users/promote` | POST   | Promote user to Admin (Admin only)    |

### Blog Management

| Endpoint               | Method | Description                           |
|------------------------|--------|-------------------------------------|
| `/api/blogs`           | POST   | Create a new blog post                |
| `/api/blogs`           | GET    | Retrieve paginated blog posts         |
| `/api/blogs/:id`       | PUT    | Update blog post (author only)        |
| `/api/blogs/:id`       | DELETE | Delete blog post (author/Admin)       |
| `/api/blogs/search`    | GET    | Search blogs by title/author           |
| `/api/blogs/filter`    | GET    | Filter blogs by tags, date, popularity |
| `/api/blogs/:id/like`  | POST   | Like a blog post                       |
| `/api/blogs/:id/dislike`| POST  | Dislike a blog post                    |

### AI Integration

| Endpoint           | Method | Description                               |
|--------------------|--------|-----------------------------------------|
| `/api/blogs/suggest`| POST   | Generate blog content suggestions via AI|

---

## Getting Started

### Prerequisites

- Go 1.18+
- MongoDB instance
- Mistral AI API key (set in `.env` file)

### Setup

1. Clone the repository:
   ```
      git clone https://github.com/abelfx/ArchScribe-cleanArch-Blog
      cd blogcleanarc
   ```
2. Create a .env file
```
MONGODB_URI=mongodb://localhost:27017
MISTRAL_API_KEY=your_mistral_api_key_here
JWT_SECRET=your_jwt_secret_here
```
3. Run the project
```
go run main.go
```
4.Access API at http://localhost:3000

## Security
- Passwords are hashed using bcrypt before storing in the database.
- JWT tokens are signed and validated for authentication.
- Role-Based Access Control enforced via middleware.
- Tokens stored securely with expiration and refresh mechanisms.

## Performance & Scalability
- Implements pagination and sorting for large datasets.
- Utilizes Go’s goroutines for concurrent request handling.
- Potential caching layers can be added for expensive queries.

## Contribution
Contributions are welcome! Please open issues or pull requests for bug fixes, feature requests, or improvements.

**License**
MIT License © 2025 Abel Tesfa


