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

	realIP := getInternalIP(&a)
	Resolver(m, r, realIP) // 开始解析
	w.WriteMsg(m)          // 响应

	log.Debugf("%s %s %s %fms", realIP, dns.TypeToString[r.Question[0].Qtype], r.Question[0].Name, float64(time.Now().UnixNano()-start)/1000.0/1000.0)
}

func getInternalIP(ip *net.IP) string {
	if isPublicIP(ip) {
		return ip.String()
	}
	return servicePublicIP
}

func isPublicIP(ip *net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	if ip6 := ip.To16(); ip6 != nil {
		if ip6[0] == 254 && ip6[1] == 192 {
			return false
		}
		return true
	}
	return false
}
