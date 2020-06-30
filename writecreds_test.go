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

type AddRes struct {
    x int
    y int
    expected int
}

var AddRess = []AddRes{
    {1,1,2},
}

func TestSum(t *testing.T) {
    for _, i := range AddRess{
    res := Add(i.x,i.y)

    if res != i.expected {
        t.Fatal("values do not match")
    }
  }
}

























