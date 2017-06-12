package history

import (
  "os"
  "fmt"
  "errors"
)

type HistoryElement struct {
  mContents interface{}
}

type History struct {
  mElements []*HistoryElement
}

func New() *History {
  return &History{}
}

func (aHistory *History) ImportFromFile(aPath string) (*History, error) {

  tFileStat, tError := os.Stat(aPath)
  if tError != nil {
    return nil, fmt.Errorf("history.go ImportFromFile(): ファイルパスが異常です. \n error = %v \n",
      tError)
  }
  if tFileStat.IsDir() {
    return nil, errors.New("history.go ImportFromFile(): ファイルパスにディレクトリが指定されました. \n")
  }
  
  tElements, tError := readElements(aPath)
  if tError != nil {
    return nil, fmt.Errorf("history.go ImportFromFile(): ファイルからの読み込み途中でエラーが発生しました. \n" +
      "error = %v \n", tError)
  }
  
  tNewHistory := &History{ tElements }
  return tNewHistory, nil
}

// 履歴の先頭要素と，先頭要素は有効な要素であるかどうかを返す．
// 履歴がnilかまたは空である場合，戦闘要素は有効な要素ではない．
func (aHistory *History) Front() (interface{}, bool) {
  if aHistory.mElements == nil || len(aHistory.mElements) == 0 {
    return nil, false
  }
  return aHistory.mElements[0], true
}

// 履歴の先頭要素をテキストで返す．
// Front()の第2返り値がfalse, または先頭要素がstringにキャストできない場合は
// 空文字を返す．
func (aHistory *History) FrontText() string {
  tFrontElement, tIsValid := aHistory.Front()
  if !tIsValid {
    return ""
  }
  tFrontText, tIsValidCast := tFrontElement.(string)
  if !tIsValidCast {
    return ""
  }
  return tFrontText
}

// 履歴の要素数を返す．
func (aHistory *History) Len() int {
  if aHistory.mElements == nil {
    return 0
  }
  return len(aHistory.mElements)
}

// 添字指定で履歴要素を返す．
func (aHistory *History) Element(aIndex int) (*HistoryElement, error) {
  if aHistory.Len() == 0 {
    return nil, errors.New("Elements is empty or nil.")
  }
  if aIndex < 0 || aIndex > aHistory.Len() - 1 {
    return nil, fmt.Errorf("aIndex out of range. aIndex = %d, Len() = %d \n",
      aIndex,
      aHistory.Len())
  }
  return aHistory.mElements[aIndex], nil
}

// 引数のファイルパスから履歴の要素を全て読み取って返す．
// TODO 画像データとかどうするよ
func readElements(aValidFilePath string) ([]*HistoryElement, error) {
  
  tFile, tError := os.Open(aValidFilePath)
  if tError != nil {
    return nil, fmt.Errorf("history.go readElements(): ファイルを開く時にエラーが発生しました. \n" +
    "error = %v \n", tError)
  }
  defer tFile.Close()
  
  tElements := make([]*HistoryElement, 0, 1024)

  
  //for tScanner.Scan() {
  //  tText := tScanner.Text()
  //  tHistoryElement := HistoryElement{tText}
  //  tElements = append(tElements, &tHistoryElement)
  //}
  
  return tElements, nil
}