# Calculator API

A calculator API written in Golang

## Run with Docker

### Prerequisites
   * [Docker](https://docs.docker.com/install/)

### Steps

1. `git clone git@github.com:alohr51/Go-Calculator-API.git`
1. `cd Go-Calculator-API`
1. `docker build -t gocalc .`
1. `docker run --rm -p 8080:8080 gocalc`
1. Open `http://localhost:8080/api/calc/sum/1/2` in your browser to sum 1 + 2, for example.
   * Subtract example
      * `http://localhost:8080/api/calc/subtract/10/5`
   * Multiply example
      * `http://localhost:8080/api/calc/multiply/10/5`
   * Divide example
      * `http://localhost:8080/api/calc/divide/10/5`

## Run with Golang

### Prerequisites
   * [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
   * [Golang](https://golang.org/doc/install)

### Steps

1. `git clone git@github.com:alohr51/Go-Calculator-API.git`
1. `cd Go-Calculator-API`
1. `go run main.go`

### Test

1. From root of project Go-Calculator-API: `go test ./...` 