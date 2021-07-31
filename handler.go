package main

import (
	"encoding/json"
	"net"
	"net/http"
	"regexp"
	"strings"

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

func Dispatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		isBrowser := IsBrowser(c)
		path := c.Request.URL.String()
		r := strings.Split(path, "/")[1]

		/*
			if r == "json" {
				//ReplyPrettyJson(c, API)
				GetJson(c)
			} else if !isBrowser {
				//PassToApi(c)
				GetIP(c)
			} else {
				fmt.Println("is browser")
				// Serve static resources for browser
				file := STATIC + "/" + c.Request.URL.Path
				http.ServeFile(c.Writer, c.Request, file)
			}
		*/

		if isBrowser {
			// Serve static resources for browser
			file := STATIC + "/" + c.Request.URL.Path
			http.ServeFile(c.Writer, c.Request, file)
		} else if r == "json" {
			GetJson(c)
		} else {
			GetJson(c)
		}
	}
}

func IsBrowser(c *gin.Context) bool {
	r, _ := regexp.Compile("Gecko|WebKit|Presto|Trident|EdgeHTML|Blink")
	return r.MatchString(c.Request.UserAgent())
}
