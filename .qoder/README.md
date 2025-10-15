# Qoder Project Configuration

This directory contains Qoder IDE configuration for the aiq project.

## Project Overview

**aiq** is a command-line tool for AI queries using Google Gemini API.

## Quick Commands

### Build
```bash
go build -o aiq
```

### Run
```bash
./aiq "your query here"
```

### Install
```bash
go install
```

## Development Notes

### Key Files
- `main.go` - Main application entry point
- `go.mod` - Go module dependencies
- `.env.example` - Environment configuration template

### Environment Setup
1. Copy `.env.example` to `.env`
2. Add your `GEMINI_API_KEY`
3. Run the tool

### Flags
- `-m` - Model version
- `-p` - Provider (gemini)
- `-s` - System instruction
- `-shell` - Shell command output mode
- `-json` - JSON output mode

### Testing Ideas
```bash
# Basic query
./aiq "test query"

# With stdin
echo "data" | ./aiq "analyze this"

# Shell mode
./aiq -shell "list files"

# JSON mode
./aiq -json "structured data request"
```

## Architecture

The application follows a simple architecture:
1. Parse command-line flags
2. Load environment from `.env` if available
3. Read stdin data if provided
4. Configure Gemini client with API key
5. Apply system instructions based on flags
6. Generate and output response
