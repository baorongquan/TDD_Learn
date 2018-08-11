package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteCheck(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://www.google.com",
		"http://blog.github.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteCheck, websites)

	want := len(websites)
	got := len(actualResults)
	if got != want {
		t.Fatalf("want %v got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://www.google.com":   true,
		"http://blog.github.com":  true,
		"waat://furhurterwe.geds": false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("expect %v actual %v", expectedResults, actualResults)
	}
}
