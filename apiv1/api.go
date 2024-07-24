package apiv1

import (
	"cloudflare_ClearCache/cache"
	"cloudflare_ClearCache/settings"
	"cloudflare_ClearCache/zoneid"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestData struct {
	Domain string `json:"domain" binding:"required"`
}

func DomainPost(c *gin.Context) {
	var alldomain []string
	var data RequestData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 获取 domain 的值
	domain := data.Domain

	switch domain {
	case "all":
		alldomain = settings.Conf.Domain
		for d := range alldomain {
			fmt.Println(alldomain[d])
			// 执行获取zonesid
			zonesid := zoneid.Getzid(alldomain[d])
			if zonesid != "nil" {
				fmt.Printf("执行清理缓存任务接口ZoneID是:%s\n", zonesid)
				cache.IfDomain(domain, zonesid)
			} else {
				fmt.Printf("ZoneID未能获取\n")
			}
		}
		c.JSON(http.StatusOK, gin.H{"domain": domain})
		fmt.Println("清理所有缓存", alldomain)
	default:
		c.JSON(http.StatusOK, gin.H{"domain": domain})

		// 执行获取zonesid
		zonesid := zoneid.Getzid(domain)
		if zonesid != "nil" {
			fmt.Printf("执行清理缓存任务接口ZoneID是:%s\n", zonesid)
			cache.IfDomain(domain, zonesid)
		} else {
			fmt.Printf("ZoneID未能获取\n")
		}
	}
}
