package utils

import (
	"fmt"
	"github.com/avct/uasurfer"
	"github.com/gin-gonic/gin"
)

type UserAgent struct {
	Ip             string `bson:"ip" json:"ip"`
	UserAgent      string `bson:"user_agent" json:"user_agent"`
	Browser        string `bson:"browser" json:"browser"`
	BrowserVersion string `bson:"browser_version" json:"browser_version"`
	OS             string `bson:"os" json:"os"`
	DeviceType     string `bson:"device_type" json:"device_type"`
}

func (userAgent *UserAgent) Constructor(c *gin.Context) *UserAgent {
	uaString := c.GetHeader("User-Agent")
	ua := uasurfer.Parse(uaString)
	return &UserAgent{
		Ip:             c.ClientIP(),
		UserAgent:      uaString,
		Browser:        ua.Browser.Name.String(),
		BrowserVersion: fmt.Sprintf("%v", ua.Browser.Version),
		OS:             ua.OS.Name.String(),
		DeviceType:     ua.OS.Name.String(),
	}
}
