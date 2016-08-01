package main

import "testing"
import "fmt"
import "net/http"
import "net/http/httptest"
import "log"
import "io/ioutil"
import "github.com/willf/bloom"
import "bytes"

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

func TestCheckRequest(t *testing.T) {

	var jsonStr = []byte(`{
	"uid": "magoo",
	"ip": "1.1.1.1",
	"mid": "ASDFQWERASDF"
}`)

	req, err := http.NewRequest("POST", "/check", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checkRequest)

	handler.ServeHTTP(rr, req)

	// Correct Response?
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `OK`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAddRequest(t *testing.T) {

	var jsonStr = []byte(`{
	"uid": "magooadd",
	"ip": "1.1.1.1",
	"mid": "ASDFQWERASDF"
}`)

	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addRequest)

	handler.ServeHTTP(rr, req)

	// Correct Response?
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `ADD`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
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
