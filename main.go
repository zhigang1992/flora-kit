package main

import (
	"github.com/zhigang1992/flora-kit/flora"
	"flag"
)



func main() {
	var configFile, geoipdb string
	flag.StringVar(&configFile, "s", "flora.default.conf", "specify surge config file")
	flag.StringVar(&geoipdb, "d", "geoip.mmdb", "specify geoip db file")
	flora.Run(configFile, geoipdb)

}
