# Seam for Go

Control locks, lights, and other internet of things devices with Seam's simple API. Checkout the [documentation](https://docs.getseam.com) or examples.

## Setup
```sh
go get github.com/iamwacko/seamapi-go
go mod tidy
```

## Usage
```go
import "github.com/iamwacko/seamapi-go/api"

func main() {
}
```

## Development

This project is written in Go, so the latest stable version is advised.

- To run tests, run `go test`

Our tests use a seam sandbox environment given by environment variable `SEAM_SANDBOX_API_KEY`. If you want to run the tests, you should first create a sandbox workspace [on your dashboard](https://dashboard.getseam.com).

