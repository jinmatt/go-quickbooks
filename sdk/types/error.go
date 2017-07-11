package types

// Error QB error response type
type Error struct {
	Fault struct {
		Error []struct {
			Message string `json:"Message"`
			Detail  string `json:"Detail"`
			Code    string `json:"code"`
		} `json:"Error"`
		Type string `json:"type"`
	} `json:"Fault"`
	Time string `json:"time"`
}

// IntuitResponse xml error response from QB
type IntuitResponse struct {
	Fault fault `xml:"Fault"`
}

type fault struct {
	Type  string     `xml:"type,attr"`
	Error faultError `xml:"Error"`
}

type faultError struct {
	Code    string `xml:"code,attr"`
	Message string `xml:"Message"`
	Detail  string `xml:"Detail"`
}
