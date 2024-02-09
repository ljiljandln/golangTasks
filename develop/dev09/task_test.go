package main

import "testing"

func TestEmptyUrl(t *testing.T) {
	err := wget("")
	if err == nil {
		t.Errorf("The test had to return an error, but hasn't")
	}
}

func TestNotValidUrl(t *testing.T) {
	err := wget("uuu")
	if err == nil {
		t.Errorf("The test had to return an error, but hasn't")
	}
}

func TestOk(t *testing.T) {
	err := wget("http://kermlinrussia.com")
	if err != nil {
		t.Errorf("The test shouldn't return an error")
	}
}
