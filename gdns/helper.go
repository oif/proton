package gdns

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// proxyAddr instance of proxy for query API
var proxyAddr *url.URL

// SetProxyAddr set proxyAddr
func SetProxyAddr(prot, addr string, port uint) {
	var err error
	proxyAddr, err = url.Parse(fmt.Sprintf("%s://%s:%d", prot, addr, port))
	if err != nil {
		log.Errorf("set proxy error %v", err.Error())
	}
}

// QueryAPI 请求 API [GET]
func QueryAPI(urlAddr string, params map[string]interface{}) ([]byte, error) {
	request, _ := http.NewRequest("GET", urlAddr+"?"+paramsFormator(params), nil)
	proxy := proxyAddr //url.Parse("http://127.0.0.1:6152")
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		//Timeout: time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New("timeout")
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
