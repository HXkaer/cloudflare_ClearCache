package cache

import (
	"bytes"
	"cloudflare_ClearCache/settings"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 此接口备用 用于判断是否是cf还是aws托管dns域名
func IfDomain(domain string, zonesid string) {
	apiKey := settings.Conf.Token
	apiEmail := settings.Conf.Email
	DelCfCache(domain, apiKey, apiEmail, zonesid)
}

func DelCfCache(domain string, apiKey string, apiEmail string, zonesid string) {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/purge_cache", zonesid)
	method := "POST"

	// JSON 数据
	jsonData := `{"purge_everything":true}`

	// 创建一个新的请求
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求头
	req.Header.Add("X-Auth-Email", apiEmail)
	req.Header.Add("X-Auth-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s delete Cache: %s", domain, string(body))
}
