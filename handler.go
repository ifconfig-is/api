package main

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func HDLR() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ContextHandler(c, c.Writer, c.Request)
	}
}

func GetIPAddress(c *gin.Context) {
	// Write response
	c.String(200, c.ClientIP()+"\n")
}

func GetJsonData(c *gin.Context) {
	ip := net.ParseIP(c.ClientIP())
	data := GetSimpleData(ip)

	// Write response
	c.JSON(200, data)
}

func GetJsonDataAddress(c *gin.Context) {
	ip := ParseIP(c.Param("address"))
	data := GetSimpleData(ip)

	// Write response
	c.JSON(200, data)
}
