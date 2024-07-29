package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-stomp/stomp/v3"
	"github.com/go-stomp/stomp/v3/frame"
	"github.com/matnich89/network-rail-client/model/movement"
	"github.com/matnich89/network-rail-client/model/realtime"
	"log"
	"sync"
	"time"
)

const (
	host = "publicdatafeeds.networkrail.co.uk"
	port = "61618"
)

type Connection interface {
	Subscribe(destination string, ack stomp.AckMode, opts ...func(frame *frame.Frame) error) (*stomp.Subscription, error)
}

type NetworkRailClient struct {
	stompConnection       Connection
	wg                    *sync.WaitGroup
	rtppmChan             chan *realtime.RTPPMDataMsg
	allTrainMovementsChan chan movement.Body
	companyMovementsChans []chan movement.Body
	ErrCh                 chan error
	ctx                   context.Context
}

func NewNetworkRailClient(ctx context.Context, username, password string) (*NetworkRailClient, error) {
	conn, err := stomp.Dial("tcp", host+":"+port,
		stomp.ConnOpt.Login(username, password),
		stomp.ConnOpt.HeartBeat(time.Minute, time.Minute),
		stomp.ConnOpt.Host("/"),
	)

	if err != nil {
		return nil, fmt.Errorf("could not connect to Network Rail: %v", err)
	}

	client := &NetworkRailClient{stompConnection: conn, wg: &sync.WaitGroup{}, ErrCh: make(chan error, 100), ctx: ctx}

	go func() {
		<-ctx.Done()
		client.wg.Wait()
	}()

	return client, nil
}

func (nr *NetworkRailClient) SubRTPPM() (chan *realtime.RTPPMDataMsg, error) {
	sub, err := nr.stompConnection.Subscribe("/topic/RTPPM_ALL", stomp.AckAuto)

	if err != nil {
		return nil, fmt.Errorf("could not subscribe to RTPPM Topic: %v", err)
	}

	nr.rtppmChan = make(chan *realtime.RTPPMDataMsg, 10)

	nr.wg.Add(1)
	go func() {
		defer func() {
			nr.wg.Done()
			close(nr.rtppmChan)
		}()

		for {
			select {
			case msg := <-sub.C:
				var rtppmMsg realtime.RTPPMDataMsg
				err := json.Unmarshal(msg.Body, &rtppmMsg)
				if err != nil {
					nr.ErrCh <- fmt.Errorf("could not unmarshal RTPPM data: %v", err)
					continue
				}
				nr.rtppmChan <- &rtppmMsg
			case <-nr.ctx.Done():
				log.Println("RTPPM Ending...")
				return
			}
		}
	}()

	return nr.rtppmChan, nil
}
