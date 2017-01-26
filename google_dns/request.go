package google_dns

import (
	"fmt"
	"proton/util"
)

const GOOGLE_DNS_API = "https://dns.google.com/"

const (
	DOMAIN_NAME    = "name"
	RR_TYPE        = "type"
	DISABLE_DNSSEC = "cd"
	EDNS           = "edns_client_subnet"
)

type GoogleDNSRequest struct {
	Name             string
	Type             uint16
	CD               bool
	EDNSClientSubnet string
}

func NewGoogleDNSRequest() *GoogleDNSRequest {
	return &GoogleDNSRequest{
		Type: 1,
		CD:   false,
	}
}

func (g *GoogleDNSRequest) ResolveName(domain string) *GoogleDNSRequest {
	g.Name = domain
	return g
}

func (g *GoogleDNSRequest) ResolveType(qtype uint16) *GoogleDNSRequest {
	g.Type = qtype
	return g
}

func (g *GoogleDNSRequest) ClientSubnet(subnet string) *GoogleDNSRequest {
	g.EDNSClientSubnet = subnet
	return g
}

func (g *GoogleDNSRequest) DisableDNSSEC(cd bool) *GoogleDNSRequest {
	g.CD = cd
	return g
}

func (g *GoogleDNSRequest) Query() (*GoogleDNSResponse, error) {
	params := map[string]interface{}{
		DOMAIN_NAME:    g.Name,
		RR_TYPE:        g.Type,
		DISABLE_DNSSEC: g.CD,
		EDNS:           g.EDNSClientSubnet,
	}
	fmt.Println("query google dns api")
	resp, err := util.QueryAPI(GOOGLE_DNS_API+"resolve", params)
	var response *GoogleDNSResponse
	if err == nil {
		response, err = BytesToGoogleDNSResponse(resp)
	}
	return response, err
}
