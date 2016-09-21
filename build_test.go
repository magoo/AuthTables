package main

import (
				"testing"
 				"fmt"
 				"net/http"
 				"net/http/httptest"
 				"log"
 				"io/ioutil"
 				"github.com/willf/bloom"
 				"bytes"
 				"encoding/json"
			)

var testRec = Record{
	Uid: "testUID",
	Mid: "testMID",
	Ip:  "1.1.1.1",
}

var filterTest = bloom.NewWithEstimates(c.BloomSize, 1e-3) // Configurable in environment var.

func TestRedisConnectivity (t *testing.T) {

	_, err := client.Ping().Result()
	if err != nil {
		t.Errorf("We can't ping redis.")
	}
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
	if filterTest.Test([]byte("shouldnotexist")) {
		log.Fatal("Bloom filter detected a string that wasn't in filter")
	}
	filterTest.Add([]byte("exists"))
	if !filterTest.Test([]byte("exists")) {
		log.Fatal("Bloom filter could not detect a string that was in filter")
	}

}

func TestCheckRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/check", bytes.NewBuffer(testRec.Marshaler()))
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

func TestResetRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/reset", bytes.NewBuffer(testRec.Marshaler()))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(resetRequest)

	handler.ServeHTTP(rr, req)

	// Correct Response?
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `RESET`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAddRequest(t *testing.T) {

	jsonStr, err := json.Marshal(testRec)
	if err != nil {
		t.Fatal(err)
	}

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
	// Check a key was written
	if !(canGetKey("testUID:1.1.1.1") && canGetKey("testUID:testMID")) {
		t.Errorf("keys not being written")
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
