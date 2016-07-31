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

var filterTest = bloom.New(1000000*n, 5)

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
	if filterTest.Test([]byte("shouldnotexist")) {
		log.Fatal("Bloom filter detected a string that wasn't in filter")
	}
	filterTest.Add([]byte("exists"))
	if !filterTest.Test([]byte("exists")) {
		log.Fatal("Bloom filter could not detect a string that was in filter")
	}

}

func BenchmarkBloomAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// We reset the timer because the bloom filter is only created on boot.
		//b.ResetTimer()
		filterTest.Add([]byte("exists"))
	}
}

func BenchmarkBloomTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// We reset the timer because the bloom filter is only created on boot.
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
