package main

import (
	"fmt"
	"net"
)

type Data struct {
	Continent string  `json:"Continent"`
	Country   string  `json:"Country"`
	City      string  `json:"City"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	TimeZone  string  `json:"TimeZone"`
	IsEU      bool    `json:"IsEU"`
	ASN       uint    `json:"ASN"`
	ORG       string  `json:"ORG"`
}

func GetData(ip net.IP) Data {
	asn, err := db.ASN(ip)
	if err != nil {
		fmt.Println(err)
	}
	city, err := db.City(ip)
	if err != nil {
		fmt.Println(err)
	}

	data := Data{
		Continent: city.Continent.Names["en"],
		Country:   city.Country.Names["en"],
		City:      city.City.Names["en"],
		Latitude:  city.Location.Latitude,
		Longitude: city.Location.Longitude,
		TimeZone:  city.Location.TimeZone,
		IsEU:      city.Country.IsInEuropeanUnion,
		ASN:       asn.AutonomousSystemNumber,
		ORG:       asn.AutonomousSystemOrganization,
	}

	return data
}
