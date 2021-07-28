package main

import (
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
	c.JSON(200, data)
}

func GetJsonWithIP(c *gin.Context) {
	ip := ParseIP(c.Param("address"))
	data := GetData(ip)

	// Write response
	c.JSON(200, data)
}
