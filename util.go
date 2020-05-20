package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/oschwald/geoip2-golang"
)

func GetIP(p graphql.ResolveParams) net.IP {
	var ip net.IP
	if p.Args["address"] != nil && p.Args["address"] != "" {
		ip = ParseIP(p.Args["address"].(string))
	} else {
		ctx := p.Context.(*gin.Context)
		ip = net.ParseIP(ctx.ClientIP())
	}
	return ip
}

func ParseIP(address string) net.IP {
	ip := net.ParseIP(address)
	if ip != nil {
		return ip
	} else {
		var host string
		u, err := url.ParseRequestURI(address)
		if err == nil {
			host = u.Host
		} else {
			host = address
		}
		ips, err := net.LookupIP(host)
		if err != nil {
		}
		return ips[0]
	}
}

func PPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func SetupParam(ip string) graphql.ResolveParams {
	var p graphql.ResolveParams
	p.Args = make(map[string]interface{})
	p.Args["address"] = ip
	return p
}

func AddNames(vdata interface{}, vtype string) interface{} {
	if vtype == "city" {
		data := vdata.(*geoip2.City)
		if data.City.Names["pt-BR"] != "" {
			data.City.Names["pt"] = data.City.Names["pt-BR"]
			data.City.Names["zh"] = data.City.Names["zh-CN"]
		}
		if data.Continent.Names["pt-BR"] != "" {
			data.Continent.Names["pt"] = data.Continent.Names["pt-BR"]
			data.Continent.Names["zh"] = data.Continent.Names["zh-CN"]
		}
		if data.Country.Names["pt-BR"] != "" {
			data.Country.Names["pt"] = data.Country.Names["pt-BR"]
			data.Country.Names["zh"] = data.Country.Names["zh-CN"]
		}
		if data.RegisteredCountry.Names["pt-BR"] != "" {
			data.RegisteredCountry.Names["pt"] = data.RegisteredCountry.Names["pt-BR"]
			data.RegisteredCountry.Names["zh"] = data.RegisteredCountry.Names["zh-CN"]
		}
		if data.RepresentedCountry.Names["pt-BR"] != "" {
			data.RepresentedCountry.Names["pt"] = data.RepresentedCountry.Names["pt-BR"]
			data.RepresentedCountry.Names["zh"] = data.RepresentedCountry.Names["zh-CN"]
		}
		for i, _ := range data.Subdivisions {
			if data.Subdivisions[i].Names["pt-BR"] != "" {
				data.Subdivisions[i].Names["pt"] = data.Subdivisions[i].Names["pt-BR"]
				data.Subdivisions[i].Names["zh"] = data.Subdivisions[i].Names["zh-CN"]
			}
		}
		return data
	}
	if vtype == "country" {
		data := vdata.(*geoip2.Country)
		if data.Continent.Names["pt-BR"] != "" {
			data.Continent.Names["pt"] = data.Continent.Names["pt-BR"]
			data.Continent.Names["zh"] = data.Continent.Names["zh-CN"]
		}
		if data.Country.Names["pt-BR"] != "" {
			data.Country.Names["pt"] = data.Country.Names["pt-BR"]
			data.Country.Names["zh"] = data.Country.Names["zh-CN"]
		}
		if data.RegisteredCountry.Names["pt-BR"] != "" {
			data.RegisteredCountry.Names["pt"] = data.RegisteredCountry.Names["pt-BR"]
			data.RegisteredCountry.Names["zh"] = data.RegisteredCountry.Names["zh-CN"]
		}
		if data.RepresentedCountry.Names["pt-BR"] != "" {
			data.RepresentedCountry.Names["pt"] = data.RepresentedCountry.Names["pt-BR"]
			data.RepresentedCountry.Names["zh"] = data.RepresentedCountry.Names["zh-CN"]
		}
		return data
	}
	return nil
}
