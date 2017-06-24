package history

import (
  "os"
  "fmt"
  "errors"
  "strings"
  "encoding/json"
  "bytes"
)

type HistoryElement struct {
  mContents string
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

func (aHistory *History) ExportToFile(aPath string) (*History, error) {
  if aHistory.mElements == nil || len(aHistory.mElements) == 0 {
    return aHistory, nil
  }
  tError := writeElements(aPath, aHistory.mElements)
  if tError != nil {
    return aHistory, fmt.Errorf("ExportToFile(): error in writeElements(), error = %v", tError)
  }
  return aHistory, nil
}

// 履歴の先頭テキストと，先頭要素は有効な要素であるかどうかを返す．
// 履歴がnilかまたは空である場合，戦闘要素は有効な要素ではない．
func (aHistory *History) Front() (string, bool) {
  if aHistory.mElements == nil || len(aHistory.mElements) == 0 {
    return "", false
  }
  return aHistory.mElements[0].mContents, true
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

// 履歴の aIndex 番目の要素を先頭に配置換えする.
// 履歴が1つも含まれていない, または aIndex が履歴の範囲外を指している場合
// 何もしない.
func (aHistory *History) MoveUp(aIndex int) (*History, error) {
  if aHistory.Len() == 0 {
    return aHistory, nil
  }
  if aIndex < 0 || aIndex > aHistory.Len() - 1 {
    return aHistory, nil
  }
  tTargetElement := aHistory.mElements[aIndex]
  
  tNewElements := make([]*HistoryElement, 0, aHistory.Len())
  tNewElements = append(tNewElements, tTargetElement)
  for tCurrrentIndex, tElement := range aHistory.mElements {
    if tCurrrentIndex == aIndex {
      continue
    }
    tNewElements = append(tNewElements, tElement)
  }
  aHistory.mElements = tNewElements
  return aHistory, nil
}

// 引数のファイルパスから履歴の要素を全て読み取って返す．
// TODO 画像データとかどうするよ→めんどいのでパス
func readElements(aValidFilePath string) ([]*HistoryElement, error) {
  
  tFile, tError := os.Open(aValidFilePath)
  if tError != nil {
    return nil, fmt.Errorf("history.go readElements(): ファイルを開く時にエラーが発生しました. \n" +
    "error = %v \n", tError)
  }
  defer tFile.Close()
  
  tElements := make([]*HistoryElement, 0, 1024)

  // JSONパースして渡す
  tJsonDecoder := newJson()
  tJsonContents, tError := tJsonDecoder.ImportFromJsonPath(aValidFilePath)
  if tError != nil {
    return nil, fmt.Errorf("history.go readElements(): JSONのパースでエラーが発生しました. \n" +
    "error = %v \n", tError)
  }
  
  for _, tJsonContent := range tJsonContents.Contents {
    // 複数行あるならばつなげる
    tClipboardText := ""
    tLineNum := 0
    for _, tContentText := range tJsonContent.Content {
      if tLineNum > 0 {
        tClipboardText += "\n"
      }
      tClipboardText += tContentText
      tLineNum += 1
    }
    
    tElements = append(tElements, &HistoryElement{
      mContents:tClipboardText,
    })
  }
  
  return tElements, nil
}

func writeElements(aFilePath string, aElements []*HistoryElement) error {
  
  tContents := make([]mJsonContent, 0, 1024)
  for _, tElement := range aElements {
    tRawContent := tElement.mContents
    tSplittedByNewLine := strings.Split(tRawContent, "\n")
    tContent := make([]string, 0, 128)
    for _, tSingleLineContent := range tSplittedByNewLine {
      tContent = append(tContent, tSingleLineContent)
    }
    tContents = append(tContents, mJsonContent{Content:tContent})
  }
  
  tJsonContents := mJsonContents{
    Contents: tContents,
  }
  
  tMarshalled, tError := json.Marshal(tJsonContents)
  if tError != nil {
    return fmt.Errorf("writeElements(): JSONをマーシャルできず. error = %v", tError)
  }
  tIndentedMarshall, tError := func(aRawBytes []byte) ([]byte, error) {
    var tBuffer bytes.Buffer
    tErr := json.Indent(&tBuffer, aRawBytes, "", "  ")
    return tBuffer.Bytes(), tErr
  }(tMarshalled)
  
  tOutFile, tError := os.Create(aFilePath)
  if tError != nil {
    return fmt.Errorf("writeElements(): ファイルを開けず. path = %s, error = %v", aFilePath, tError)
  }
  
  _, tError = tOutFile.Write(tIndentedMarshall)
  if tError != nil {
    return fmt.Errorf("writeElements(): ファイルに書けず. error = %v", tError)
  }
  
  tOutFile.Sync()
  tOutFile.Close()
  
  return nil
}