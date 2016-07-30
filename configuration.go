package main

import (
	"flag"
	"os"
)

//JSON Config Struct
type Configuration struct {
	Host     string
	Port     string
	Password string
	Loglevel string
}

//Global access to configuration variables
var c Configuration = readConfig()

func readConfig() (c Configuration) {
	//Command Line Flags. If command line flag is blank, use ENV instead
	var flag_host string
	flag.StringVar(&flag_host, "host", os.Getenv("AUTHTABLES_HOST"), "hostname for redis")
	var flag_port string
	flag.StringVar(&flag_port, "port", os.Getenv("AUTHTABLES_PORT"), "port for redis")
	var flag_pw string
	flag.StringVar(&flag_pw, "password", os.Getenv("AUTHTABLES_PW"), "password for redis")
	var flag_loglevel string
	flag.StringVar(&flag_loglevel, "loglevel", os.Getenv("AUTHTABLES_LOGLEVEL"), "level of logging (debug, info, warn, error)")
	flag.Parse()

	//We're going to load this with config data. See struct!
	configuration := Configuration{}

	configuration.Host = flag_host
	configuration.Port = flag_port
	configuration.Password = flag_pw
	configuration.Loglevel = flag_loglevel

	return configuration
}
