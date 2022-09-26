# Compy


## Prerequisite
- Go - https://go.dev/doc/install
- Docker - https://docs.docker.com/engine/install/
- Postgres - https://www.postgresql.org/download/

## Starting the server:
There are 2 ways to start the GRPC server:

### Using Docker:
- Run the command `docker-compose up` and it will take care of everything and will start the server in a docker container. Port 9000 of container is exposed and can be used to access the server.

### Using Go:
- Install required go modules using `go mod download`.
- Open the code in a text editor and update the postgres credentials like database name, username and password in `sqlboiler.toml`, `config.toml` and `Makefile`.
- Start the server using `go run main.go`

## Usage
After starting the server, grpc services can be accesssed on port 9000. You can either use:
1. Postman - https://blog.postman.com/postman-now-supports-grpc/
2. Grpcurl - https://github.com/fullstorydev/grpcurl

List of available services can be fetched using server reflection.
