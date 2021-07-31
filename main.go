package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

var (
	PORT         string
	STATIC       string
	NG_URL       string
	PROD         bool
	DBIP_PATH    string
	ASN_MMDB     string
	City_MMDB    string
	Country_MMDB string
)

type DB struct {
	ASN     func(net.IP) (*geoip2.ASN, error)
	City    func(net.IP) (*geoip2.City, error)
	Country func(net.IP) (*geoip2.Country, error)
}

var db DB

func init_database() {
	// Set maxmind path
	DBIP_PATH = os.Getenv("DBIP_PATH")
	if DBIP_PATH == "" {
		DBIP_PATH = "./dbip"
	}
	// Set mmdb path
	ASN_MMDB = DBIP_PATH + "/asn.mmdb"
	City_MMDB = DBIP_PATH + "/city.mmdb"
	Country_MMDB = DBIP_PATH + "/country.mmdb"

	var err error
	db_asn, err := geoip2.Open(ASN_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", ASN_MMDB)
	}
	db_city, err := geoip2.Open(City_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", City_MMDB)
	}
	db_country, err := geoip2.Open(Country_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", Country_MMDB)
	}
	db = DB{
		ASN:     db_asn.ASN,
		City:    db_city.City,
		Country: db_country.Country,
	}
}

func init() {
	init_flag()
	init_database()
}

func main() {
	// Set global variables
	PORT = os.Getenv("IFCONFIGIS_API_PORT")
	if PORT == "" {
		PORT = "5000"
	}

	STATIC = os.Getenv("IFCONFIGIS_STATIC")
	if STATIC == "" {
		STATIC = "./static"
	}

	NG_URL = os.Getenv("IFCONFIGIS_NG_URL")
	if NG_URL == "" {
		NG_URL = "http://127.0.0.1:5080"
	}

	// Set router
	r := gin.Default()

	// Allow all origins
	r.Use(cors.Default())

	// Add dispatcher
	r.Use(Dispatcher())

	// Start
	fmt.Println()
	fmt.Println("Listening...")
	r.Run(":" + PORT)
}
