package main

import (
	"net"
)

type Data struct {
	IP        string  `json:"ip"`
	Continent string  `json:"continent"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Lat       float64 `json:"latitude"`
	Lon       float64 `json:"longitude"`
	ASN       uint    `json:"asn"`
	ORG       string  `json:"organization"`
}

func GetData(ip net.IP) Data {
	asn, err := db.ASN(ip)
	if err != nil {
		//fmt.Println(err)
	}
	city, err := db.City(ip)
	if err != nil {
		//fmt.Println(err)
	}

	data := Data{
		IP:        ip.String(),
		Continent: city.Continent.Names["en"],
		Country:   city.Country.Names["en"],
		City:      city.City.Names["en"],
		Lat:       city.Location.Latitude,
		Lon:       city.Location.Longitude,
		ASN:       asn.AutonomousSystemNumber,
		ORG:       asn.AutonomousSystemOrganization,
	}

	return data
}
