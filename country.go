package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/oschwald/geoip2-golang"
)

var gcountryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoIP2Country",
		Fields: graphql.Fields{
			"Continent": &graphql.Field{
				Type: continentType,
			},
			"Country": &graphql.Field{
				Type: countryType,
			},
			"RegisteredCountry": &graphql.Field{
				Type: countryType,
			},
			"RepresentedCountry": &graphql.Field{
				Type: countryType,
			},
			"Traits": &graphql.Field{
				Type: traitsType,
			},
		},
	},
)

var gcountryField = &graphql.Field{
	Type:        gcountryType,
	Description: "Get Country",
	Args: graphql.FieldConfigArgument{
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: GetCountry,
}

func GetCountry(p graphql.ResolveParams) (interface{}, error) {
	ip := GetIP(p)
	data, err := db.Country(ip)
	data = AddNames(data, "country").(*geoip2.Country)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
