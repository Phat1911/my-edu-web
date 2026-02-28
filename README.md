# EduWeb

A fullstack web application for educational content, built with Next.js (frontend) and Go/Gin (backend), backed by PostgreSQL on Neon.

## Tech Stack

| Layer | Technology |
|---|---|
| Frontend | Next.js 16, React 19, TypeScript, Tailwind CSS 4 |
| Backend | Go, Gin, JWT authentication |
| Database | PostgreSQL (Neon cloud) |
| Auth | JWT via httpOnly cookies |

## Project Structure

```
edu-web/
├── frontend/          # Next.js app
│   ├── src/
│   ├── package.json
│   └── tsconfig.json
└── backend/           # Go/Gin API server
    ├── cmd/main.go    # Entry point
    ├── config/        # App config (loads .env)
    ├── internal/
    │   ├── config/    # Shared config (JWT secret)
    │   ├── handlers/  # HTTP handlers (auth, messages)
    │   ├── middleware/ # JWT auth middleware
    │   ├── models/    # Data models
    │   └── repository/# Database layer (pgxpool)
    ├── bin/server     # Compiled binary
    └── go.mod
```

## Getting Started

### Prerequisites

- Go 1.24+
- Node.js 22+
- PostgreSQL database (or Neon account)

### Backend

1. Copy the example env file and fill in your values:

```bash
cp backend/.env.example backend/.env
```

2. Edit `backend/.env`:

```env
DATABASE_URL=postgresql://user:password@host/dbname?sslmode=require
PORT=8080
FRONTEND_URL=http://localhost:3000
JWT_SECRET=your-strong-secret-here
```

3. Build and run:

```bash
cd backend
go build -o bin/server ./cmd/main.go
./bin/server
```

The server starts on `http://localhost:8080`. On first run it automatically runs all database migrations and seeds initial data.

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The frontend starts on `http://localhost:3000`.

## API Endpoints

### Auth

| Method | Endpoint | Description | Auth required |
|---|---|---|---|
| POST | `/api/v1/auth/register` | Register new account | No |
| POST | `/api/v1/auth/login` | Login | No |
| POST | `/api/v1/auth/logout` | Logout (clears cookie) | No |
| GET | `/api/v1/auth/me` | Get current user | Yes |

### Content

| Method | Endpoint | Description | Auth required |
|---|---|---|---|
| GET | `/api/v1/health` | Health check | No |
| GET | `/api/v1/videos` | List videos | No |
| GET | `/api/v1/audios` | List audios | No |
| GET | `/api/v1/qrcodes` | List QR codes | No |
| POST | `/api/v1/qrcodes/generate` | Generate QR code | No |
| POST | `/api/v1/chat` | Send chat message | No |
| GET | `/api/v1/chat/:session_id` | Get chat history | No |

### Messaging (protected)

| Method | Endpoint | Description | Auth required |
|---|---|---|---|
| POST | `/api/v1/messages` | Send direct message | Yes |
| GET | `/api/v1/messages/:other_user_id` | Get messages with user | Yes |
| GET | `/api/v1/users` | List all users | Yes |

## Authentication

Auth uses JWT tokens stored as `httpOnly` cookies (`eduhub_token`). The token expires after 24 hours.

API clients (non-browser) can also pass the token via `Authorization: Bearer <token>` header.

## Environment Variables

| Variable | Required | Default | Description |
|---|---|---|---|
| `DATABASE_URL` | Yes | - | PostgreSQL connection string |
| `PORT` | No | `8080` | Server port |
| `FRONTEND_URL` | No | `http://localhost:3000` | Allowed CORS origin |
| `JWT_SECRET` | No | `eduweb-secret-key-2026` | JWT signing secret (set this in production!) |
| `ENV` | No | - | Set to `production` to enable Secure cookie flag |

## Development

### Rebuild backend

```bash
cd backend
go build -o bin/server ./cmd/main.go && ./bin/server
```

### Run without building (dev mode)

```bash
cd backend
go run ./cmd/main.go
```

### Verify build only

```bash
cd backend
go build ./...
```
