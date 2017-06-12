package history

import (
  "os"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "bytes"
  "github.com/pkg/errors"
)

type mJsonContents struct {
  Contents []struct {
    Content []string `json:"content"`
  } `json:"contents"`
}

func newJson() *mJsonContents {
  return &mJsonContents{}
}

func (aContents *mJsonContents) ImportFromJsonPath(aPath string) (*mJsonContents, error) {
  tFileStat, tError := os.Stat(aPath)
  if tError != nil || tFileStat.IsDir() {
    tNewError := tError
    if tError == nil {
      tNewError = fmt.Errorf("指定されたパスはディレクトリです. path = %s \n", aPath)
    }
    return nil, errors.Wrap(tNewError, "指定されたパスを開く時にエラーが発生しました. \n")
  }
  
  tRawBytes, tError := ioutil.ReadFile(aPath)
  if tError != nil {
    return nil, errors.Wrap(tError,"json.go ImportFromJsonPath(); ファイルを読めません. \n")
  }
  var newJsonContents mJsonContents
  tError = json.NewDecoder(bytes.NewReader(tRawBytes)).Decode(&newJsonContents)
  if tError != nil {
    return nil, errors.Wrap(tError, "JSONのパースでエラーが発生しました. ")
  }
  
  return &newJsonContents, nil
}