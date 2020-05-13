package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var genterpriseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoIP2Enterprise",
		Fields: graphql.Fields{
			"City": &graphql.Field{
				Type: cityType,
			},
			"Continent": &graphql.Field{
				Type: continentType,
			},
			"Country": &graphql.Field{
				Type: countryType,
			},
			"Location": &graphql.Field{
				Type: locationType,
			},
			"Postal": &graphql.Field{
				Type: postalType,
			},
			"RegisteredCountry": &graphql.Field{
				Type: countryType,
			},
			"RepresentedCountry": &graphql.Field{
				Type: countryType,
			},
			"Subdivisions": &graphql.Field{
				Type: graphql.NewList(subdivisionsType),
			},
			"Traits": &graphql.Field{
				Type: traitsType,
			},
		},
	},
)

var genterpriseField = &graphql.Field{
	Type:        genterpriseType,
	Description: "Get Enterprise",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetEnterprise,
}

func GetEnterprise(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data, err := db.Enterprise(ip)
	//data = AddCompatibleI18NFields(data)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
