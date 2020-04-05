# auction-bid-tracker

Server for Auction Bid Tracker

### Build

    # Clone
    git clone https://github.com/cakirmuha/auction-bid-tracker.git
    cd auction-bid-tracker

    export GO111MODULE=on

    # Generate assets
    go generate ./cmd/auction-bid-tracker/assets/gen.go 

    # Compile
    go build ./cmd/auction-bid-tracker

### Run
   
    Set up env vars, then:
   
    ./auction-bid-tracker

### Options

    Usage of ./auction-bid-tracker:
      -listen string
            Listen addr (default ":8181")
      -log string
            Log level (debug, info, warn, error) (default "debug")
            
### Concurrency Test

     go test ./cmd/bid.tracker/concurrency_test.go
            
### Generate

Generated API spec (openapi.yml) is served at `/api/v1/assets/openapi.yml`

### Go Vet

    go vet ./...

## Solution Instruction

In this task, we have 3 objects(User, Item, Bid). 
100 mock users and 4 mock items are created by starting the application. Map data structure is used to access all 3 objects. 
Maps are not thread-safe by default in golang, so sync.Map can be used to handle concurrency problems. Default map is prefered to handle concurrency issues by code in some part of the code.
- User map(to access user elements for a user)    - default map: sync.RWMutex is used to read user data.
  - key: userid, value: User object
- Item map(to access item elements for an item)   - default map: sync.RWMutex is used to read item data.
  - key: itemid, value: Item object
- ItemBid map(to access bid elements for an item) - sync.Map: No need to consider concurrency problems like deadlock, synchronization...
  - key: itemid, value: List of bids
  - List of bids: Linked list is chosen to access bids for an item. When a new bid come, bid is prepended to list, so head of list will be the last bid and/or the bid having the highest amount. Slice can also be used for list of birds, but it will be sorted by ascending(not so important since we always need to access the highest element to check bid amount, and it is easily accessible by both linked list and slice without visiting list). If we want to delete a user and/or remove all bids of user for an item, the cost of delete/add of linked list is lower than slice(it can be thought as the plus of linked list).

### Solution test

Application can tested using Goroutines(eg. `go test ./cmd/bid.tracker/concurrency_test.go`). Bid request is sent by all 100 mock users for an item, there will be no runtime error, and winning bid will be equal to the highest bid amount of users.
