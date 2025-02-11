# Key-Value Store

## Overview

This is a high-performance, in-memory key-value store written in Go with support for persistence, TTL (Time-To-Live), LRU caching, batch operations, and both CLI and HTTP API interfaces.

## Features

- **In-Memory Storage**: Fast key-value operations using Go's built-in data structures.
- **Persistent Storage**: Uses BadgerDB to save key-value pairs to disk.
- **TTL Support**: Keys can have expiration times, automatically deleting expired data. If no TTL is provided, the key is stored indefinitely.
- **LRU Cache**: Least Recently Used cache mechanism to manage memory efficiently.
- **Batch Operations**: Allows bulk insert and retrieval of key-value pairs.
- **CLI Support**: A command-line interface for interacting with the store.
- **HTTP API**: RESTful endpoints for setting, getting, and deleting key-value pairs.
- **Unit Tests**: Test cases written to ensure reliability and correctness.
- **Docker Support**: Easily run the key-value store in a Docker container.

## Installation

### Prerequisites

- Go 1.18+
- Docker (optional, for containerized deployment)

### Clone the Repository

```sh
git clone https://github.com/Yashh56/keyValueStore.git
cd keyValueStore
```

### Install Dependencies

```sh
go mod tidy
```

## Usage

### Running the Application

The project supports both API and CLI modes.

#### Start as API Server

```sh
go run main.go api
```

#### Start as CLI

```sh
go run main.go cli
```

### Using the CLI

```sh
> store myKey myValue         # Stores 'myKey' with value 'myValue' indefinitely
> store myKey myValue 10      # Stores 'myKey' with value 'myValue' and TTL of 10 seconds
> get myKey                   # Retrieves value for 'myKey'
> delete myKey                # Deletes 'myKey'
> exit                        # Exits the CLI
```

### Using the HTTP API

#### Set Key-Value Pair (Optional TTL)

```sh
curl -X POST "http://localhost:8080/set" -d '{"key":"name","value":"Yash"}' -H "Content-Type: application/json"
```

or with TTL:

```sh
curl -X POST "http://localhost:8080/set" -d '{"key":"name","value":"Yash","ttl":5}' -H "Content-Type: application/json"
```

#### Get Key-Value Pair

```sh
curl -X GET "http://localhost:8080/get?key=name"
```

#### Delete Key-Value Pair

```sh
curl -X DELETE "http://localhost:8080/delete?key=name"
```

### Batch Operations

#### Batch Set (Optional TTL)

```sh
curl -X POST "http://localhost:8080/batch/set" -d '{"items":{"key1":"value1","key2":"value2"}}' -H "Content-Type: application/json"
```

or with TTL:

```sh
curl -X POST "http://localhost:8080/batch/set" -d '{"items":{"key1":"value1","key2":"value2"}, "ttl":10}' -H "Content-Type: application/json"
```

#### Batch Get

```sh
curl -X POST "http://localhost:8080/batch/get" -d '{"keys":["key1","key2"]}' -H "Content-Type: application/json"
```

## Docker Installation

You can run the key-value store using Docker.

### Build the Docker Image

```sh
docker build -t keyvaluestore .
```

### Run the Container (API Mode)

```sh
docker run -p 8080:8080 keyvaluestore api
```

### Run the Container (CLI Mode)

```sh
docker run -it keyvaluestore cli
```

## Project Structure

```
keyValueStore/
│── internal/
│   ├── store/          # Core key-value store logic
│   ├── ttl/            # TTL handling
│   ├── cache/            # LRU cache implementation
│   ├── persist/        # Persistent storage using BadgerDB
│── api/                # HTTP API implementation
│── cli/                # CLI commands
│── tests/              # Unit tests for the key-value store
│── Dockerfile          # Docker setup
│── main.go             # Entry point
│── go.mod              # Dependencies
```

## Running Tests

Unit tests are included to ensure the correctness of the implementation.

```sh
go test ./...
```

## Future Enhancements

- **Distributed Sharding**: Scale the store across multiple nodes.
- **Replication**: Ensure high availability by replicating data.
- **Logging & Monitoring**: Track store operations and performance.

## License

MIT License

