package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-stomp/stomp/v3"
	"github.com/matnich89/network-rail-client/model"
	"sync"
	"time"
)

const (
	host = "publicdatafeeds.networkrail.co.uk"
	port = "61618"
)

type NetworkRailClient struct {
	stompConnection *stomp.Conn
	wg              *sync.WaitGroup
	rtppmChan       chan *model.RTPPMDataMsg
	stopChan        chan struct{}
}

func NewNetworkRailClient(username, password string) (*NetworkRailClient, error) {
	conn, err := stomp.Dial("tcp", host+":"+port,
		stomp.ConnOpt.Login(username, password),
		stomp.ConnOpt.HeartBeat(time.Minute, time.Minute),
		stomp.ConnOpt.Host("/"),
	)

	if err != nil {
		return nil, fmt.Errorf("could not connect to Network Rail: %v", err)
	}

	return &NetworkRailClient{stompConnection: conn, wg: &sync.WaitGroup{}, stopChan: make(chan struct{})}, nil
}

func (nr *NetworkRailClient) Disconnect() error {
	close(nr.stopChan)
	nr.wg.Wait()
	return nr.stompConnection.Disconnect()
}

func (nr *NetworkRailClient) SubRTPPM() (chan *model.RTPPMDataMsg, error) {
	sub, err := nr.stompConnection.Subscribe("/topic/RTPPM_ALL", stomp.AckAuto)

	if err != nil {
		return nil, fmt.Errorf("could not subscribe to Topic Rail: %v", err)
	}

	nr.rtppmChan = make(chan *model.RTPPMDataMsg, 10)

	nr.wg.Add(1)
	go func() {
		defer nr.wg.Done()
		defer close(nr.rtppmChan)

		for {
			select {
			case msg := <-sub.C:
				var rtppmMsg model.RTPPMDataMsg
				err := json.Unmarshal(msg.Body, &rtppmMsg)
				if err != nil {
					fmt.Printf("Error unmarshaling message: %v\n", err)
					continue
				}
				nr.rtppmChan <- &rtppmMsg
			case <-nr.stopChan:
				return
			}
		}
	}()

	return nr.rtppmChan, nil
}
