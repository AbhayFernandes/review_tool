[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)
[![Maintainability](https://api.codeclimate.com/v1/badges/3b594526ba8ebeef8bcf/maintainability)](https://codeclimate.com/github/AbhayFernandes/review_tool/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/3b594526ba8ebeef8bcf/test_coverage)](https://codeclimate.com/github/AbhayFernandes/review_tool/test_coverage)
[![DeepSource](https://app.deepsource.com/gh/AbhayFernandes/review_tool.svg/?label=active+issues&show_trend=true&token=Uq3ZVq8ITznBv5BSfi6DCH6-)](https://app.deepsource.com/gh/AbhayFernandes/review_tool/)

# Review Tool

## Introduction

Welcome to the Review Tool project! This project consists of several components including a CLI, API, Job Processor, and a web frontend. This README provides details about the commands in the Makefile, the docker-compose setup, and the Dockerfiles used in the project.

## Makefile Commands

The Makefile includes various commands for building, running, testing, and cleaning up the project components. Below are the details of the available commands:

### Build Commands

- `make build`: Build all components (API, Job Processor, CLI, and web).
- `make build-go`: Build all Go components (API, Job Processor, and CLI).
- `make build-cli`: Build the CLI binary.
- `make build-api`: Build the API binary.
- `make build-job-processor`: Build the Job Processor binary.
- `make build-web`: Build the web frontend.
- `make build-docker`: Build all Docker images.

### Run Commands

- `make run-cli`: Run the CLI with arguments.
- `make run-api`: Run the API service.
- `make run-job-processor`: Run the Job Processor service.
- `make run-web`: Run the React frontend.
- `make up`: Start all services with Docker Compose.
- `make down`: Stop all services with Docker Compose.

### Test Commands

- `make test`: Run all tests.

### Clean Commands

- `make clean`: Clean up generated files.

## Project Architecture

The Review Tool project is composed of the following components:

1. **CLI**: A command-line interface for interacting with the Review Tool services. This tool will primarily upload git diffs to the API server with SSH keys as authentication. 
2. **API**: A gRPC-based API service that handles requests and responses, and will orchestrate any jobs/actions needed on code reviews.
    - Authentication will be done by taking an SSH private key from the user and hashing the commit message, which the server will verify using the user's public key, submitted
    through web UI.
3. **Job Processor**: A service responsible for processing jobs on code reviews. Things like static analysis, linting etc. will be done here.
4. **Web**: A React-based web frontend for user interaction, allowing users to view changes, leave comments and approve changes.
5. **ScyllaDB**: A NoSQL database used for storing metadata about every review. The reviews themselves will be stored on the Filesystem.

The components communicate with each other using gRPC and are containerized using Docker. The `docker-compose.yml` file orchestrates the services, ensuring they are built and run in the correct order with the necessary dependencies.

## Note on Go Development:
`/cmd` contains all the standalone packages that are a part of this package and `/pkg` contains any code that is shared between packages in `cmd`. These are things like the proto files
and the go code that is generated. 
