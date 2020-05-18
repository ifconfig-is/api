package main

import (
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"IP":                   ipField,
			"GeoIP2ASN":            gasnField,
			"GeoIP2AnonymousIP":    gaipField,
			"GeoIP2City":           gcityField,
			"GeoIP2ConnectionType": gctypeField,
			"GeoIP2Country":        gcountryField,
			"GeoIP2Domain":         gdomainField,
			"GeoIP2Enterprise":     genterpriseField,
			"GeoIP2ISP":            gispField,
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
