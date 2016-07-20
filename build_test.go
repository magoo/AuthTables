package main

import "testing"
import "fmt"
import "net/http"
import "net/http/httptest"
import "log"
import "io/ioutil"
import "github.com/willf/bloom"

func TestPrintLine(t *testing.T) {
	// test stuff here...
	fmt.Println("Print line works, so there's that.")
}

func TestLoad(t *testing.T) {
	writeRecord([]byte("asdf"))
}

func TestWWWServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Testing the client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

func TestBloom(t *testing.T) {
	var filter = bloom.New(1000000*n, 5) // load of 20, 5 keys
	if filter.Test([]byte("shouldnotexist")) {
		log.Fatal("Bloom filter detected a string that wasn't in filter")
	}
	filter.Add([]byte("exists"))
	if !filter.Test([]byte("exists")) {
		log.Fatal("Bloom filter could not detect a string that was in filter")
	}

}
