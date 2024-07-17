# Network Rail Client

This Go package provides a client for connecting to and consuming data from the Network Rail data feeds. It specifically handles the Real-Time Public Performance Measure (RTPPM) feed.

## Installation

To use this package in your Go project, run:

```bash
go get github.com/matnich89/network-rail-client

```

## Features

* Easy connection to Network Rail's STOMP server
* Automatic handling of subscription and message parsing
* Graceful shutdown and resource cleanup

## API

### `NewNetworkRailClient(username, password string) (*NetworkRailClient, error)`

Creates a new Network Rail client and establishes a connection to the STOMP server.

### `(nr *NetworkRailClient) Disconnect() error`

Disconnects the client from the STOMP server and cleans up resources.

### `(nr *NetworkRailClient) SubRTPPM() (chan *model.RTPPMDataMsg, error)`

Subscribes to the RTPPM feed and returns a channel of RTPPM messages.

## Data Model

The `model` package contains structs representing the RTPPM data. The main struct is `RTPPMDataMsg`, which contains detailed information retail time info about national train performance.

## Usage 
Here is a very basic example of how to use the client 

```go
package main

import (
	"fmt"
	"log"
	"github.com/matnich89/network-rail-client/client"
)

func main() {
	// Create a new client
	nrClient, err := client.NewNetworkRailClient("your-username", "your-password")
	if err != nil {
		log.Fatal(err)
	}
	defer nrClient.Disconnect()

	// Subscribe to the RTPPM feed
	rtppmChan, err := nrClient.SubRTPPM()
	if err != nil {
		log.Fatal(err)
	}

	// Process incoming messages
	for msg := range rtppmChan {
		fmt.Printf("Received RTPPM message: %+v\n", msg)
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