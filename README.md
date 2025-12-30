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

Jump from an interface usage to its definition in the core library.

- In `main.go`, click on `Notifier` (line 23) → **Go to Definition**
- → Highlights `Notifier` (line 7) interface in `notifier-core/notifier.go`

**Benefit:** Instantly understand the contract your service depends on—see the interface definition without switching repos or searching documentation.

### 2. Go to Definition (to notifier-email)

Navigate from a constructor call to its implementation in another repository.

- In `main.go`, click on `NewEmailNotifier` (line 13) → **Go to Definition**
- → Highlights `NewEmailNotifier` (line 18) function in `notifier-email/email.go`

**Benefit:** Dive into the implementation details of your dependencies—debug issues or understand behavior without cloning and searching multiple repos.

### 3. Find Implementations (cross-repo)

Discover all concrete implementations of an interface across repositories.

- In `main.go`, click on `Notifier` (line 23) → **Find Implementations**
- → Highlights `EmailNotifier` (line 11) struct in `notifier-email/email.go`

**Benefit:** See all available implementations when working with dependency injection—critical for understanding which adapters exist and could be swapped in.
