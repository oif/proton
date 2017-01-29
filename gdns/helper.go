package gdns

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// QueryAPI 请求 API [GET]
func QueryAPI(urlAddr string, params map[string]interface{}) ([]byte, error) {
	request, _ := http.NewRequest("GET", urlAddr+"?"+paramsFormator(params), nil)
	proxy, err := url.Parse("http://127.0.0.1:6152")
	if err != nil {
		return nil, err
	}
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
