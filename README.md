# README.md
# MapReduce Implementation in Go

A distributed MapReduce implementation using Go, Redis, and RabbitMQ.

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Make (optional)

## Setup and Running

1. Clone the repository:
```bash
git clone <your-repo>
cd mapreduce
```

2. Initialize Go module:
```bash
go mod init mapreduce
go mod tidy
```

3. Build and run with Docker Compose:
```bash
docker-compose up --build
```

## Testing the System

1. Create an input file:
```bash
echo "example input data for testing word count" > input.txt
```

2. Process the file:
```bash
cat input.txt | docker-compose exec master ./master
```

3. View results:
```bash
docker-compose exec redis redis-cli HGETALL map_results
```

## Architecture

- Master Node: Coordinates the MapReduce workflow
- Worker Nodes: Process map and reduce tasks
- Redis: Stores intermediate results
- RabbitMQ: Handles task distribution

## Monitoring

- RabbitMQ Management UI: http://localhost:15672 (guest/guest)
- Redis Commander: http://localhost:8081

## Scaling Workers

To scale the number of workers:
```bash
docker-compose up --scale worker=5
```

# example_input.txt
Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris 
nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in 
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.