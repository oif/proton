package gdns

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// client instance of request for query API
var client = &http.Client{
	Timeout: 5 * time.Second,
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// SetProxyAddr set proxyAddr
func SetProxyAddr(prot, addr string, port uint) {
	proxyAddr, err := url.Parse(fmt.Sprintf("%s://%s:%d", prot, addr, port))
	if err != nil {
		log.Errorf("set proxy error %v", err.Error())
	}
	client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyAddr),
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

// QueryAPI 请求 API [GET]
func QueryAPI(urlAddr string, params map[string]interface{}) ([]byte, error) {
	request, _ := http.NewRequest("GET", urlAddr+"?"+paramsFormator(params), nil)
	resp, err := client.Do(request)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body, err
}

// 构造参数 (URLEncode)
func paramsFormator(params map[string]interface{}) string {
	if params == nil || len(params) == 0 {
		// 无参数
		return ""
	}
	var result string
	for key, val := range params {
		result += fmt.Sprintf("%s=%v&", key, val)
	}
	return strings.TrimRight(result, "&")
}
