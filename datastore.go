package main

import (
	"github.com/willf/bloom"
	"gopkg.in/redis.v4"
)

//Bloom Filter
var n uint = 1000
var filter = bloom.New(1000000*n, 5) // load of 20, 5 keys

//Record is the main struct that is passed from applications to AuthTables as JSON.
//Applications send us these, and AuthTables responds with `OK`s or `BAD`
type Record struct {
	UID string `json:"uid"`
	IP  string `json:"ip"`
	MID string `json:"mid"`
}

//RecordHashes is a struct ready for use in the bloom filter or redis.
type RecordHashes struct {
	uid    []byte
	uidMID []byte
	uidIP  []byte
	uidALL []byte
	ipMID  []byte
	midIP  []byte
}

//Take us online to Redis
var client = redis.NewClient(&redis.Options{
	Addr:     c.Host + ":" + c.Port,
	Password: c.Password, // no password set
	DB:       0,          // use default DB
})
