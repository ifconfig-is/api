package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var gaipType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoIP2AnonymousIP",
		Fields: graphql.Fields{
			"IsAnonymous": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsAnonymousVPN": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsHostingProvider": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsPublicProxy": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsTorExitNode": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var gaipField = &graphql.Field{
	Type:        gaipType,
	Description: "Get AnonymousIP",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetAIP,
}

func GetAIP(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data, err := db.AnonymousIP(ip)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
