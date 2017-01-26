package google_dns

import (
	"encoding/json"
)

type GoogleDNSResponse struct {
	Status   int  `json:"Status"` // 0 success, 2 fail
	TC       bool `json:"TC"`     // Whether the response is truncated
	RD       bool `json:"RD"`     // Always true for Google Public DNS
	RA       bool `json:"RA"`     // Always true for Google Public DNS
	AD       bool `json:"AD"`     // Whether all response data was validated with DNSSEC
	CD       bool `json:"CD"`     // Whether the client asked to disable DNSSEC
	Question []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"Question"`
	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  uint32 `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`
	Additional       []interface{} `json:"Additional"`
	EdnsClientSubnet string        `json:"edns_client_subnet"`
	Comment          string        `json:"Comment"`
}

func BytesToGoogleDNSResponse(resp []byte) (*GoogleDNSResponse, error) {
	response := GoogleDNSResponse{}
	var err error
	err = json.Unmarshal(resp, &response)
	return &response, err
}

func (g *GoogleDNSResponse) Success() (bool, string) {
	if g.Status == 0 {
		return true, ""
	}
	return false, g.Comment
}
