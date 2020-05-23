package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

var (
	Port            string
	MaxMind_PATH    string
	ASN_MMDB        string
	AIP_MMDB        string
	City_MMDB       string
	CType_MMDB      string
	Country_MMDB    string
	Domain_MMDB     string
	Enterprise_MMDB string
	ISP_MMDB        string
)

type DB struct {
	ASN            func(net.IP) (*geoip2.ASN, error)
	AnonymousIP    func(net.IP) (*geoip2.AnonymousIP, error)
	City           func(net.IP) (*geoip2.City, error)
	ConnectionType func(net.IP) (*geoip2.ConnectionType, error)
	Country        func(net.IP) (*geoip2.Country, error)
	Domain         func(net.IP) (*geoip2.Domain, error)
	Enterprise     func(net.IP) (*geoip2.Enterprise, error)
	ISP            func(net.IP) (*geoip2.ISP, error)
}

var db DB

func init_database() {
	// Set maxmind path
	MaxMind_PATH = os.Getenv("MAXMIND_PATH")
	// Set mmdb path
	ASN_MMDB = MaxMind_PATH + "/GeoLite2-ASN.mmdb"
	AIP_MMDB = MaxMind_PATH + "/GeoIP2-Anonymous-IP.mmdb"
	City_MMDB = MaxMind_PATH + "/GeoIP2-City.mmdb"
	CType_MMDB = MaxMind_PATH + "/GeoIP2-Connection-Type.mmdb"
	Country_MMDB = MaxMind_PATH + "/GeoIP2-Country.mmdb"
	Domain_MMDB = MaxMind_PATH + "/GeoIP2-Domain.mmdb"
	Enterprise_MMDB = MaxMind_PATH + "/GeoIP2-Enterprise.mmdb"
	ISP_MMDB = MaxMind_PATH + "/GeoIP2-ISP.mmdb"

	var err error
	db_asn, err := geoip2.Open(ASN_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", ASN_MMDB)
	}
	db_aip, err := geoip2.Open(AIP_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", AIP_MMDB)
	}
	db_city, err := geoip2.Open(City_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", City_MMDB)
	}
	db_ctype, err := geoip2.Open(CType_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", CType_MMDB)
	}
	db_country, err := geoip2.Open(Country_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", Country_MMDB)
	}
	db_domain, err := geoip2.Open(Domain_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", Domain_MMDB)
	}
	db_enterprise, err := geoip2.Open(Enterprise_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", Enterprise_MMDB)
	}
	db_isp, err := geoip2.Open(ISP_MMDB)
	if err != nil {
		fmt.Println("Not Found: ", ISP_MMDB)
	}
	db = DB{
		ASN:            db_asn.ASN,
		AnonymousIP:    db_aip.AnonymousIP,
		City:           db_city.City,
		ConnectionType: db_ctype.ConnectionType,
		Country:        db_country.Country,
		Domain:         db_domain.Domain,
		Enterprise:     db_enterprise.Enterprise,
		ISP:            db_isp.ISP,
	}
}

func init() {
	init_database()
}

func main() {
	// Set running port
	Port = os.Getenv("GEOIP2GQL_PORT")
	if Port == "" {
		Port = "3000"
	}

	// Set router
	r := gin.Default()

	// GraphQL API
	r.POST("/gql", HDLR())

	// Simple API
	r.GET("/", GetIPAddress)
	r.GET("/json", GetSimpleData)

	// Start
	fmt.Println()
	fmt.Println("Listening...")
	r.Run(":" + Port)
}
