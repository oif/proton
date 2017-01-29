package gdns

import (
	"encoding/json"
	"github.com/miekg/dns"
	"net"
)

type GoogleDNSResponse struct {
	Status           int           `json:"Status"`             // 0 success, 2 fail
	TC               bool          `json:"TC"`                 // Whether the response is truncated
	RD               bool          `json:"RD"`                 // Always true for Google Public DNS
	RA               bool          `json:"RA"`                 // Always true for Google Public DNS
	AD               bool          `json:"AD"`                 // Whether all response data was validated with DNSSEC
	CD               bool          `json:"CD"`                 // Whether the client asked to disable DNSSEC
	Question         []Question    `json:"Question"`           // Question
	Answer           []Answer      `json:"Answer"`             // Answer
	Additional       []interface{} `json:"Additional"`         // Additional response
	EdnsClientSubnet string        `json:"edns_client_subnet"` // IP address / scope prefix-length
	Comment          string        `json:"Comment"`            // comment
}

type Question struct {
	Name string `json:"name"` // FQDN with trailing dot
	Type uint32 `json:"type"` // Standard DNS RR type
}

type Answer struct {
	Name string `json:"name"` // Always matches name in the Question section
	Type uint16 `json:"type"` // Standard DNS RR type
	TTL  uint32 `json:"TTL"`  // Record's time-to-live in seconds
	Data string `json:"data"` // IP address as text
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
	case dns.TypeNS:
		return &dns.NS{
			Hdr: a.GetRRHeader(),
			Ns:  a.Data,
		}
	default:
		return &dns.TXT{
			Hdr: dns.RR_Header{
				Name:   a.Name,
				Rrtype: dns.TypeTXT,
				Class:  dns.ClassINET,
				Ttl:    0,
			},
			Txt: []string{"do not support TYPE: " + dns.TypeToString[a.Type] + " currently"},
		}
	}
}

func (a *Answer) GetRRHeader() dns.RR_Header {
	return dns.RR_Header{
		Name:   a.Name,
		Rrtype: a.Type,
		Class:  dns.ClassINET,
		Ttl:    a.TTL,
	}
}
