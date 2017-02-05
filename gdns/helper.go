package gdns

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// client instance of request for query API
var client *http.Client

// SetProxyAddr set proxyAddr
func SetProxyAddr(prot, addr string, port uint) {
	proxyAddr, err := url.Parse(fmt.Sprintf("%s://%s:%d", prot, addr, port))
	if err != nil {
		log.Errorf("set proxy error %v", err.Error())
	}
	client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyAddr),
		},
	}
}

// QueryAPI 请求 API [GET]
func QueryAPI(urlAddr string, params map[string]interface{}) ([]byte, error) {
	request, _ := http.NewRequest("GET", urlAddr+"?"+paramsFormator(params), nil)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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
