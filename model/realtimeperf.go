package model

type RTPPMDataMsg struct {
	RTPPMDataMsgV1 RTPPMDataMsgV1 `json:"RTPPMDataMsgV1"`
}

type RTPPMDataMsgV1 struct {
	RTPPMData      RTPPMData   `json:"RTPPMData"`
	Sender         Sender      `json:"Sender"`
	Publication    Publication `json:"Publication"`
	Classification string      `json:"classification"`
	Timestamp      string      `json:"timestamp"`
	Owner          string      `json:"owner"`
}

type RTPPMData struct {
	RAGThresholds      []RAGThreshold     `json:"RAGThresholds"`
	WebPPMLink         string             `json:"WebPPMLink"`
	PPT                PPT                `json:"PPT"`
	NationalPage       NationalPage       `json:"NationalPage"`
	OOCPage            OOCPage            `json:"OOCPage"`
	FOCPage            FOCPage            `json:"FOCPage"`
	CommonOperatorPage CommonOperatorPage `json:"CommonOperatorPage"`
	OperatorPage       []OperatorPage     `json:"OperatorPage"`
	SnapshotTStamp     string             `json:"snapshotTStamp"`
}

type RAGThreshold struct {
	Type   string `json:"type"`
	Good   string `json:"good"`
	Medium string `json:"medium"`
}

type PPT struct {
	Text           string `json:"text"`
	Rag            string `json:"rag"`
	RagDisplayFlag string `json:"ragDisplayFlag"`
}

type NationalPage struct {
	WebFixedMsg1     string      `json:"WebFixedMsg1"`
	WebFixedMsg2     string      `json:"WebFixedMsg2"`
	StaleFlag        string      `json:"StaleFlag"`
	NationalPPM      NationalPPM `json:"NationalPPM"`
	Sector           []Sector    `json:"Sector"`
	Operator         []Operator  `json:"Operator"`
	WebDisplayPeriod string      `json:"WebDisplayPeriod"`
}

type NationalPPM struct {
	Total          string     `json:"Total"`
	OnTime         string     `json:"OnTime"`
	Late           string     `json:"Late"`
	CancelVeryLate string     `json:"CancelVeryLate"`
	PPM            PPMData    `json:"PPM"`
	RollingPPM     RollingPPM `json:"RollingPPM"`
}

type Sector struct {
	SectorPPM  SectorPPM `json:"SectorPPM"`
	SectorCode string    `json:"sectorCode"`
	SectorDesc string    `json:"sectorDesc"`
}

type SectorPPM struct {
	Total          string     `json:"Total"`
	OnTime         string     `json:"OnTime"`
	Late           string     `json:"Late"`
	CancelVeryLate string     `json:"CancelVeryLate"`
	PPM            PPMData    `json:"PPM"`
	RollingPPM     RollingPPM `json:"RollingPPM"`
}

type Operator struct {
	Total      string     `json:"Total"`
	PPM        PPMData    `json:"PPM"`
	RollingPPM RollingPPM `json:"RollingPPM"`
	Code       string     `json:"code"`
	Name       string     `json:"name"`
	KeySymbol  string     `json:"keySymbol,omitempty"`
}

type OOCPage struct {
	WebFixedMsg1     string     `json:"WebFixedMsg1"`
	Operator         []Operator `json:"Operator"`
	WebDisplayPeriod string     `json:"WebDisplayPeriod"`
}

type FOCPage struct {
	WebFixedMsg2     string      `json:"WebFixedMsg2"`
	NationalPPM      NationalPPM `json:"NationalPPM"`
	Operator         []Operator  `json:"Operator"`
	WebDisplayPeriod string      `json:"WebDisplayPeriod"`
}

type CommonOperatorPage struct {
	WebDisplayPeriod string `json:"WebDisplayPeriod"`
	WebFixedMsg1     string `json:"WebFixedMsg1"`
}

type OperatorPage struct {
	Operator          OperatorData `json:"Operator"`
	OprToleranceTotal interface{}  `json:"OprToleranceTotal"`
	OprServiceGrp     interface{}  `json:"OprServiceGrp"`
}

type PPMData struct {
	Text           string `json:"text"`
	Rag            string `json:"rag"`
	RagDisplayFlag string `json:"ragDisplayFlag,omitempty"`
}

type RollingPPM struct {
	Text        string `json:"text"`
	Rag         string `json:"rag"`
	DisplayFlag string `json:"displayFlag,omitempty"`
	TrendInd    string `json:"trendInd,omitempty"`
}

type OperatorData struct {
	Total          string     `json:"Total"`
	OnTime         string     `json:"OnTime"`
	Late           string     `json:"Late"`
	CancelVeryLate string     `json:"CancelVeryLate"`
	PPM            PPMData    `json:"PPM"`
	RollingPPM     RollingPPM `json:"RollingPPM"`
	Code           string     `json:"code"`
	Name           string     `json:"name"`
	KeySymbol      string     `json:"keySymbol,omitempty"`
}

type Sender struct {
	Organisation string `json:"organisation"`
	Application  string `json:"application"`
}

type Publication struct {
	TopicID string `json:"TopicID"`
}
