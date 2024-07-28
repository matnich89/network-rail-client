package main

import (
	"context"
	"github.com/matnich89/network-rail-client/client"
	"log"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	nrClient, err := client.NewNetworkRailClient(ctx, "mathewnicholls@protonmail.com", "Tsunami123@")

	if err != nil {
		log.Fatalln(err)
	}

	realTimeChan, err := nrClient.SubRTPPM()

	if err != nil {
		log.Fatalln(err)
	}

	movementChan, err := nrClient.SubAllTrainMovement()

	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case realTimeMsg := <-realTimeChan:
			log.Println(realTimeMsg)
		case movementMsg := <-movementChan:
			log.Println(movementMsg)
		case <-ctx.Done():
			log.Println("time out triggered")
			return
		}
	}

}
