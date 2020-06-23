package main

import (
	"io/ioutil"
	"testing"
)

func TestWriteCreds(t *testing.T) {
	mockWrite := []byte("foo")
	err := setCredentials(&setCredentialsInput{
		toWrite:  mockWrite,
		location: "/tmp/foo",
	})
	if err != nil {
		t.Error(err.Error())
	}
	f, err := ioutil.ReadFile("/tmp/foo")
	if err != nil {
		t.Error(err.Error())
	}
	if string(f) != string(mockWrite) {
		t.Errorf("The input is different than the output")
	}
}
