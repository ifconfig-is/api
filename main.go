package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

const (
	Port            string = "3000"
	ASN_MMDB        string = "maxmind/GeoLite2-ASN-Test.mmdb"
	AIP_MMDB        string = "maxmind/GeoIP2-Anonymous-IP-Test.mmdb"
	City_MMDB       string = "maxmind/GeoIP2-City-Test.mmdb"
	CType_MMDB      string = "maxmind/GeoIP2-Connection-Type-Test.mmdb"
	Country_MMDB    string = "maxmind/GeoIP2-Country-Test.mmdb"
	Domain_MMDB     string = "maxmind/GeoIP2-Domain-Test.mmdb"
	Enterprise_MMDB string = "maxmind/GeoIP2-Enterprise-Test.mmdb"
	ISP_MMDB        string = "maxmind/GeoIP2-ISP-Test.mmdb"
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
	var err error
	db_asn, err := geoip2.Open(ASN_MMDB)
	if err != nil {
		log.Fatal(err)
	}
	db_aip, err := geoip2.Open(AIP_MMDB)
	if err != nil {
		log.Fatal(err)
	}
	db_city, err := geoip2.Open(City_MMDB)
	if err != nil {
		log.Fatal(err)
	}
	db_ctype, err := geoip2.Open(CType_MMDB)
	if err != nil {
		log.Fatal(err)
	}
	db_country, err := geoip2.Open(Country_MMDB)
	if err != nil {
		log.Fatal(err)
	}
	db_domain, err := geoip2.Open(Domain_MMDB)
	if err != nil {
		log.Fatal(err)
	}
	db_enterprise, err := geoip2.Open(Enterprise_MMDB)
	if err != nil {
		log.Fatal(err)
	}
	db_isp, err := geoip2.Open(ISP_MMDB)
	if err != nil {
		log.Fatal(err)
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
	// Set router
	r := gin.Default()

	// GraphQL API Endpoint
	r.POST("/gql", HDLR())

	// Start
	fmt.Println()
	fmt.Println("Listening...")
	r.Run(":" + Port)
}
