package main

import (
	"github.com/graphql-go/graphql"
)

var ipType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IP",
		Fields: graphql.Fields{
			"Query": &graphql.Field{
				Type: graphql.String,
			},
			"Address": &graphql.Field{
				Type: graphql.String,
			},
			"IsPrivate": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsIPv4": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsIPv6": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var ipField = &graphql.Field{
	Type:        ipType,
	Description: "Get IP",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetIPInfo,
}

type IPInfo struct {
	Query     string
	Address   string
	IsPrivate bool
	IsIPv4    bool
	IsIPv6    bool
}

func GetIPInfo(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data := IPInfo{
		Query:     p.Args["address"].(string),
		Address:   ip.String(),
		IsPrivate: false,
		IsIPv4:    false,
		IsIPv6:    false,
	}
	return data, nil
}
