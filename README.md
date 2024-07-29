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
	"sync"
	"syscall"

	"github.com/matnich89/network-rail-client/client"
	"github.com/matnich89/network-rail-client/model"
	"github.com/matnich89/network-rail-client/model/movement"
	"github.com/matnich89/network-rail-client/model/realtime"
)

func main() {
	// Replace with your actual Network Rail API credentials (MAKE SURE TO USE EMV VARS!!!)
	username := "your-username"
	password := "your-password"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new NetworkRailClient
	nrClient, err := client.NewNetworkRailClient(ctx, username, password)
	if err != nil {
		log.Fatalf("Failed to create NetworkRailClient: %v", err)
	}

	// Create a WaitGroup to manage our goroutines
	var wg sync.WaitGroup

	// Subscribe to RTPPM data
	rtppmChan, err := nrClient.SubRTPPM()
	if err != nil {
		log.Fatalf("Failed to subscribe to RTPPM: %v", err)
	}

	// Subscribe to all train movements
	allMovementsChan, err := nrClient.SubAllTrainMovement()
	if err != nil {
		log.Fatalf("Failed to subscribe to all train movements: %v", err)
	}

	// Subscribe to specific train company movements
	operators := []model.TrainOperator{
		model.ElizabethLine,
		model.CrossCountry,
	}
	companySubChannels, err := nrClient.SubMultiTrainCompanyMovements(operators)
	if err != nil {
		log.Fatalf("Failed to subscribe to company-specific movements: %v", err)
	}

	// Set up a channel to handle shutdown gracefully
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Start goroutines to handle incoming data
	wg.Add(1)
	go handleRTPPM(ctx, &wg, rtppmChan)
	wg.Add(1)
	go handleAllMovements(ctx, &wg, allMovementsChan)
	for _, companySub := range companySubChannels {
		wg.Add(1)
		go handleCompanyMovements(ctx, &wg, companySub)
	}

	// Handle errors
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-nrClient.ErrCh:
				log.Printf("Error: %v", err)
			case <-ctx.Done():
				return
			}
		}
	}()

	// Wait for shutdown signal
	<-shutdown
	fmt.Println("Shutting down...")
	cancel()

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Shutdown complete")
}

func handleRTPPM(ctx context.Context, wg *sync.WaitGroup, rtppmChan <-chan *realtime.RTPPMDataMsg) {
	defer wg.Done()
	for {
		select {
		case msg := <-rtppmChan:
			fmt.Printf("Received RTPPM data: %+v\n", msg)
		case <-ctx.Done():
			return
		}
	}
}

func handleAllMovements(ctx context.Context, wg *sync.WaitGroup, movementsChan <-chan movement.Body) {
	defer wg.Done()
	for {
		select {
		case movement := <-movementsChan:
			fmt.Printf("Received movement: %+v\n", movement)
		case <-ctx.Done():
			return
		}
	}
}

func handleCompanyMovements(ctx context.Context, wg *sync.WaitGroup, companySub *client.TrainCompanySub) {
	defer wg.Done()
	for {
		select {
		case movement := <-companySub.SubChan:
			fmt.Printf("Received %s movement: %+v\n", companySub.Name, movement)
		case <-ctx.Done():
			return
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
