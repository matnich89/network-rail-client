package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/matnich89/network-rail-client/model/movement"
	"github.com/matnich89/network-rail-client/model/realtime"
	"sync"
	"testing"
	"time"

	"github.com/go-stomp/stomp/v3"
	"github.com/go-stomp/stomp/v3/frame"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockConnection implements the Connection interface for testing
type MockConnection struct {
	mock.Mock
}

func (m *MockConnection) Subscribe(destination string, ack stomp.AckMode, opts ...func(*frame.Frame) error) (*stomp.Subscription, error) {
	args := m.Called(destination, ack)
	return args.Get(0).(*stomp.Subscription), args.Error(1)
}

func TestSubRTPPM(t *testing.T) {
	mockConn := new(MockConnection)

	// Create a real stomp.Subscription
	subChan := make(chan *stomp.Message, 10)
	sub := &stomp.Subscription{C: subChan}

	client := &NetworkRailClient{
		stompConnection: mockConn,
		ctx:             context.Background(),
		wg:              &sync.WaitGroup{},
	}

	mockConn.On("Subscribe", "/topic/RTPPM_ALL", stomp.AckAuto).Return(sub, nil)

	rtppmChan, err := client.SubRTPPM()
	assert.NoError(t, err)
	assert.NotNil(t, rtppmChan)

	// Simulate receiving a message with the correct structure
	testMsg := &realtime.RTPPMDataMsg{
		RTPPMDataMsgV1: realtime.RTPPMDataMsgV1{
			RTPPMData: realtime.RTPPMData{
				NationalPage: realtime.NationalPage{
					NationalPPM: realtime.NationalPPM{
						PPM: realtime.PPMData{
							Text: "90.0",
							Rag:  "G",
						},
					},
				},
			},
			Sender: realtime.Sender{
				Organisation: "NETVIS",
				Application:  "RTPPM",
			},
			Timestamp: "2023-07-28T12:00:00Z",
		},
	}
	msgBody, _ := json.Marshal(testMsg)
	subChan <- &stomp.Message{Body: msgBody}

	select {
	case receivedMsg := <-rtppmChan:
		assert.Equal(t, "90.0", receivedMsg.RTPPMDataMsgV1.RTPPMData.NationalPage.NationalPPM.PPM.Text)
		assert.Equal(t, "G", receivedMsg.RTPPMDataMsgV1.RTPPMData.NationalPage.NationalPPM.PPM.Rag)
		assert.Equal(t, "NETVIS", receivedMsg.RTPPMDataMsgV1.Sender.Organisation)
		assert.Equal(t, "RTPPM", receivedMsg.RTPPMDataMsgV1.Sender.Application)
		assert.Equal(t, "2023-07-28T12:00:00Z", receivedMsg.RTPPMDataMsgV1.Timestamp)
	case <-time.After(time.Second):
		t.Fatal("Timed out waiting for message")
	}

	mockConn.AssertExpectations(t)
}

func TestSubAllTrainMovement(t *testing.T) {
	mockConn := new(MockConnection)

	subChan := make(chan *stomp.Message, 10)
	sub := &stomp.Subscription{C: subChan}

	client := &NetworkRailClient{
		stompConnection: mockConn,
		ctx:             context.Background(),
		wg:              &sync.WaitGroup{},
	}

	mockConn.On("Subscribe", "/topic/TRAIN_MVT_ALL_TOC", stomp.AckAuto).Return(sub, nil)

	movementChan, err := client.SubAllTrainMovement()
	assert.NoError(t, err)
	assert.NotNil(t, movementChan)

	// Simulate receiving a message with the correct structure
	testMovement := []movement.Message{
		{
			Header: movement.Header{MsgType: movement.TrainMovement},
			Body: map[string]interface{}{
				"event_type": "DEPARTURE",
				"train_id":   "1A23",
			},
		},
	}
	msgBody, _ := json.Marshal(testMovement)
	subChan <- &stomp.Message{Body: msgBody}

	select {
	case receivedMovement := <-movementChan:
		assert.IsType(t, &movement.TrainMovementBody{}, receivedMovement)
		movementBody := receivedMovement.(*movement.TrainMovementBody)
		assert.Equal(t, "DEPARTURE", movementBody.EventType)
		assert.Equal(t, "1A23", movementBody.TrainID)
	case <-time.After(time.Second):
		t.Fatal("Timed out waiting for message")
	}

	mockConn.AssertExpectations(t)
}

func TestSubRTPPM_Error(t *testing.T) {
	mockConn := new(MockConnection)
	client := &NetworkRailClient{
		stompConnection: mockConn,
		ctx:             context.Background(),
		wg:              &sync.WaitGroup{},
	}

	mockConn.On("Subscribe", "/topic/RTPPM_ALL", stomp.AckAuto).Return((*stomp.Subscription)(nil), errors.New("subscription error"))

	_, err := client.SubRTPPM()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not subscribe to RTPPM Topic")

	mockConn.AssertExpectations(t)
}

func TestSubAllTrainMovement_Error(t *testing.T) {
	mockConn := new(MockConnection)
	client := &NetworkRailClient{
		stompConnection: mockConn,
		ctx:             context.Background(),
		wg:              &sync.WaitGroup{},
	}

	mockConn.On("Subscribe", "/topic/TRAIN_MVT_ALL_TOC", stomp.AckAuto).Return((*stomp.Subscription)(nil), errors.New("subscription error"))

	_, err := client.SubAllTrainMovement()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not subscribe to All Train Movements topic")

	mockConn.AssertExpectations(t)
}

func TestContextCancellation(t *testing.T) {
	mockConn := new(MockConnection)

	subChan := make(chan *stomp.Message, 10)
	sub := &stomp.Subscription{C: subChan}

	ctx, cancel := context.WithCancel(context.Background())
	client := &NetworkRailClient{
		stompConnection: mockConn,
		ctx:             ctx,
		wg:              &sync.WaitGroup{},
	}

	mockConn.On("Subscribe", "/topic/RTPPM_ALL", stomp.AckAuto).Return(sub, nil)

	rtppmChan, err := client.SubRTPPM()
	assert.NoError(t, err)

	cancel()

	time.Sleep(100 * time.Millisecond)

	_, ok := <-rtppmChan
	assert.False(t, ok, "Channel should be closed after context cancellation")

	mockConn.AssertExpectations(t)
}
