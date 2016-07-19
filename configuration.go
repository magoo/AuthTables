package main

import (
      "fmt"
      "encoding/json"
      "os"
      "flag"
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
  flag.Parse()
  //Command Line Flags
  var config_file string
  flag.StringVar(&config_file, "conf", "./conf.json", "path to config file")

  //May need to come from CLI, built in for now
  file, _ := os.Open(config_file)
  decoder := json.NewDecoder(file)
  configuration := Configuration{}
  err := decoder.Decode(&configuration)
  if err != nil {
    fmt.Println("error:", err)
  }
  return configuration
}
