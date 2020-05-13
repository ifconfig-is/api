package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/oschwald/geoip2-golang"
)

var gcityType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoIP2City",
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

var gcityField = &graphql.Field{
	Type:        gcityType,
	Description: "Get City",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetCity,
}

func GetCity(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data, err := db.City(ip)
	data = AddNames(data, "city").(*geoip2.City)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
