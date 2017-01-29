package gdns

// GoogleDNSAPI used in query api
const GoogleDNSAPI = "https://dns.google.com/"

// API request params
const (
	DomainName    = "name"               // 解析域名
	RRType        = "type"               // 解析类型
	DisableDNSSEC = "cd"                 // 关闭 DNSSEC
	EDNS          = "edns_client_subnet" // EDNS
)

// GoogleDNSRequest parse request data to struct
type GoogleDNSRequest struct {
	Name             string // 解析域名
	Type             uint16 // 解析类型
	CD               bool   // 关闭 DNSSEC，默认关闭
	EDNSClientSubnet string // EDNS 子网
}

// NewGoogleDNSRequest will return a GoogleDNSRequest isntance with default value
func NewGoogleDNSRequest() *GoogleDNSRequest {
	return &GoogleDNSRequest{
		Type: 1,
		CD:   false,
	}
}

// ResolveName  设置解析域名
func (g *GoogleDNSRequest) ResolveName(domain string) *GoogleDNSRequest {
	g.Name = domain
	return g
}

// ResolveType 设置解析类型
func (g *GoogleDNSRequest) ResolveType(qtype uint16) *GoogleDNSRequest {
	g.Type = qtype
	return g
}

// ClientSubnet 设置子网
func (g *GoogleDNSRequest) ClientSubnet(subnet string) *GoogleDNSRequest {
	g.EDNSClientSubnet = subnet
	return g
}

// DisableDNSSEC DNSSEC 开关
func (g *GoogleDNSRequest) DisableDNSSEC(cd bool) *GoogleDNSRequest {
	g.CD = cd
	return g
}

// Query 请求 Google DNS
func (g *GoogleDNSRequest) Query() (*GoogleDNSResponse, error) {
	params := map[string]interface{}{
		DomainName:    g.Name,
		RRType:        g.Type,
		DisableDNSSEC: g.CD,
		EDNS:          g.EDNSClientSubnet,
	}
	resp, err := QueryAPI(GoogleDNSAPI+"resolve", params)

	var response *GoogleDNSResponse
	if err == nil {
		response, err = BytesToGoogleDNSResponse(resp)
	}

	return response, err
}
