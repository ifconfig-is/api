package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var gdomainType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoIP2Domain",
		Fields: graphql.Fields{
			"Domain": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var gdomainField = &graphql.Field{
	Type:        gdomainType,
	Description: "Get Domain",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetDomain,
}

func GetDomain(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data, err := db.Domain(ip)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
