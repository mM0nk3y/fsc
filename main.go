package main

import (
	"test/config"
	"test/scan"
)

func main() {
	config.API_init()
	scan.FofaGet()
	scan.ShodanGet()
}
