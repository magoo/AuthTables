package main

import (
	"github.com/willf/bloom"
	"gopkg.in/redis.v4"
)

//Bloom Filter
var n uint = 1000
var filter = bloom.New(1000000*n, 5) // load of 20, 5 keys

type Record struct {
	UID string `json:"uid"`
	IP  string `json:"ip"`
	MID string `json:"mid"`
}

//Main data structure for Bloom and Redis.
type RecordHashes struct {
	uid     []byte
	uid_mid []byte
	uid_ip  []byte
	uid_all []byte
	ip_mid  []byte
	mid_ip  []byte
}

//Take us online to Redis
var client = redis.NewClient(&redis.Options{
	Addr:     c.Host + ":" + c.Port,
	Password: c.Password, // no password set
	DB:       0,          // use default DB
})
