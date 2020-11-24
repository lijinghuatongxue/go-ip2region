package src

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ip2region/util"
	"net/http"
)

func Ip2Region(c *gin.Context) {
	type IpStruct struct {
		Ip string `form:"ip" json:"ip" uri:"ip" xml:"ip" binding:"required"`
	}

	data := make(map[string]interface{})
	var IpPayload IpStruct
	if err := c.ShouldBindJSON(&IpPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data["IP"] = IpPayload.Ip
	// Get ip'region
	IpRegionInfo, err := util.Ip2Area(IpPayload.Ip)
	if err != true {
		data["Region"] = "Null"
		c.JSON(200, gin.H{
			"code": 500,
			"data": data,
			"msg":  "ok",
		})
	}
	data["IP"] = IpPayload.Ip
	data["Region"] = IpRegionInfo
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "ok",
	})
}
