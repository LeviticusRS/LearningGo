package main

import "testing"

const WebsiteToTest = "https://google.com"

func TestStatus(t *testing.T) {
	result := getHTTPResponse(WebsiteToTest).StatusCode
	expected := 200
	if result != expected {
		t.Errorf("getHTTPResponse().StatusCode test returned an unexpected result: got %v want %v", result, expected)
	}
}
