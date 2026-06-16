# InterviewSetup — HighLevel Live Coding Interview

## Repo layout

This is a monorepo. The two sub-projects are **completely independent** — no shared
go.mod, no shared package.json, no shared config of any kind.

| Directory  | Stack        | Port | Status        |
|------------|--------------|------|---------------|
| /backend   | Go (stdlib)  | 8080 | active        |
| /frontend  | Next.js (JS, App Router) | 3000 | active |

## Backend

- Language: Go (chosen over Node.js — HighLevel's real backend is Go)
- Router: stdlib `net/http` with Go 1.22+ method patterns; no third-party router
- DB: MongoDB via local Docker (`mongodb://localhost:27017`), connected via `go.mongodb.org/mongo-driver/mongo`
- Config: `github.com/joho/godotenv` — secrets from `.env`, never hardcoded
- Module: `github.com/krish8learn/InterviewSetup/backend`

## Interview format

- Design data models
- Write REST APIs
- Implement business logic

Add interview routes in `backend/main.go` under the clearly marked TODO comment.

## Frontend

- Framework: Next.js (JavaScript, App Router) — independent project in `/frontend`
- Port: 3000
- Fully independent of `/backend` — own `package.json`, own `node_modules`, no shared config
- Home page (`/frontend/app/page.js`): client component that fetches `GET /` from the Go backend on mount
- Config: `NEXT_PUBLIC_API_URL` in `.env.local` (gitignored); see `.env.local.example`
