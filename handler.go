package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

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
			if PROD {
				// Serve static resources for browser
				if c.Request.URL.Path == "/" {
					c.FileFromFS("static/index.htm", http.FS(f))
				} else {
					c.FileFromFS("static/"+c.Request.URL.Path, http.FS(f))
				}
			} else {
				PassToNg(c)
			}
		}
	}
}

func IsBrowser(c *gin.Context) bool {
	r, _ := regexp.Compile("Gecko|WebKit|Presto|Trident|EdgeHTML|Blink")
	return r.MatchString(c.Request.UserAgent())
}

func PassToNg(c *gin.Context) {
	url, err := url.Parse(NG_URL)
	if err != nil {
		fmt.Println(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(c.Writer, c.Request)
}

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
