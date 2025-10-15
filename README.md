# aiq

Simple command line tool for AI query, similar to `gemini -p` command.

[![Go Reference](https://pkg.go.dev/badge/github.com/chinglinwen/aiq.svg)](https://pkg.go.dev/github.com/chinglinwen/aiq)
[![Go Report Card](https://goreportcard.com/badge/github.com/chinglinwen/aiq)](https://goreportcard.com/report/github.com/chinglinwen/aiq)

## Quick Start

```bash
# Install
go install github.com/chinglinwen/aiq@latest

# Configure (choose one method)
echo "GEMINI_API_KEY=your-api-key" > .env
# or
export GEMINI_API_KEY="your-api-key"

# Use
aiq "What is the capital of France?"
echo "Hello, World!" | aiq "Translate to Spanish"
aiq -shell "list all go files" | sh
aiq -json "List 3 programming languages"
```

## Features

- Support for Google Gemini API
- Read data from stdin
- Query prompt from command arguments
- Flags:
  - `-m` for model version (default: "gemini-2.5-flash")
  - `-p` for provider (default: "gemini")
  - `-s` for system instruction to guide AI behavior
  - `-shell` for shell command output mode (outputs only commands, no explanations)
  - `-json` for JSON output format (outputs valid JSON only)
- Reads API key from environment variables

## Installation

### Option 1: Install with go install (Recommended)

```bash
go install github.com/chinglinwen/aiq@latest
```

This will install the `aiq` binary to your `$GOPATH/bin` directory (usually `~/go/bin`).

Make sure `$GOPATH/bin` is in your PATH:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Option 2: Build from source

```bash
git clone https://github.com/chinglinwen/aiq.git
cd aiq
go build -o aiq
```

Optionally, move the binary to a location in your PATH:
```bash
sudo mv aiq /usr/local/bin/
```

## Configuration

### Option 1: Using .env file (Recommended)

Create a `.env` file in the same directory as the `aiq` binary or in your working directory:

```bash
cp .env.example .env
# Edit .env and add your API key
```

Example `.env` file:
```
GEMINI_API_KEY=your-api-key-here
```

The `.env` file will be automatically loaded when you run `aiq`.

### Option 2: Environment Variables

Set your Gemini API key as an environment variable:

```bash
export GEMINI_API_KEY="your-api-key-here"
# or
export GOOGLE_API_KEY="your-api-key-here"
```

## Usage

Basic query:
```bash
aiq "What is the capital of France?"
```

With stdin data:
```bash
echo "Hello, World!" | aiq "Translate this to Spanish"
```

```bash
cat data.txt | aiq "Summarize this text"
```

With custom model:
```bash
aiq -m gemini-1.5-flash "Explain quantum computing"
```

With system instruction:
```bash
aiq -s "You are a helpful coding assistant. Be concise." "Explain Go interfaces"
```

Shell command mode (output only commands):
```bash
aiq -shell "list all go files recursively"
# Output: find . -name "*.go"

aiq -shell "find large files over 100MB"
# Can be piped directly: aiq -shell "find large files" | sh
```

JSON output mode:
```bash
aiq -json "List 3 programming languages with their use cases"
# Output: {"languages": [{"name": "Go", "use_case": "Systems programming"}, ...]}

cat data.txt | aiq -json "Analyze this data and return structured results"
# Output: {"analysis": {...}}
```

Combined:
```bash
cat logs.txt | aiq -m gemini-1.5-flash "Analyze these logs for errors"
```

## Available Models

- `gemini-2.5-flash` (default)
- `gemini-2.5-pro`
- `gemini-1.5-flash`
- `gemini-1.5-pro`
- `gemini-pro`

## Examples

```bash
# Simple question
aiq "What is Go programming language?"

# Code review
cat main.go | aiq "Review this code for potential improvements"

# Data analysis
cat sales.csv | aiq "Analyze this sales data and provide insights"

# Translation
echo "Hello" | aiq "Translate to French, Spanish, and German"

# Generate and execute shell commands
aiq -shell "compress all log files in current directory" | sh

# Custom AI persona
aiq -s "You are a senior Go developer" "Best practices for error handling"

# Get structured JSON data
aiq -json "List top 5 Go frameworks for web development with descriptions"

# Combine flags for JSON shell commands (advanced)
cat logs.txt | aiq -json -s "Extract error patterns" "Analyze these logs"
```
