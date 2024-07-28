package movement

import (
	"encoding/json"
)

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
	Departure   PlannedEventType = "DEPARTURE"
	Destination PlannedEventType = "DESTINATION"
)

type Header struct {
	MsgType MsgType `json:"msg_type"`
}

type Body interface {
	GetType() MsgType
}

type Message struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

type TrainActivationBody struct {
	ScheduleSource     string `json:"schedule_source"`
	TrainFileAddress   string `json:"train_file_address"`
	ScheduleEndDate    string `json:"schedule_end_date"`
	TrainID            string `json:"train_id"`
	TPOriginTimestamp  string `json:"tp_origin_timestamp"`
	CreationTimestamp  string `json:"creation_timestamp"`
	TPOriginStanox     string `json:"tp_origin_stanox"`
	OriginDepTimestamp string `json:"origin_dep_timestamp"`
	TrainServiceCode   string `json:"train_service_code"`
	TocID              string `json:"toc_id"`
	D1266RecordNumber  string `json:"d1266_record_number"`
	TrainCallType      string `json:"train_call_type"`
	TrainUID           string `json:"train_uid"`
	TrainCallMode      string `json:"train_call_mode"`
	ScheduleType       string `json:"schedule_type"`
	SchedOriginStanox  string `json:"sched_origin_stanox"`
	ScheduleWTTID      string `json:"schedule_wtt_id"`
	ScheduleStartDate  string `json:"schedule_start_date"`
}

func (ta *TrainActivationBody) GetType() MsgType {
	return TrainActivation
}

type TrainCancellationBody struct {
	TrainFileAddress   string `json:"train_file_address"`
	TrainServiceCode   string `json:"train_service_code"`
	OrigLocStanox      string `json:"orig_loc_stanox"`
	TocID              string `json:"toc_id"`
	DepTimestamp       string `json:"dep_timestamp"`
	DivisionCode       string `json:"division_code"`
	LocStanox          string `json:"loc_stanox"`
	CancellationTime   string `json:"canx_timestamp"`
	CancellationReason string `json:"canx_reason_code"`
	TrainID            string `json:"train_id"`
	OrigLocTimestamp   string `json:"orig_loc_timestamp"`
	CancellationType   string `json:"canx_type"`
}

func (tc *TrainCancellationBody) GetType() MsgType {
	return TrainCancellation
}

type TrainMovementBody struct {
	EventType            string `json:"event_type"`
	GbttTimestamp        string `json:"gbtt_timestamp,omitempty"`
	OriginalLocStanox    string `json:"original_loc_stanox,omitempty"`
	OriginalLocTimestamp string `json:"original_loc_timestamp,omitempty"`
	PlannedTimestamp     string `json:"planned_timestamp"`
	TimetableVariation   string `json:"timetable_variation"`
	CurrentTrainID       string `json:"current_train_id,omitempty"`
	DelayMonitoringPoint string `json:"delay_monitoring_point"`
	NextReportRunTime    string `json:"next_report_run_time,omitempty"`
	ReportingStanox      string `json:"reporting_stanox,omitempty"`
	ActualTimestamp      string `json:"actual_timestamp"`
	CorrectionInd        string `json:"correction_ind"`
	EventSource          string `json:"event_source"`
	TrainFileAddress     string `json:"train_file_address,omitempty"`
	Platform             string `json:"platform,omitempty"`
	DivisionCode         string `json:"division_code"`
	TrainTerminated      string `json:"train_terminated"`
	TrainID              string `json:"train_id"`
	OffrouteInd          string `json:"offroute_ind"`
	VariationStatus      string `json:"variation_status"`
	TrainServiceCode     string `json:"train_service_code"`
	TocID                string `json:"toc_id"`
	LocStanox            string `json:"loc_stanox"`
	AutoExpected         string `json:"auto_expected,omitempty"`
	DirectionInd         string `json:"direction_ind,omitempty"`
	Route                string `json:"route,omitempty"`
	PlannedEventType     string `json:"planned_event_type"`
	NextReportStanox     string `json:"next_report_stanox,omitempty"`
	LineInd              string `json:"line_ind,omitempty"`
}

