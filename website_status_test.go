package main

import "testing"

const WebsiteToTest = "https://google.com"

func TestStatus(t *testing.T) {
	result := getHttpResponse(WebsiteToTest).StatusCode
	expected := 200
	if result != expected {
		t.Errorf("getHttpResponse().StatusCode test returned an unexpected result: got %v want %v", result, expected)
	}
}
