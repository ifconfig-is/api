package main

import (
	"encoding/json"
	"net"

	"github.com/gin-gonic/gin"
)

func GetIP(c *gin.Context) {
	// Write response
	c.String(200, c.ClientIP()+"\n")
}

func GetJson(c *gin.Context) {
	ip := net.ParseIP(c.ClientIP())
	data := GetData(ip)

	// Write response
	s, _ := json.MarshalIndent(data, "", "  ")
	c.String(200, string(s)+"\n")
}

func GetJsonWithIP(c *gin.Context) {
	ip := ParseIP(c.Param("address"))
	data := GetData(ip)

	// Write response
	s, _ := json.MarshalIndent(data, "", "  ")
	c.String(200, string(s)+"\n")
}
