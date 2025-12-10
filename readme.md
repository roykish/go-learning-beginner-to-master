# Go Learning: Beginner → Master

Concise, hands-on learning repo for Go developers. This guide contains short lessons, sample projects, and example code to move from beginner to confident Go programmer.

**Status:** Work in progress — Topic 1 (CLI Calculator) implemented.

**Quick links:** `go` >= 1.20 recommended

**Table of Contents**

- [Overview](#overview)
- [What You Will Learn](#what-you-will-learn)
- [Project: CLI Calculator](#project-cli-calculator)
  - [Features](#features)
  - [Usage](#usage)
  - [Examples](#examples)
- [Repository Layout](#repository-layout)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Install & Run](#install--run)
- [Contributing](#contributing)
- [Resources](#resources)
- [License](#license)

## Overview

This repository provides bite-sized topics and a small project to practice idiomatic Go: modules, package layout, functions, error handling, basic types, and pointers. It is intended for learners who already know basic programming concepts and want practical exercises.

## What You Will Learn

- How to initialize and use Go modules (`go mod init`, dependency management)
- How to structure packages and multi-file programs
- How the `main` package and entrypoint work
- Variables, constants, and primary data types
- Control flow: loops and conditionals
- Writing and testing functions with parameters and returns
- Basic error handling patterns
- Intro to pointers and simple usage

## Project: CLI Calculator

Small command-line calculator implemented as a learning project to demonstrate package structure, functions, and input validation.

### Features

- Accepts two numbers and an operator (`+ - * /`) via command-line arguments or interactive input
- Validates inputs and provides helpful error messages
- Organized across multiple files/packages to show real-project layout

### Usage

Run from the repository root (PowerShell example shown):

```powershell
cd "c:\Users\kisho\Documents\Personal\Learning & Implementation\Backend Technology\Go\go-learning-beginner-to-master"
go run ./go-basics-topic-1/calc
```

Alternatively, build a binary:

```powershell
go build -o bin/calc ./go-basics-topic-1/calc
.\bin\calc 12 + 5
```

### Examples

- Interactive: run with no args and follow prompts
- Direct: `./bin/calc 10 / 2` → prints `5`

## Repository Layout

`/go-basics-topic-1/` — Topic 1 code (CLI Calculator and supporting packages)

Other files:

- `readme.md` — repository overview (this file)

## Getting Started

### Prerequisites

- Go (recommended >= 1.20). Download: https://go.dev/dl/
- Basic command-line experience

### Install & Run

1. Clone the repo (if not already):

```powershell
git clone <repo-url>
cd "go-learning-beginner-to-master"
```

2. Inspect the topic folder and run the calculator example:

```powershell
go run ./go-basics-topic-1/calc
```

3. To build a standalone binary:

```powershell
go build -o bin/calc ./go-basics-topic-1/calc
```

## Contributing

Contributions welcome. Suggested workflow:

1. Fork the repo
2. Create a topic branch: `git checkout -b topic/<brief-desc>`
3. Implement changes and add tests where appropriate
4. Open a pull request with a concise description and small scope

Please follow idiomatic Go style and keep modules tidy.

## Resources

- Official Go documentation: https://go.dev/doc
- A Tour of Go: https://tour.golang.org
- Effective Go: https://go.dev/doc/effective_go

## License

This repository is provided as-is for learning purposes. Add a license file if you plan to share or reuse the content publicly.
