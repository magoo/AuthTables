package main

import (
      "os"
      "flag"
      "fmt"
    )


//JSON Config Struct
type Configuration struct {
  Host      string
  Port      string
  Password  string
}

//Global access to configuration variables
var c Configuration = readConfig()

func readConfig() (c Configuration) {
  fmt.Println("Reading config!")
  //Command Line Flags
  var flag_host string
  flag.StringVar(&flag_host, "host", os.Getenv("AUTHTABLES_HOST"), "hostname for redis")
  var flag_port string
  flag.StringVar(&flag_port, "port", os.Getenv("AUTHTABLES_PORT"), "port for redis")
  var flag_pw string
  flag.StringVar(&flag_pw, "password", os.Getenv("AUTHTABLES_PW"), "password for redis")
  flag.Parse()

  fmt.Println("Reading config: " + flag_host + flag_port + flag_pw)
  //We're going to load this with config data. See struct!
  configuration := Configuration{}

  //command line flag is blank, use ENV instead

  configuration.Host = os.Getenv("AUTHTABLES_HOST")
  configuration.Port = os.Getenv("AUTHTABLES_PORT")
  configuration.Password = os.Getenv("AUTHTABLES_PW")

  //configuration.Host = flag_host
  //configuration.Port = flag_port
  //configuration.Password = flag_pw

  return configuration
}
