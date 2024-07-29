package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-stomp/stomp/v3"
	"github.com/matnich89/network-rail-client/model"
	"github.com/matnich89/network-rail-client/model/movement"
	"log"
)

type TrainCompanySub struct {
	model.TrainOperator
	SubChan chan movement.Body
}

func (nr *NetworkRailClient) SubAllTrainMovement() (<-chan movement.Body, error) {
	sub, err := nr.stompConnection.Subscribe("/topic/TRAIN_MVT_ALL_TOC", stomp.AckAuto)
	if err != nil {
		return nil, fmt.Errorf("could not subscribe to All Train Movements topic: %w", err)
	}

	nr.allTrainMovementsChan = make(chan movement.Body, 100)

	nr.wg.Add(1)
	go nr.handleSubscription(sub, nr.allTrainMovementsChan)

	return nr.allTrainMovementsChan, nil
}

func (nr *NetworkRailClient) SubMultiTrainCompanyMovements(operators []model.TrainOperator) ([]*TrainCompanySub, error) {
	if len(operators) == 0 {
		return nil, fmt.Errorf("no operators provided")
	}

	var trainCompanySubs []*TrainCompanySub
	for _, operator := range operators {
		trainCompanySub := &TrainCompanySub{
			TrainOperator: operator,
			SubChan:       make(chan movement.Body, 50),
		}
		trainCompanySubs = append(trainCompanySubs, trainCompanySub)

		topic := fmt.Sprintf("/topic/TRAIN_MVT_%s_TOC", operator.TOC)
		sub, err := nr.stompConnection.Subscribe(topic, stomp.AckAuto)
		if err != nil {
			return nil, fmt.Errorf("could not subscribe to Train Movements topic for %s: %w", operator.Name, err)
		}

		nr.companyMovementsChans = append(nr.companyMovementsChans, trainCompanySub.SubChan)

		nr.wg.Add(1)
		go nr.handleSubscription(sub, trainCompanySub.SubChan)
	}

	return trainCompanySubs, nil
}

func (nr *NetworkRailClient) handleSubscription(sub *stomp.Subscription, movementChan chan<- movement.Body) {
	defer func() {
		nr.wg.Done()
		close(movementChan)
	}()

	log.Printf("Subscribing to Train Movements Topic: %s", sub.Destination())

	for {
		select {
		case msg := <-sub.C:
			nr.processMessage(msg, movementChan)
		case <-nr.ctx.Done():
			log.Printf("Ending movement sub for %s", sub.Destination())
			return
		}
	}
}

func (nr *NetworkRailClient) processMessage(msg *stomp.Message, movementChan chan<- movement.Body) {
	var messages []movement.Message
	if err := json.Unmarshal(msg.Body, &messages); err != nil {
		nr.ErrCh <- err
		return
	}

	for _, m := range messages {
		data := movement.Convert(m.Body, m.Header.MsgType)
		movementChan <- data
	}
}
