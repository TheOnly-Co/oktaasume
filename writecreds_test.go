package main 

import (
    "io/ioutil"    
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

type authTest struct {
    str string
    factor string
}
var authTests = []authTest {
    authTest {
     str: "hi",
     factor: "hi",
    },
    authTest {
     str: "hi",
     factor: "hi",
    },
}


func TestSearchAuthMethod (t *testing.T) {

    mockData ,err := searchAuthMethod(authTests,"hi")
    if err != nil {
    t.Error(err.Error())
    } 
    for _, i := range authTests{
        if i.fator != "hi" {
            t.Error("values do not match")
        }
    }    
     
}



























