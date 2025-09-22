# README.md
# Practice Hexagonal - Q&A System with OAuth 2.0

A Q&A system built with hexagonal architecture, featuring Google OAuth 2.0 authentication and role-based access control.

## 🏗️ Architecture

This project follows **Hexagonal Architecture** (Ports & Adapters) with clean separation of concerns:

```
internal/
├── core/                    # Business Logic (Pure)
│   ├── domain/             # Entities
│   │   ├── user.go
│   │   └── question.go
│   ├── port/               # Interfaces
│   │   ├── oauth.go
│   │   ├── auth.go
│   │   └── qa.go
│   └── service/            # Business Logic
│       ├── auth.go
│       └── qa.go
└── adapter/                # External Dependencies
    ├── googleoauth/        # OAuth Client
    │   └── client.go
    ├── sqlite/             # Database
    │   ├── db.go
    │   ├── user_repository.go
    │   └── qa_repository.go
    └── fiber/              # HTTP Server
        ├── middleware/
        ├── models/
        ├── controllers/
        └── routes/
```

## 🚀 Features

- **OAuth 2.0 + SSO**: Google authentication with cookie-based state validation
- **JWT Authentication**: Secure token-based API access
- **Role-based Access**: User and Admin roles
- **Q&A System**: Users ask questions, admins provide answers
- **Hexagonal Architecture**: Clean, testable, maintainable code
- **SQLite Database**: Lightweight with auto-migration

## 📋 Setup Instructions

### 1. Google OAuth Setup
1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create project → Enable Google+ API
3. Create OAuth 2.0 Client ID
4. Set redirect URI: `http://localhost:3000/auth/google/callback`

### 2. Environment Configuration
Create `.env` file:
```env
GOOGLE_CLIENT_ID=YOUR_GOOGLE_CLIENT_ID
GOOGLE_CLIENT_SECRET=YOUR_GOOGLE_CLIENT_SECRET
BASE_URL=http://localhost:3000
JWT_SECRET=9b8c7d6e5f4a3b2c1d0e9f8a7b6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e1f0a9b8c7d6e5f4a
SESSION_SECRET=a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6
```

### 3. Run Application
```bash
go mod tidy
go run main.go
```

## 🔗 API Endpoints

### Authentication (Public)
- `GET /` - Home page with login link
- `GET /auth/google/login` - Redirect to Google OAuth
- `GET /auth/google/callback` - OAuth callback handler

### Authentication (Protected)
- `GET /auth/profile` - Get user profile
- `POST /auth/logout` - Logout user

### Q&A System (Protected)
- `POST /qa/` - Ask question (User)
- `GET /qa/` - Get all questions (User)
- `PUT /qa/:id/answer` - Answer question (Admin only)

### System
- `GET /health` - Health check

## 🧪 Testing

### Browser Testing (Recommended)
1. Open `http://localhost:3000`
2. Click login link → redirected to Google
3. Complete OAuth flow → get JWT token in response

### HTML Test Interface
1. Open `test.html` in browser
2. Click "Login with Google" (opens popup)
3. Complete login → popup closes, token saved
4. Test Q&A functionality

### cURL Testing
```bash
# Get login URL (will redirect)
curl -v http://localhost:3000/auth/google/login

# After getting JWT token from callback:
export TOKEN="your-jwt-token-here"

# Ask question
curl -X POST http://localhost:3000/qa/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content": "What is hexagonal architecture?"}'

# Get questions
curl -H "Authorization: Bearer $TOKEN" \
  http://localhost:3000/qa/

# Answer question (admin only)
curl -X PUT http://localhost:3000/qa/1/answer \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"answer": "Hexagonal architecture separates business logic from external concerns."}'
```

## 👥 User Roles

- **User**: Ask questions, view Q&As
- **Admin**: Everything users can do + answer questions

To make someone admin, update database directly:
```sql
UPDATE users SET role = 'admin' WHERE email = 'admin@example.com';
```

## 🔒 Security Features

- **State Parameter Validation**: Prevents CSRF attacks
- **Cookie-based State Storage**: Secure state management
- **JWT Token Authentication**: Stateless API security
- **Role-based Authorization**: Admin-only endpoints
- **Input Validation**: Request data validation
- **CORS Protection**: Cross-origin security

## 📊 Database Schema

### Users
- `id`, `google_id`, `email`, `name`, `picture`, `role`, `created_at`, `updated_at`

### Questions
- `id`, `content`, `answer`, `user_id`, `created_at`, `updated_at`

## 🎯 Key Differences from Previous Version

1. **Cookie-based State Validation**: More secure OAuth flow
2. **Redirect-based Login**: Better user experience
3. **Simplified Error Handling**: Cleaner responses
4. **Working Example Integration**: Proven OAuth flow
5. **Enhanced Logging**: Better debugging information

This implementation combines the robustness of hexagonal architecture with the reliability of the working OAuth example you provided!