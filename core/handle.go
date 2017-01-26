package core

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	gd "proton/google_dns"
)

func protonHandle(w dns.ResponseWriter, r *dns.Msg) {

	var (
		v4 bool
		a  net.IP
	)
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = true
	if ip, ok := w.RemoteAddr().(*net.UDPAddr); ok {
		a = ip.IP
		v4 = a.To4() != nil
	}
	if ip, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		a = ip.IP
		v4 = a.To4() != nil
	}

	fmt.Printf("recieve resolve %s\n", r.Question[0].Name)

	switch r.Question[0].Qtype {
	case dns.TypeA, dns.TypeAAAA:
		if v4 {
			response, err := gd.NewGoogleDNSRequest().ResolveName(r.Question[0].Name).ResolveType(r.Question[0].Qtype).ClientSubnet(a.String()).Query()
			if err != nil {
				fmt.Printf("google dns request error %v\n", err.Error())
				return
			}
			// Success
			if ok, _ := response.Success(); ok {
				for _, anw := range response.Answer {
					m.Answer = append(m.Answer, &dns.A{
						Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: anw.TTL},
						A:   net.ParseIP(anw.Data),
					})
				}
			}
		} else {

		}
	case dns.TypeCNAME:
	default:
	}
	w.WriteMsg(m)
}
