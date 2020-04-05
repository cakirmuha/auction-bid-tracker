# auction-bid-tracker

Server for Auction Bid Tracker

## Build

    # Clone
    git clone https://github.com/cakirmuha/auction-bid-tracker.git
    cd auction-bid-tracker

    export GO111MODULE=on

    # Generate assets
    go generate ./cmd/auction-bid-tracker/assets/gen.go 

    # Compile
    go build ./cmd/auction-bid-tracker

## Run
   
    Set up env vars, then:
   
    ./auction-bid-tracker

## Options

    Usage of ./auction-bid-tracker:
      -listen string
            Listen addr (default ":8181")
      -log string
            Log level (debug, info, warn, error) (default "debug")
            
## Concurrency Test

     go test ./cmd/bid.tracker/concurrency_test.go
            
## Generate

Generated API spec (openapi.yml) is served at `/api/v1/assets/openapi.yml`

#### Go Vet

    go vet ./...