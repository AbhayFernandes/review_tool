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
3. **Job Processor**: A service responsible for processing jobs on code reviews. Things like static analysis, linting etc. will be done here.
4. **Web**: A React-based web frontend for user interaction, allowing users to view changes, leave comments and approve changes.
5. **ScyllaDB**: A NoSQL database used for storing metadata about every review. The reviews themselves will be stored on the Filesystem.

The components communicate with each other using gRPC and are containerized using Docker. The `docker-compose.yml` file orchestrates the services, ensuring they are built and run in the correct order with the necessary dependencies.
