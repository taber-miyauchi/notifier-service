# notifier-service

## Project Structure

```
├── notifier-core/          ← The foundation
│   ├── go.mod
│   ├── notifier.go         ← Notifier interface
│   ├── message.go          ← Message struct, Priority enum
│   └── README.md
│
├── notifier-email/         ← Implementation
│   ├── go.mod              ← depends on notifier-core
│   ├── email.go            ← EmailNotifier implements Notifier
│   └── README.md
│
└── notifier-service/       ← Consumer/API (you are here)
    ├── go.mod              ← depends on both
    ├── main.go             ← HTTP server using both packages
    └── README.md
```

## Overview

HTTP API server that sends notifications using the notifier system.

## Endpoints

```
POST / 
  - to: recipient email
  - subject: email subject
  - body: email body
```

## Dependencies

- `github.com/sourcegraph-ce/notifier-core` - For `Notifier` interface and `Message` type
- `github.com/sourcegraph-ce/notifier-email` - For `EmailNotifier` implementation

## Testing Precise Code Navigation

Open this repo in Sourcegraph and try:

### 1. Go to Definition (to notifier-core)

- In `notifier-service/main.go`, click on `Notifier` (line 23) → **Go to Definition**
- → Jumps to `Notifier` (line 7) interface in `notifier-core/notifier.go`

### 2. Go to Definition (to notifier-email)

- In `notifier-service/main.go`, click on `email.NewEmailNotifier` (line 13) → **Go to Definition**
- → Jumps to `NewEmailNotifier` (line 18) function in `notifier-email/email.go`

### 3. Find Implementations (cross-repo)

- In `notifier-service/main.go`, click on `Notifier` (line 23) → **Find Implementations**
- → Shows `EmailNotifier` in `notifier-email/email.go`
