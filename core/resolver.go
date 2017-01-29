package core

import (
	log "github.com/Sirupsen/logrus"
	"github.com/miekg/dns"
	"proton/gdns"
)

func Resolver(m *dns.Msg, r *dns.Msg, clientIP string) {
	if r.Question[0].Qtype == dns.TypeANY { // 拒绝
		return
	}

	statistics.Resolve()

	// 从缓存中获取
	dnsCache, err := getDNSCache(r.Question, clientIP)
	if err == nil { // 有缓存
		statistics.Hit()

		m.Answer = dnsCache.Answer
		m.Extra = dnsCache.Extra
		m.Ns = dnsCache.Ns
		return
	}

	response, err := gdns.NewGoogleDNSRequest().ResolveName(r.Question[0].Name).ResolveType(r.Question[0].Qtype).ClientSubnet(clientIP).Query()
	if err != nil {
		log.Errorf("google dns request error %v", err.Error())
		return
	}

	// Success
	if ok, comment := response.Success(); ok {
		for _, ans := range response.Answer {
			m.Answer = append(m.Answer, ans.GetAnswer())
		}
		// 缓存记录
		addDNSCache(m, clientIP)
	} else {
		m.Answer = append(m.Answer, &dns.TXT{
			Hdr: dns.RR_Header{
				Name:   clientIP,
				Rrtype: dns.TypeTXT,
				Class:  dns.ClassINET,
				Ttl:    0,
			},
			Txt: []string{comment},
		})
	}
}
