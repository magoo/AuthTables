package main

import (
	"flag"
	"os"
 	"strconv"
)

//Configuration holds data cleaned from our ENV variables or passed through cmd line
type Configuration struct {
	Host     string
	Port     string
	Password string
	Loglevel string
	BloomSize uint
}

//Global access to configuration variables
var c = readConfig()

func readConfig() (c Configuration) {
	//Command Line Flags. If command line flag is blank, use ENV instead
	var flagHost string
	flag.StringVar(&flagHost, "host", os.Getenv("AUTHTABLES_HOST"), "hostname for redis")
	var flagPort string
	flag.StringVar(&flagPort, "port", os.Getenv("AUTHTABLES_PORT"), "port for redis")
	var flagPW string
	flag.StringVar(&flagPW, "password", os.Getenv("AUTHTABLES_PW"), "password for redis")
	var flagLoglevel string
	flag.StringVar(&flagLoglevel, "loglevel", os.Getenv("AUTHTABLES_LOGLEVEL"), "level of logging (debug, info, warn, error)")
	var flagBloomSize uint
	d, _  := strconv.ParseUint(os.Getenv("AUTHTABLES_BLOOMSIZE"), 0, 32)
	flag.UintVar(&flagBloomSize, "bloomsize", uint(d), "size of bloom filter (default 1e9)")
	flag.Parse()

	//We're going to load this with config data. See struct!
	configuration := Configuration{}

	configuration.Host = flagHost
	configuration.Port = flagPort
	configuration.Password = flagPW
	configuration.Loglevel = flagLoglevel
	configuration.BloomSize = flagBloomSize

	return configuration
}
