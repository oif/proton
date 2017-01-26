package core

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	gd "proton/google_dns"
)

func protonHandle(w dns.ResponseWriter, r *dns.Msg) {
	var (
		a net.IP
	)
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = true
	if ip, ok := w.RemoteAddr().(*net.UDPAddr); ok {
		a = ip.IP
	}
	if ip, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		a = ip.IP
	}

	fmt.Printf("%s %s %s\n", a.String(), dns.TypeToString[r.Question[0].Qtype], r.Question[0].Name)

	Resolver(m, r, a.String()) // 开始解析

	w.WriteMsg(m)
}

func Resolver(m *dns.Msg, r *dns.Msg, clientIP string) {
	response, err := gd.NewGoogleDNSRequest().ResolveName(r.Question[0].Name).ResolveType(r.Question[0].Qtype).ClientSubnet(clientIP).Query()
	if err != nil {
		fmt.Printf("google dns request error %v\n", err.Error())
		return
	}
	// Success
	if ok, _ := response.Success(); ok {
		for _, ans := range response.Answer {
			m.Answer = append(m.Answer, ans.GetAnswer())
		}
	}
}
