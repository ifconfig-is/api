## Intro

**ifconfig.is** is an IP lookup online service. Live demo: [https://ifconfig.is](https://ifconfig.is)

The IP data source is [https://db-ip.com/](https://db-ip.com/)

## Command Line Usage

```
$ curl ifconfig.is
1.2.3.4

$ curl ifconfig.is/json
{
  "ip": "1.2.3.4",
  "continent": "Asia",
  "country": "Singapore",
  "city": "Singapore",
  "latitude": 1.35208,
  "longitude": 103.82,
  "asn": 14061,
  "organization": "DigitalOcean, LLC"
}

$ curl ifconfig.is/json/www.google.com
{
  "ip": "74.125.68.101",
  "continent": "Asia",
  "country": "Singapore",
  "city": "Singapore (Queenstown Estate)",
  "latitude": 1.27623,
  "longitude": 103.8,
  "asn": 15169,
  "organization": "Google LLC"
}
```

## License

See the [LICENSE](https://github.com/i3h/ifconfig/blob/master/LICENSE.md) file for license rights and limitations (MIT).

# Acknowledgements

[mpolden/echoip](https://github.com/mpolden/echoip)

[oschwald/geoip2-golang](https://github.com/oschwald/geoip2-golang)

[DB-IP](https://db-ip.com)
