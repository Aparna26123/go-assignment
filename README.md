# Tech Stack

- **Language**: Go
- **Framework**: Gorilla Mux
- **Database**: SQLite
- **ORM**: GORM
- **Authentication**: JWT
- **Testing**: Go’s built-in testing
- **Deployment**: Local (Docker optional)



# Project Structure
task-manager/
├── cmd/                  # Main application entry point
├── internal/
│   ├── handlers/         # HTTP handlers
│   ├── models/           # Data models
│   ├── services/         # Business logic
│   ├── repository/       # DB access layer
│   └── middleware/       # Auth, logging, etc.
├── pkg/                  # Utility packages
├── config/               # Configuration files
├── migrations/           # DB migration scripts
├── go.mod
└── README.md

