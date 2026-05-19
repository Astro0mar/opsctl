# opsctl

Simple DevOps CLI written in Go using Cobra.

## Features

- Cluster health checks
- Failed pod detection
- Version command
- Docker support

## Installation

```bash
go mod tidy
go build -o opsctl
```

## Usage

### Check cluster

```bash
./opsctl cluster check
```

### List failed pods

```bash
./opsctl pods failed
```

### Version

```bash
./opsctl version
```

## Run with Docker

```bash
docker build -t opsctl .
docker run opsctl
```

## Future Improvements

- Kubernetes API integration
- Prometheus metrics
- Helm validation
- GitHub Actions CI/CD
