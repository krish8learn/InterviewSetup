# InterviewSetup — HighLevel Live Coding Interview

## Repo layout

This is a monorepo. The two sub-projects are **completely independent** — no shared
go.mod, no shared package.json, no shared config of any kind.

| Directory  | Stack        | Port | Status        |
|------------|--------------|------|---------------|
| /backend   | Go (stdlib)  | 8080 | active        |
| /frontend  | Next.js      | 3000 | added later   |

## Backend

- Language: Go (chosen over Node.js — HighLevel's real backend is Go)
- Router: stdlib `net/http` with Go 1.22+ method patterns; no third-party router
- DB: MongoDB Atlas, connected via `go.mongodb.org/mongo-driver/mongo`
- Config: `github.com/joho/godotenv` — secrets from `.env`, never hardcoded
- Module: `github.com/krish8learn/InterviewSetup/backend`

## Interview format

- Design data models
- Write REST APIs
- Implement business logic

Add interview routes in `backend/main.go` under the clearly marked TODO comment.
