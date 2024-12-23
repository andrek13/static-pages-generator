# Static Pages Generator

This project is a static pages generator written in Go. It converts Markdown files into HTML and serves them over HTTP. The project includes functionality to build the static HTML files and to serve them using a simple HTTP server.

## Project Structure

- **main.go**: The main entry point of the application.
- **content/**: Directory containing the Markdown files to be converted.
- **public/**: Directory where the generated HTML files will be stored.
- **public/dashboard-template.html**: HTML template used for generating the final HTML pages.

## Prerequisites

- [Go 1.16 or later](https://golang.org/dl/)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/andrek13/static-pages-generator.git
   cd static-pages-generator
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

### Build

To convert the Markdown files in the `content/` directory to HTML and save them in the `public/` directory, run:

```bash
go run main.go build
```

### Serve

To serve the generated HTML files over HTTP, run:

```bash
go run main.go serve
```

The server will start on [http://localhost:8080](http://localhost:8080).
