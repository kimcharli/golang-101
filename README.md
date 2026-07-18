# golang-101

A consolidated repository for early Go experiments, tutorials, and practice exercises.

## Workspace Structure

This repository is configured as a Go Workspace (`go.work`), allowing all modules to be developed, compiled, and linted seamlessly.

### Modules & Projects

1. **`001-main/`**
   - A standalone Go program printing `"Hello Go!"`.
2. **`002-basic-main/`**
   - A standalone web server experimenting with hot-reload and graceful shutdown via signal handling (`os.Signal`).
3. **`003-config/`**
   - A Go module demonstrating how to load and parse YAML configurations using `gopkg.in/yaml.v3`.
4. **`004-htmx/`**
   - A full-featured web application using [Echo v4](https://github.com/labstack/echo), HTMX templates, and YAML configuration.
5. **`mod101/hello/`**
   - The classic "Get started with Go" hello-world tutorial with custom utilities and unit tests.
6. **`templates/template001/`**
   - Basic Go standard template engine demo (`temp001.go`).
   - `REST/`: A lightweight REST API server showcasing JSON marshaling/unmarshaling, synchronizing store with mutexes, and a basic Dockerfile.
