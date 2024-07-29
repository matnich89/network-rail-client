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

func (nr *NetworkRailClient) SubAllTrainMovement() (chan movement.Body, error) {
	sub, err := nr.stompConnection.Subscribe("/topic/TRAIN_MVT_ALL_TOC", stomp.AckAuto)

	if err != nil {
		return nil, fmt.Errorf("could not subscribe to All Train Movements topic: %v", err)
	}

	nr.allTrainMovementsChan = make(chan movement.Body, 100)

	nr.wg.Add(1)
	go func() {
		defer func() {
			nr.wg.Done()
			close(nr.allTrainMovementsChan)
		}()

		for {
			select {
			case msg := <-sub.C:
				var message []movement.Message
				err = json.Unmarshal(msg.Body, &message)
				if err != nil {
					fmt.Printf("Error unmarshaling message: %v\n", err)
					continue
				}
				for _, m := range message {
					data := movement.Convert(m.Body, m.Header.MsgType)
					nr.allTrainMovementsChan <- data
				}
			case <-nr.ctx.Done():
				log.Println("All Train Movements Ending...")
				return
			}
		}
	}()
	return nr.allTrainMovementsChan, nil
}

func (nr *NetworkRailClient) SubMultiTrainCompanyMovements(operators []model.TrainOperator) ([]*TrainCompanySub, error) {
	if len(operators) == 0 {
		return nil, fmt.Errorf("no operators provided")
	}
	var trainCompanySubs []*TrainCompanySub
	for _, operator := range operators {
		trainCompanySubs = append(trainCompanySubs, &TrainCompanySub{
			TrainOperator: operator,
			SubChan:       make(chan movement.Body, 50),
		})
	}

	for i, trainCompanySub := range trainCompanySubs {
		topic := fmt.Sprintf("/topic/TRAIN_MVT_%s_TOC", trainCompanySub.TOC)
		sub, err := nr.stompConnection.Subscribe(topic, stomp.AckAuto)
		if err != nil {
			return nil, fmt.Errorf("could not subscribe to Train Movements topic: %v", err)
		}
		nr.companyMovementsChans = append(nr.companyMovementsChans, trainCompanySub.SubChan)

		nr.wg.Add(1)
		go func(chanIndex int, companyName string) {
			log.Println("Subscribing to Train Movements Topic:", topic)
			defer func() {
				close(nr.companyMovementsChans[i])
				nr.wg.Done()
			}()

			for {
				select {
				case msg := <-sub.C:
					var message []movement.Message
					err := json.Unmarshal(msg.Body, &message)
					if err != nil {
						fmt.Printf("Error unmarshaling message: %v\n", err)
						continue
					}
					for _, m := range message {
						log.Println(msg.Body)
						data := movement.Convert(m.Body, m.Header.MsgType)
						nr.companyMovementsChans[i] <- data
					}
				case <-nr.ctx.Done():
					log.Printf("Ending movement sub for train company %s", companyName)
					return
				}
			}
		}(i, trainCompanySub.Name)
	}
	return trainCompanySubs, nil
}
