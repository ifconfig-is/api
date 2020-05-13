package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var gasnType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoIP2ASN",
		Fields: graphql.Fields{
			"AutonomousSystemNumber": &graphql.Field{
				Type: graphql.String,
			},
			"AutonomousSystemOrganization": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var gasnField = &graphql.Field{
	Type:        gasnType,
	Description: "Get ASN",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetASN,
}

func GetASN(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data, err := db.ASN(ip)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
