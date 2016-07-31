package main

import "testing"
import "fmt"
import "net/http"
import "net/http/httptest"
import "log"
import "io/ioutil"
import "github.com/willf/bloom"

var testRec = Record{
	UID: "testUID",
	MID: "testMID",
	IP:  "1.1.1.1",
}

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
	if err != nil {
		log.Fatal(err)
	}

	err = res.Body.Close()
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

func BenchmarkBloomAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var filter = bloom.New(1000000*n, 5) // load of 20, 5 keys
		filter.Add([]byte("exists"))
	}
}

func BenchmarkBloomTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var filter = bloom.New(1000000*n, 5) // load of 20, 5 keys
		filter.Test([]byte("shouldnotexist"))
	}
}

func BenchmarkWriteRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(testRec)
	}
}

func BenchmarkReadRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		check(testRec)
	}
}
