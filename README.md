# opsctl

A lightweight DevOps CLI tool built with Go and Cobra to simulate real-world infrastructure operations and Kubernetes diagnostics.

---

# Features

* Cluster health checks
* Failed pod inspection
* CLI command structure using Cobra
* Docker support
* GitHub Actions CI pipeline
* Beginner-friendly project layout

---

# Project Structure

```text
opsctl/
├── cmd/
│   ├── root.go
│   ├── cluster.go
│   ├── pods.go
│   └── version.go
├── internal/
├── .github/
│   └── workflows/
│       └── go.yml
├── Dockerfile
├── Makefile
├── go.mod
├── main.go
└── README.md
```

---

# Requirements

* Go 1.24+
* Docker (optional)

---

# Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/opsctl.git
cd opsctl
```

Install dependencies:

```bash
go mod tidy
```

Build the project:

```bash
go build -o opsctl
```

---

# Usage

## Show Version

```bash
./opsctl version
```

Example output:

```text
opsctl v0.1.0
```

---

# Cluster Health Check

```bash
./opsctl cluster check
```

Example output:

```text
Cluster status: HEALTHY
API Server: Reachable
Nodes Ready: 3/3
```

---

# List Failed Pods

```bash
./opsctl pods failed
```

Example output:

```text
NAMESPACE     POD                  STATUS
default       nginx-xyz            CrashLoopBackOff
monitoring    grafana-0            Error
```

---

# Run Without Building

```bash
go run main.go version
```

---

# Docker Support

Build the image:

```bash
docker build -t opsctl .
```

Run the container:

```bash
docker run opsctl
```

---

# Makefile Commands

Build:

```bash
make build
```

Run:

```bash
make run
```

Docker build:

```bash
make docker-build
```

Clean binaries:

```bash
make clean
```

---

# GitHub Actions CI

The project includes a CI workflow located at:

```text
.github/workflows/go.yml
```

It automatically:

* installs Go
* validates dependencies
* builds the project

---

# Future Improvements

* Real Kubernetes API integration
* Prometheus metrics support
* Helm validation checks
* SSH automation
* OpenTelemetry tracing
* Multi-cluster support
* YAML linting
* JSON export reports

---

# Why This Project Matters

This project demonstrates:

* Golang development
* CLI engineering
* DevOps tooling
* Infrastructure automation
* CI/CD basics
* Containerization
* Clean project organization

---

# Resume Description

> Built a DevOps CLI tool in Go using Cobra to automate cluster diagnostics and infrastructure health checks with Docker and CI/CD integration.

---

# Tech Stack

* Go
* Cobra
* Docker
* GitHub Actions
* Linux CLI tooling

---

# License

MIT License
