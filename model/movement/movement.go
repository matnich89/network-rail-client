package movement

type MsgType string
type PlannedEventType string

const (
	TrainActivation       MsgType = "0001"
	TrainCancellation     MsgType = "0002"
	TrainMovement         MsgType = "0003"
	TrainReinstatement    MsgType = "0005"
	TrainChangeOfOrigin   MsgType = "0006"
	TrainChangeOfIdentity MsgType = "0007"
	TrainChangeOfLocation MsgType = "0008"

	Arrival     PlannedEventType = "ARRIVAL"
	Departure                    = "DEPARTURE"
	Destination                  = "DESTINATION"
)

type Header struct {
	MsgType MsgType `json:"msg_type"`
}

type TrainMovementBody struct {
	TrainId            string  `json:"train_id"`
	CurrentTrainId     *string `json:"current_train_id"`
	PlannedTimestamp   int64   `json:"planned_timestamp"`
	TimetableVariation string  `json:"timetable_variation"`
	ActualTimestamp    int64   `json:"actual_timestamp"`
	TocID              string  `json:"toc_id"`
	TrainTerminated    bool    `json:"train_terminated"`
	TrainServiceCode   string  `json:"train_service_code"`
	LocStanox          string  `json:"loc_stanox"`
	NextReportStanox   *string `json:"next_report_stanox"`
	Offroute_ind       bool    `json:"offroute_ind"`
}

type MovementDataMsg struct {
	Header Header `json:"header"`
}
