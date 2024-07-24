package zoneid

import (
	"cloudflare_ClearCache/settings"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CloudflareResponse struct {
	Result  []Zone  `json:"result"`
	Success bool    `json:"success"`
	Errors  []Error `json:"errors"`
}

type Zone struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Getzid(domain string) (zoneID string) {
	// 替换为你的API信息
	apiEmail := settings.Conf.Email
	apiKey := settings.Conf.Token

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones?name=%s", domain)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Add("X-Auth-Email", apiEmail)
	req.Header.Add("X-Auth-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response code")
		os.Exit(1)
	}

	var cfResponse CloudflareResponse
	err = json.NewDecoder(resp.Body).Decode(&cfResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	if !cfResponse.Success {
		fmt.Println("Failed to fetch Zone ID")
		for _, e := range cfResponse.Errors {
			fmt.Printf("Error: %d - %s\n", e.Code, e.Message)
		}
		os.Exit(1)
	}

	if len(cfResponse.Result) > 0 {
		zoneID := cfResponse.Result[0].ID
		return fmt.Sprintf(zoneID)
	} else {
		fmt.Println("No zones found for the specified domain")
		return "nil"
	}
}
