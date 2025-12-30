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

## Precise Code Navigation Demo

This repo is the **consumer** that ties everything together. Try these:

1. **"Go to Definition"** on `core.Notifier` → jumps to `notifier-core/notifier.go`
2. **"Go to Definition"** on `core.NewMessage` → jumps to `notifier-core/message.go`
3. **"Go to Definition"** on `email.NewEmailNotifier` → jumps to `notifier-email/email.go`
4. **"Find References"** on `Message` in `notifier-core` → shows usage in all 3 repos
