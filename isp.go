package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var gispType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoIP2ISP",
		Fields: graphql.Fields{
			"AutonomousSystemNumber": &graphql.Field{
				Type: graphql.String,
			},
			"AutonomousSystemOrganization": &graphql.Field{
				Type: graphql.String,
			},
			"ISP": &graphql.Field{
				Type: graphql.String,
			},
			"Organization": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var gispField = &graphql.Field{
	Type:        gispType,
	Description: "Get ISP",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetISP,
}

func GetISP(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data, err := db.ISP(ip)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
