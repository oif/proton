package gdns

import (
	"encoding/json"
	"github.com/miekg/dns"
	"net"
)

type GoogleDNSResponse struct {
	Status           int           `json:"Status"` // 0 success, 2 fail
	TC               bool          `json:"TC"`     // Whether the response is truncated
	RD               bool          `json:"RD"`     // Always true for Google Public DNS
	RA               bool          `json:"RA"`     // Always true for Google Public DNS
	AD               bool          `json:"AD"`     // Whether all response data was validated with DNSSEC
	CD               bool          `json:"CD"`     // Whether the client asked to disable DNSSEC
	Question         []Question    `json:"Question"`
	Answer           []Answer      `json:"Answer"`
	Additional       []interface{} `json:"Additional"`
	EdnsClientSubnet string        `json:"edns_client_subnet"`
	Comment          string        `json:"Comment"`
}

type Question struct {
	Name string `json:"name"`
	Type uint32 `json:"type"`
}

type Answer struct {
	Name string `json:"name"`
	Type uint16 `json:"type"`
	TTL  uint32 `json:"TTL"`
	Data string `json:"data"`
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

func (a *Answer) GetAnswer() dns.RR {
	switch a.Type {
	case dns.TypeA:
		return &dns.A{
			Hdr: a.GetRRHeader(),
			A:   net.ParseIP(a.Data),
		}
	case dns.TypeAAAA:
		return &dns.AAAA{
			Hdr:  a.GetRRHeader(),
			AAAA: net.ParseIP(a.Data),
		}
	case dns.TypeCNAME:
		return &dns.CNAME{
			Hdr:    a.GetRRHeader(),
			Target: a.Data,
		}
	}
	return nil
}

func (a *Answer) GetRRHeader() dns.RR_Header {
	return dns.RR_Header{
		Name:   a.Name,
		Rrtype: a.Type,
		Class:  dns.ClassINET,
		Ttl:    a.TTL,
	}
}
