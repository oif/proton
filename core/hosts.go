package core

import (
	log "github.com/Sirupsen/logrus"
	"github.com/miekg/dns"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var hostsCache map[string]*dns.A

func readHosts(addr string) string {
	if addr == "" {
		addr = "https://raw.githubusercontent.com/racaljk/hosts/master/hosts"
	}
	resp, err := http.Get(addr)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Error(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
	}
	return string(body)
}

func loadHostToCache(h []string) {
	temp := []string{}

	for _, x := range h {
		if len(x) > 0 && x[0] != '#' {
			temp = strings.Fields(x)
			if len(temp) == 2 {
				temp[1] += "."
				hostsCache[temp[1]] = &dns.A{
					Hdr: dns.RR_Header{
						Name:   temp[1],
						Rrtype: dns.TypeA,
						Class:  dns.ClassINET,
						Ttl:    60,
					},
					A: net.ParseIP(temp[0]),
				}
			}
		}
	}
}

func refreshHost(hostFileAddr string) {
	hostsCache = make(map[string]*dns.A)

	rawHost := readHosts(hostFileAddr)
	hosts := strings.Split(rawHost, "\n")
	loadHostToCache(hosts)
}
