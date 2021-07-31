package main

import (
	"embed"
	"fmt"
	"net"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

//go:embed dbip static
var f embed.FS

var (
	PORT   string
	STATIC string
	NG_URL string
	DEV    bool
)

type DB struct {
	ASN     func(net.IP) (*geoip2.ASN, error)
	City    func(net.IP) (*geoip2.City, error)
	Country func(net.IP) (*geoip2.Country, error)
}

var db DB

func init_database() {
	bytes_asn, err := f.ReadFile("dbip/asn.mmdb")
	db_asn, err := geoip2.FromBytes(bytes_asn)
	if err != nil {
		fmt.Println(err)
	}
	bytes_city, err := f.ReadFile("dbip/city.mmdb")
	db_city, err := geoip2.FromBytes(bytes_city)
	if err != nil {
		fmt.Println(err)
	}
	bytes_country, err := f.ReadFile("dbip/country.mmdb")
	db_country, err := geoip2.FromBytes(bytes_country)
	if err != nil {
		fmt.Println(err)
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
