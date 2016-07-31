package main

import (
	"flag"
	"os"
)

//Configuration holds data cleaned from our ENV variables or passed through cmd line
type Configuration struct {
	Host     string
	Port     string
	Password string
	Loglevel string
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
	flag.Parse()

	//We're going to load this with config data. See struct!
	configuration := Configuration{}

	configuration.Host = flagHost
	configuration.Port = flagPort
	configuration.Password = flagPW
	configuration.Loglevel = flagLoglevel

	return configuration
}
