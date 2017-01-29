package core

import (
	log "github.com/Sirupsen/logrus"
	"github.com/miekg/dns"
	"net"
	"time"
)

// protonHandle handle func of Proton
func protonHandle(w dns.ResponseWriter, r *dns.Msg) {
	start := time.Now().UnixNano()
	var (
		a net.IP
	)
	m := new(dns.Msg)
	m.SetEdns0(4096, true)
	m.SetReply(r)
	m.Compress = true

	if ip, ok := w.RemoteAddr().(*net.UDPAddr); ok {
		a = ip.IP
	}
	if ip, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		a = ip.IP
	}

	Resolver(m, r, a.String()) // 开始解析
	w.WriteMsg(m)              // 响应

	log.Debugf("%s %s %s %fms", a.String(), dns.TypeToString[r.Question[0].Qtype], r.Question[0].Name, float64(time.Now().UnixNano()-start)/1000.0/1000.0)
}
