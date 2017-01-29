package gdns

const GOOGLE_DNS_API = "https://dns.google.com/" // Google DNS API

// 请求参数
const (
	DOMAIN_NAME    = "name"               // 解析域名
	RR_TYPE        = "type"               // 解析类型
	DISABLE_DNSSEC = "cd"                 // 关闭 DNSSEC
	EDNS           = "edns_client_subnet" // EDNS
)

type GoogleDNSRequest struct {
	Name             string // 解析域名
	Type             uint16 // 解析类型
	CD               bool   // 关闭 DNSSEC，默认关闭
	EDNSClientSubnet string // EDNS 子网
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
	resp, err := QueryAPI(GOOGLE_DNS_API+"resolve", params)

	var response *GoogleDNSResponse
	if err == nil {
		response, err = BytesToGoogleDNSResponse(resp)
	}

	return response, err
}
