package yaddress

type Address struct {
	ErrorCode            int     `json:"ErrorCode"`
	ErrorMessage         string  `json:"ErrorMessage"`
	AddressLine1         string  `json:"AddressLine1"`
	AddressLine2         string  `json:"AddressLine2"`
	Number               string  `json:"Number"`
	PreDir               string  `json:"PreDir"`
	Street               string  `json:"Street"`
	Suffix               string  `json:"Suffix"`
	PostDir              string  `json:"PostDir"`
	Sec                  string  `json:"Sec"`
	SecValidated         bool    `json:"SecValidated"`
	City                 string  `json:"City"`
	State                string  `json:"State"`
	Zip                  string  `json:"Zip"`
	Zip4                 string  `json:"Zip4"`
	UspsCarrierRoute     string  `json:"UspsCarrierRoute"`
	County               string  `json:"County"`
	StateFP              string  `json:"StateFP"`
	CountyFP             string  `json:"CountyFP"`
	CensusTract          string  `json:"CensusTract"`
	CensusBlock          string  `json:"CensusBlock"`
	Latitude             float32 `json:"Latitude"`
	Longitude            float32 `json:"Longitude"`
	GeoPrecision         int     `json:"GeoPrecision"`
	TimeZoneOffset       int     `json:"TimeZoneOffset"`
	DstObserved          bool    `json:"DstObserved"`
	PlaceFP              int     `json:"PlaceFP"`
	CityMunicipality     string  `json:"CityMunicipality"`
	SalesTaxRate         float32 `json:"SalesTaxRate"`
	SalesTaxJurisdiction int     `json:"SalesTaxJurisdiction"`
}

type Request struct {
	AddressLine1 string
	AddressLine2 string
}

// YaddressResponse
type YaddressResult struct {
	Result Address
	Debug  struct {
		ErrorCode    int
		ErrorMessage string
	}
}
