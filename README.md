# Golang Web Application

A modern web application built with Go, using Chi framework, HTMX, and Tailwind CSS.

## Features

- **Chi Framework**: Lightweight and flexible router for Go
- **HTMX**: Dynamic HTML updates without writing JavaScript
- **Tailwind CSS**: Utility-first CSS framework
- **SQLite Database**: Simple and reliable data storage
- **Live Reload**: Development with instant feedback
- **WebSocket Support**: Real-time communication capabilities

## Prerequisites

- Go 1.21 or higher
- Node.js and npm (for Tailwind CSS)
- macOS, Linux, or WSL for Windows users

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/despiegk/golang_web.git
   cd golang_web
   ```

2. Make the installation script executable and run it:
   ```bash
   chmod +x install.sh
   ./install.sh
   ```

   This script will:
   - Install required Go tools (go-blueprint, air, templ)
   - Install Tailwind CSS (requires Homebrew on macOS)
   - Download Go dependencies
   - Generate template files
   - Build CSS assets

## Running the Application

The application can be run in two modes:

1. **Standard Mode**:
   ```bash
   chmod +x start.sh
   ./start.sh
   ```

2. **Development Mode** (with live reload):
   ```bash
   ./start.sh --dev
   ```

The application will be available at `http://localhost:8080`

## Project Structure

```
heroweb/
├── cmd/
│   ├── api/        # Main application entry point
│   └── web/        # Web-related code (templates, assets)
├── internal/
│   ├── database/   # Database operations
│   ├── dns/        # DNS-related functionality
│   ├── filemanager/# File management operations
│   └── server/     # Server configuration and routes
├── .air.toml       # Live reload configuration
├── go.mod          # Go module file
├── Makefile        # Build and development commands
└── tailwind.config.js # Tailwind CSS configuration
```

## Development Commands

The following commands are available through the Makefile:

- `make build`: Build the application
- `make run`: Run the application
- `make test`: Run tests
- `make clean`: Clean build artifacts
- `make watch`: Run with live reload (requires air)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
