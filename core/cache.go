package core

import (
	"errors"
	"fmt"
	"github.com/miekg/dns"
	"strings"
)

// 根据格式生成 key
func getKey(name, qtype, ip string) string {
	dot := strings.LastIndex(ip, ".")
	if dot > 0 { // v4
		return fmt.Sprintf(CacheKeyFormat, name, qtype, ip[0:dot])
	}

	colon := strings.LastIndex(ip, ":")
	if colon > 0 { // v6
		return fmt.Sprintf(CacheKeyFormat, name, qtype, ip[0:colon])
	}
	return ""
}

// 增加 DNS 缓存
func addDNSCache(m *dns.Msg, ip string) error {
	if len(m.Answer) > 0 {
		packed, err := m.Pack()
		if err != nil {
			return err
		}
		key := getKey(m.Question[0].Name, dns.TypeToString[m.Question[0].Qtype], ip)
		if key == "" {
			return errors.New("invalid ip " + ip)
		}
		cache.Set([]byte(key), packed, int(m.Answer[0].Header().Ttl))
	}
	return nil
}

// 获取 DNS 缓存
func getDNSCache(q []dns.Question, ip string) (*dns.Msg, error) {
	temp := &dns.Msg{}
	key := getKey(q[0].Name, dns.TypeToString[q[0].Qtype], ip)
	if key == "" {
		return temp, errors.New("invalid ip" + ip)
	}
	got, err := cache.Get([]byte(key))
	if err != nil {
		return temp, err
	}
	// 有缓存
	err = temp.Unpack(got)
	if err != nil {
		return temp, err
	}

	// 更新 TTL
	realTTL, err := cache.TTL([]byte(key))
	if err != nil {
		return temp, err
	}

	for i := 0; i < len(temp.Answer); i++ {
		temp.Answer[i].Header().Ttl = realTTL
	}
	for i := 0; i < len(temp.Extra); i++ {
		temp.Extra[i].Header().Ttl = realTTL
	}
	for i := 0; i < len(temp.Ns); i++ {
		temp.Ns[i].Header().Ttl = realTTL
	}

	// 转换成功
	return temp, nil
}
