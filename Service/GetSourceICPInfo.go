package Service

import (
	"errors"
	"io"
	"mengdiao/ICP-Finder/utils"
	"net/http"
	"regexp"
	"strings"
)

func GetSourceICPInfo(url string) (map[string]interface{}, error) {
	domain, err := urlREDomain(url)
	if err != nil {
		return nil, err
	}
	body, err := newQuery(domain)
	if err != nil {
		return nil, err
	}
	str := strings.Trim(body, "()")
	m, err := utils.StringToMap(str)
	if err != nil {
		return nil, err
	}
	return m, nil

}

func newQuery(domain string) (string, error) {
	client := &http.Client{}
	url := "https://cgi.urlsec.qq.com/index.php?m=check&a=check&url=" + domain
	req, _ := http.NewRequest("GET", url, nil)
	// 设置Header
	req.Header.Set("Referer", "https://guanjia.qq.com")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	// 函数结束后关闭相关连接
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func urlREDomain(url string) (string, error) {
	pattern, _ := regexp.Compile(`^(?:https?:\/\/)?([^/:]+)`)
	match := pattern.FindStringSubmatch(url)
	if len(match) > 1 {
		return match[1], nil
	} else {
		//匹配不成功 返回err
		return "", errors.New("nothing matched re")
	}
}
