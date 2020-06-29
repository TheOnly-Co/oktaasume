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
type okta struct {
    factor string
}
type authTest struct {
    str string
    Objects []okta    
}

var authTests = []authTest {
    authTest {
     str: "hi",
     Objects: []okta {
            okta {
                factor:"hi",
            },
        },
    },
}

func TestSearchAuthMethod (t *testing.T) {

//         prototype1
//         mockData ,err := searchAuthMethod([]authTests,  authTest.str)
//       if err != nil {
//        t.Error(err.Error())
//        }
//        for _, i := range authTests{
//           if i.fator != "hi" {
//                t.Error("values do not match")
//            }
//        }
//  prototpye2
//       for _, i := range authTests {
//           mockData, err := searchAuthMethod(i.Objects,i.str)
//           if err != nil {
//               return err
//           }
//           if i.Objects.factor != i.str {
//               t.Error("Values do not match")
//           }
//       }        
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

























