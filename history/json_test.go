package history

import (
  "testing"
  "fmt"
)

var mJsonPath01 = "./history_test01.json"
func TestMJsonContents_ImportFromJsonPath(t *testing.T) {
  tJson, tError := newJson().ImportFromJsonPath(mJsonPath01)
  
  if tError != nil {
    t.Fatalf("TestMJsonContents_ImportFromJsonPath() failed. error = %v \n", tError)
  }
  fmt.Println("tJson = ", tJson)
}
