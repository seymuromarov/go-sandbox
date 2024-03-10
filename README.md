# Sandbox Code Executor

The project is a Docker-based environment for executing code snippets in multiple programming languages, including Python, JavaScript, TypeScript, Java, C++, and Go. This tool is particularly useful for educational purposes, coding assessments, and automated testing environments, providing a simple and isolated execution space.

## Features

- Support for multiple programming languages.
- Docker-based execution environment for code isolation.
- Environment variable configuration for easy customization.
- Execution time and memory limit controls.

## Prerequisites

Before you begin, ensure you have the following installed on your system:
- Docker
- Go (for running the Go-based server and utility scripts)

## Installation

1. **Clone the repository:**
```bash
git clone https://github.com/seymuromarov/go-sandbox.git
cd go-sandbox
go get
```

2. **Load Environment Variables:**
Place a `.env` file in the root directory with the necessary configuration parameters. Refer to .env.example for guidance on the required settings.

## Running the Project

To run a code snippet, use the following command format:
```bash
go run . <language extension> <file path>
```

For example, to run a Go file located at `./examples/fibonacci.go`, you would use:
```bash
go run . .go ./examples/fibonacci.go
```

Please note that the first run will take additional time due to the Docker image being built.

## Adding Support for a New Language

To add support for a new programming language, follow these steps:

1. **Update Dockerfile:**

Ensure the Docker image has the necessary compilers or interpreters for the new language.

2. **Modify `dockerutils/config.go`:**

Add a new case in the `GetDockerCommand` function to handle the new language extension and define the command to compile/run the code.
