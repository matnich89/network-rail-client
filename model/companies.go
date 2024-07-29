package model

type TrainOperator struct {
	Name string
	TOC  string
}

var (
	AllianceRail                     = TrainOperator{Name: "Alliance Rail", TOC: "ZB"}
	AmeyFleetServices                = TrainOperator{Name: "Amey Fleet Services", TOC: "RE"}
	AvantiWestCoast                  = TrainOperator{Name: "Avanti West Coast", TOC: "HF"}
	BalfourBeattyRail                = TrainOperator{Name: "Balfour Beatty Rail", TOC: "RZ"}
	C2C                              = TrainOperator{Name: "c2c", TOC: "HT"}
	CaledonianSleeper                = TrainOperator{Name: "Caledonian Sleeper", TOC: "ES"}
	CarillionRail                    = TrainOperator{Name: "Carillion Rail", TOC: "RB"}
	ChilternRailways                 = TrainOperator{Name: "Chiltern Railways", TOC: "HO"}
	ColasRail                        = TrainOperator{Name: "Colas Rail", TOC: "RG"}
	CrossCountry                     = TrainOperator{Name: "CrossCountry", TOC: "EH"}
	DBCargoCharters                  = TrainOperator{Name: "DB Cargo Charters", TOC: "FM"}
	DBCargoFreight                   = TrainOperator{Name: "DB Cargo Freight", TOC: "WA"}
	DBCargoInternational             = TrainOperator{Name: "DB Cargo International", TOC: "DA"}
	DCRail                           = TrainOperator{Name: "DC Rail", TOC: "PO"}
	DirectRailServices               = TrainOperator{Name: "Direct Rail Services", TOC: "XH"}
	EastMidlandsRailway              = TrainOperator{Name: "East Midlands Railway", TOC: "EM"}
	ElizabethLine                    = TrainOperator{Name: "Elizabeth line", TOC: "EX"}
	EuroporteChannel                 = TrainOperator{Name: "Europorte Channel", TOC: "PT"}
	Eurostar                         = TrainOperator{Name: "Eurostar", TOC: "GA"}
	FfestiniogRailway                = TrainOperator{Name: "Ffestiniog Railway", TOC: "XJ"}
	FreightlinerHeavyHaul            = TrainOperator{Name: "Freightliner Heavy Haul", TOC: "DH"}
	FreightlinerIntermodal           = TrainOperator{Name: "Freightliner Intermodal", TOC: "DB"}
	GBRailfreight                    = TrainOperator{Name: "GB Railfreight", TOC: "PE"}
	GoviaThameslink                  = TrainOperator{Name: "Govia Thameslink Railway", TOC: "ET"}
	GrandCentral                     = TrainOperator{Name: "Grand Central", TOC: "EC"}
	GrandCentralNorthWest            = TrainOperator{Name: "Grand Central (North West)", TOC: "LN"}
	GrandUnionTrains                 = TrainOperator{Name: "Grand Union Trains", TOC: "LF"}
	GreatWesternRailway              = TrainOperator{Name: "Great Western Railway", TOC: "EF"}
	GreaterAnglia                    = TrainOperator{Name: "Greater Anglia", TOC: "EB"}
	HansonAndHallRailServices        = TrainOperator{Name: "Hanson & Hall Rail Services", TOC: "YG"}
	Harsco                           = TrainOperator{Name: "Harsco", TOC: "RT"}
	HeathrowConnect                  = TrainOperator{Name: "Heathrow Connect", TOC: "EE"}
	HeathrowExpress                  = TrainOperator{Name: "Heathrow Express", TOC: "HM"}
	HullTrains                       = TrainOperator{Name: "Hull Trains", TOC: "PF"}
	InternalTesting                  = TrainOperator{Name: "Internal Testing", TOC: "RM"}
	IslandLines                      = TrainOperator{Name: "Island Lines", TOC: "HZ"}
	JSDRailResearchAndDevelopment    = TrainOperator{Name: "JSD Rail Research & Development", TOC: "RR"}
	LeggeInfrastructureServices      = TrainOperator{Name: "Legge Infrastructure Services", TOC: "LG"}
	LocomotiveServices               = TrainOperator{Name: "Locomotive Services", TOC: "LS"}
	LondonNorthEasternRailway        = TrainOperator{Name: "London North Eastern Railway", TOC: "HB"}
	LondonOverground                 = TrainOperator{Name: "London Overground", TOC: "EK"}
	LORAM                            = TrainOperator{Name: "LORAM", TOC: "LC"}
	LULBakerlooLine                  = TrainOperator{Name: "LUL Bakerloo Line", TOC: "XC"}
	LULDistrictLineRichmond          = TrainOperator{Name: "LUL District Line - Richmond", TOC: "XE"}
	LULDistrictLineWimbledon         = TrainOperator{Name: "LUL District Line - Wimbledon", TOC: "XB"}
	Lumo                             = TrainOperator{Name: "Lumo", TOC: "LD"}
	Merseyrail                       = TrainOperator{Name: "Merseyrail", TOC: "HE"}
	NetworkRailOnTrackMachines       = TrainOperator{Name: "Network Rail (On-Track Machines)", TOC: "LR"}
	NetworkRailReservedPathings      = TrainOperator{Name: "Network Rail Reserved Pathings (non-QJ)", TOC: "NR"}
	NetworkRailVirtualFreightCompany = TrainOperator{Name: "Network Rail Virtual Freight Company", TOC: "QJ"}
	NexusTyneAndWearMetro            = TrainOperator{Name: "Nexus (Tyne & Wear Metro)", TOC: "PG"}
	NorthernTrains                   = TrainOperator{Name: "Northern Trains", TOC: "ED"}
	NorthYorkshireMoorsRailway       = TrainOperator{Name: "North Yorkshire Moors Railway", TOC: "PR"}
	OnRouteLogistics                 = TrainOperator{Name: "On Route Logistics", TOC: "PM"}
	PreMetroOperations               = TrainOperator{Name: "Pre Metro Operations", TOC: "PK"}
	RailOperationsGroup              = TrainOperator{Name: "Rail Operations Group", TOC: "PH"}
	SBSwietelskyBabcockRail          = TrainOperator{Name: "SB (Swietelsky Babcock) Rail", TOC: "RD"}
	ScotRail                         = TrainOperator{Name: "ScotRail", TOC: "HA"}
	SercoRailOperations              = TrainOperator{Name: "Serco Rail Operations", TOC: "SD"}
	SLCOperations                    = TrainOperator{Name: "SLC Operations", TOC: "SO"}
	SNCFFreightServices              = TrainOperator{Name: "SNCF Freight Services", TOC: "PS"}
	SouthEastern                     = TrainOperator{Name: "Southeastern", TOC: "HU"}
	Southern                         = TrainOperator{Name: "Southern", TOC: "HW"}
	SouthWesternRailway              = TrainOperator{Name: "South Western Railway", TOC: "HY"}
	SouthYorkshireSupertram          = TrainOperator{Name: "South Yorkshire Supertram", TOC: "SJ"}
	SwanageRailway                   = TrainOperator{Name: "Swanage Railway", TOC: "SP"}
	TransPennineExpress              = TrainOperator{Name: "TransPennine Express", TOC: "EA"}
	TransportForWales                = TrainOperator{Name: "Transport for Wales", TOC: "HL"}
	VaramisRail                      = TrainOperator{Name: "Varamis Rail", TOC: "MV"}
	VintageTrains                    = TrainOperator{Name: "Vintage Trains", TOC: "TY"}
	VirtualEuropeanPaths             = TrainOperator{Name: "Virtual European Paths", TOC: "EU"}
	VolkerRail                       = TrainOperator{Name: "VolkerRail", TOC: "RH"}
	WestCoastRailways                = TrainOperator{Name: "West Coast Railways", TOC: "PA"}
	WestMidlandsTrains               = TrainOperator{Name: "West Midlands Trains", TOC: "EJ"}
)
