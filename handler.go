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

func GetJsonWithIP(c *gin.Context, ip_str string) {
	ip := ParseIP(ip_str)
	data := GetData(ip)

	// Write response
	s, _ := json.MarshalIndent(data, "", "  ")
	c.String(200, string(s)+"\n")
}

func Dispatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		isBrowser := IsBrowser(c)
		path := c.Request.URL.String()
		list := strings.Split(path, "/")
		r := list[1]
		a := ""
		if len(list) > 2 {
			a = list[2]
		}

		if r == "json" && a == "" {
			GetJson(c)
		} else if r == "json" && a != "" {
			GetJsonWithIP(c, a)
		} else if !isBrowser {
			GetIP(c)
		} else {
			// Serve static resources for browser
			file := STATIC + "/" + c.Request.URL.Path
			http.ServeFile(c.Writer, c.Request, file)
		}
	}
}

func IsBrowser(c *gin.Context) bool {
	r, _ := regexp.Compile("Gecko|WebKit|Presto|Trident|EdgeHTML|Blink")
	return r.MatchString(c.Request.UserAgent())
}
