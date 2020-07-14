package main

import (

    "io/ioutil"    
    "github.com/hunkeelin/oktalib"
    "testing"
)

func TestWriteCreds (t *testing.T) {
    mockWrite := []byte("hello")
    err := writeCredToFile(&writeCredToFileInput{
        location: "/tmp/hello",
        towrite: mockWrite,
    })
    if err != nil {
        t.Error(err.Error())
    }
    f, err := ioutil.ReadFile("/tmp/hello")
    if err != nil{
        t.Error(err.Error())
    }
    if string(f) != string(mockWrite) {
        t.Errorf("The input is different than the output")
    }
}

func TestSearchAuthMethod(t *testing.T) {
    var mockFactors []oktalib.OktaUserAuthnFactor
    mockFactors = append(mockFactors, oktalib.OktaUserAuthnFactor{
        FactorType: "foo",
    })
    if searchAuthMethod(mockFactors, "bar") {
        t.Errorf("searchAuthMethod return true when two values are not the same")
    }
    if !searchAuthMethod(mockFactors, "foo") {
        t.Errorf("searchAuthMethod return false when the two values are the same")
    }
}

























