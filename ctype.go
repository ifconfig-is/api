package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var gctypeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoIP2ConnectionType",
		Fields: graphql.Fields{
			"ConnectionType": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var gctypeField = &graphql.Field{
	Type:        gctypeType,
	Description: "Get ConnectionType",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetCType,
}

func GetCType(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data, err := db.ConnectionType(ip)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
