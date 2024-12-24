# Static Pages Generator

A simple command-line application written in Go to generate and serve static HTML pages from Markdown files.

## Features
- **Build Command**: Converts Markdown files into static HTML pages.
- **Serve Command**: Hosts the generated static pages on a local HTTP server.
- Supports custom input/output directories for the build process.
- Allows specifying a custom port for the HTTP server.

## Prerequisites
- Go 1.18 or later installed on your system.

## Installation
1. Clone this repository:
   ```bash
   git clone https://github.com/andrek13/static-pages-generator.git
   cd static-pages-generator
   ```

2. Build the executable:
   ```bash
   go build -o static-pages-generator
   ```

3. Verify the executable:
   ```bash
   ./static-pages-generator --help
   ```

## Usage
### Build Command
The `build` command converts Markdown files into HTML pages.

#### Syntax:
```bash
./static-pages-generator build --input=<input-directory> --output=<output-directory>
```

#### Options:
- `--input`: Directory containing Markdown files. Default: `content`.
- `--output`: Directory where HTML files will be generated. Default: `public`.

#### Example:
```bash
./static-pages-generator build --input=content --output=public
```

### Serve Command
The `serve` command starts a local HTTP server to serve the generated HTML files.

#### Syntax:
```bash
./static-pages-generator serve --port=<port-number>
```

#### Options:
- `--port`: Port for the HTTP server. Default: `8080`.

#### Example:
```bash
./static-pages-generator serve --port=9090
```

### Full Workflow Example
1. Build HTML pages:
   ```bash
   ./static-pages-generator build --input=my_markdown --output=my_html
   ```

2. Serve the generated pages:
   ```bash
   ./static-pages-generator serve --port=8081
   ```

3. Open your browser and navigate to `http://localhost:8081` to view the pages.

## Project Structure
```
static-pages-generator/
├── content/          # Default directory for Markdown files
├── public/           # Default output directory for generated HTML files
├── main.go           # Source code
└── README.md         # Documentation
```

## Dependencies
- [blackfriday](https://github.com/russross/blackfriday): Markdown processor for Go.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

