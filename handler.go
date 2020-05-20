package main

import (
	"fmt"
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

func GetSimpleData(c *gin.Context) {
	ip := net.ParseIP(c.ClientIP())
	asn, err := db.ASN(ip)
	if err != nil {
		fmt.Println(err)
	}
	city, err := db.City(ip)
	if err != nil {
		fmt.Println(err)
	}

	data := SimpleData{
		Continent: city.Continent.Names["en"],
		Country:   city.Country.Names["en"],
		City:      city.City.Names["en"],
		Latitude:  city.Location.Latitude,
		Longitude: city.Location.Longitude,
		TimeZone:  city.Location.TimeZone,
		IsEU:      city.Country.IsInEuropeanUnion,
		ASN:       asn.AutonomousSystemNumber,
		ORG:       asn.AutonomousSystemOrganization,
	}

	// Write response
	c.JSON(200, data)
}
