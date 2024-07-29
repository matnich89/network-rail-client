# Network Rail Client

This Go package provides a client for connecting to and consuming data from the Network Rail data feeds. It handles Real-Time Public Performance Measure (RTPPM) and Train Movement data feeds.

## Installation

To use this package in your Go project, run:

```bash
go get github.com/matnich89/network-rail-client
```

## Features

* Easy connection to Network Rail's STOMP server
* Automatic handling of subscription and message parsing for multiple data feeds
* Support for subscribing to all train movements or specific train operating companies
* Graceful shutdown and resource cleanup
* Error handling through a dedicated error channel

## API

### `NewNetworkRailClient(ctx context.Context, username, password string) (*NetworkRailClient, error)`

Creates a new Network Rail client and establishes a connection to the STOMP server.

### `(nr *NetworkRailClient) SubRTPPM() (chan *realtime.RTPPMDataMsg, error)`

Subscribes to the RTPPM feed and returns a channel of RTPPM messages.

### `(nr *NetworkRailClient) SubAllTrainMovement() (<-chan movement.Body, error)`

Subscribes to all train movement messages and returns a channel of movement data.

### `(nr *NetworkRailClient) SubMultiTrainCompanyMovements(operators []model.TrainOperator) ([]*TrainCompanySub, error)`

Subscribes to train movement messages for specific train operating companies and returns a slice of TrainCompanySub structures, each containing a channel for that company's movement data.

## Data Models

The package includes several data models:

* `realtime.RTPPMDataMsg`: Represents the RTPPM data structure.
* `movement.Body`: Interface for different types of train movement messages.
* `model.TrainOperator`: Represents a train operating company.

## Usage

Here's a basic example of how to use the client:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/matnich89/network-rail-client/client"
    "github.com/matnich89/network-rail-client/model"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Create a new client
    nrClient, err := client.NewNetworkRailClient(ctx, "your-username", "your-password")
    if err != nil {
        log.Fatal(err)
    }

    // Subscribe to the RTPPM feed
    rtppmChan, err := nrClient.SubRTPPM()
    if err != nil {
        log.Fatal(err)
    }

    // Subscribe to specific train operating companies
    operators := []model.TrainOperator{model.AvantiWestCoast, model.GreatWesternRailway}
    trainSubs, err := nrClient.SubMultiTrainCompanyMovements(operators)
    if err != nil {
        log.Fatal(err)
    }

    // Set up signal handling for graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    // Process incoming messages
    for {
        select {
        case msg := <-rtppmChan:
            fmt.Printf("Received RTPPM message: %+v\n", msg)
        case <-sigChan:
            fmt.Println("Received termination signal. Shutting down...")
            cancel()
            return
        case err := <-nrClient.ErrCh:
            fmt.Printf("Error: %v\n", err)
        default:
            for _, sub := range trainSubs {
                select {
                case mov := <-sub.SubChan:
                    fmt.Printf("Received movement for %s: %+v\n", sub.Name, mov)
                default:
                    // No message available, continue to next subscription
                }
            }
        }
    }
}
```

## Dependencies

* github.com/go-stomp/stomp/v3

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This client is unofficial and is not affiliated with Network Rail. Users must comply with Network Rail's terms of service when using this client.