func (tm *TrainMovementBody) GetType() MsgType {
	return TrainMovement
}

type TrainReinstatementBody struct {
	TrainID                string `json:"train_id"`
	CurrentTrainID         string `json:"current_train_id"`
	OriginalLocTimestamp   string `json:"original_loc_timestamp"`
	DepTimestamp           string `json:"dep_timestamp"`
	LocStanox              string `json:"loc_stanox"`
	OriginalLocStanox      string `json:"original_loc_stanox"`
	ReinstatementTimestamp string `json:"reinstatement_timestamp"`
	TocID                  string `json:"toc_id"`
	DivisionCodeID         string `json:"division_code_id"`
	TrainFileAddress       string `json:"train_file_address"`
	TrainServiceCode       string `json:"train_service_code"`
}

func (tm *TrainReinstatementBody) GetType() MsgType {
	return TrainReinstatement
}

type TrainChangeOfOriginBody struct {
	TrainID              string `json:"train_id"`
	DepTimestamp         string `json:"dep_timestamp"`
	LocStanox            string `json:"loc_stanox"`
	OriginalLocStanox    string `json:"original_loc_stanox"`
	OriginalLocTimestamp string `json:"original_loc_timestamp"`
	CurrentTrainID       string `json:"current_train_id"`
	TrainServiceCode     string `json:"train_service_code"`
	ReasonCode           string `json:"reason_code"`
	DivisionCode         string `json:"division_code"`
	TocID                string `json:"toc_id"`
	TrainFileAddress     string `json:"train_file_address"`
	CooTimestamp         string `json:"coo_timestamp"`
}

func (tc *TrainChangeOfOriginBody) GetType() MsgType {
	return TrainChangeOfOrigin
}

type TrainChangeOfIdentityBody struct {
	TrainID          string `json:"train_id"`
	CurrentTrainID   string `json:"current_train_id"`
	RevisedTrainID   string `json:"revised_train_id"`
	TrainFileAddress string `json:"train_file_address"`
	TrainServiceCode string `json:"train_service_code"`
	EventTimestamp   string `json:"event_timestamp"`
}

func (tc *TrainChangeOfIdentityBody) GetType() MsgType {
	return TrainChangeOfIdentity
}

type TrainChangeOfLocationBody struct {
	TrainID              string `json:"train_id"`
	CurrentTrainID       string `json:"current_train_id"`
	DepTimestamp         string `json:"dep_timestamp"`
	LocStanox            string `json:"loc_stanox"`
	OriginalLocStanox    string `json:"original_loc_stanox"`
	OriginalLocTimestamp string `json:"original_loc_timestamp"`
	TrainServiceCode     string `json:"train_service_code"`
	TrainFileAddress     string `json:"train_file_address"`
	EventTimestamp       string `json:"event_timestamp"`
}

func (tc *TrainChangeOfLocationBody) GetType() MsgType {
	return TrainChangeOfLocation
}

func Convert(body interface{}, msgType MsgType) Body {
	bodyMap, ok := body.(map[string]interface{})
	if !ok {
		return nil
	}

	var result Body
	switch msgType {
	case TrainActivation:
		result = &TrainActivationBody{}
	case TrainCancellation:
		result = &TrainCancellationBody{}
	case TrainMovement:
		result = &TrainMovementBody{}
	case TrainReinstatement:
		result = &TrainReinstatementBody{}
	case TrainChangeOfOrigin:
		result = &TrainChangeOfOriginBody{}
	case TrainChangeOfIdentity:
		result = &TrainChangeOfIdentityBody{}
	case TrainChangeOfLocation:
		result = &TrainChangeOfLocationBody{}
	default:
		return nil
	}

	jsonData, err := json.Marshal(bodyMap)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return nil
	}

	return result
}
