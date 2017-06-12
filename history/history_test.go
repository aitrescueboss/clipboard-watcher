package history

import (
  "testing"
  "math/rand"
  "time"
  "fmt"
)

var mHistoryFilePath01 = "./history_test01.json" //ただのテキスト
var mHistoryFilePath02 = "./history_test02.hist" //改行を含むテキスト

// 単純なインスタンス生成..
func Test_New(t *testing.T) {
  
  tHistory := New()
  if tHistory == nil {
    t.Fatal("Test failed.")
  }
}

// ファイルからの読み込み.
func Test_Import(t *testing.T) {
  tHistory, tError := New().ImportFromFile(mHistoryFilePath01)
  if tError != nil {
    t.Fatalf("Test_Import() failed. error = %v", tError)
  } else {
    fmt.Println("tHistory = ", tHistory)
  }
}

// 先頭保持.
func Test_FrontElementIsValid(t *testing.T) {
  tHistory, tError := New().ImportFromFile(mHistoryFilePath01)
  if tError != nil {
    t.Fatalf("Test_FrontElementIsValid() failed. error = %v", tError)
  }
  
  tActualFronText := "hgoe"
  // tFront := tHistory.FrontString()
  // tFront := tHistory.FrontBytes()
  tFront, tIsValid := tHistory.Front()
  
  if tIsValid {
    tFrontText := tFront.(string)
    if tFrontText != tActualFronText {
      t.Fatalf("Test_FronElementIsValid() failed. Actual front text = %s \n", tFrontText)
    }
  } else {
    t.Fatal("Test_FrontElementIsValid() failed. Front element is not valid. why?")
  }
  
  tFrontText := tHistory.FrontText()
  if tFrontText != tActualFronText {
    t.Fatalf("Test_FronElementIsValid() failed. Actual front text = %s \n", tFrontText)
  }
}

// 途中保持.
//func Test_InterElementIsValid(t *testing.T) {
//  tHistory, tError := New().ImportFromFile(mHistoryFilePath01)
//  if tError != nil {
//    t.Fatalf("Test_InterElementIsValid() failed. error = %v", tError)
//  }
//
//  tIndex := randomInt(tHistory.Len()-1)
//  tInterElement, tError := tHistory.Element(tIndex)
//  if tError != nil {
//    t.Fatalf("Test_InterElementIsValid() failed. error = %v", tError)
//  }
//
//  tMoveUppedElement, tError := tHistory.MoveUp(tIndex)
//  if tError != nil {
//    t.Fatalf("Test_InterElementIsValid() failed. error = %v", tError)
//  }
//
//  if tInterElement != tMoveUppedElement {
//    t.Fatalf("Test_InterElementIsValid() failed. interelement = %v, move-upped element = %v",
//      tInterElement,
//      tMoveUppedElement)
//  }
//}

func randomInt(aMax int) int {
  rand.Seed(time.Now().UnixNano())
  return rand.Intn(aMax)
}

// 改行を含むテキスト.
func Test_TextWithNewLine(t *testing.T) {
  tHistory, tError := New().ImportFromFile(mHistoryFilePath02)
  if tError != nil {
    t.Fatalf("Test_TextWithNewLine() failed. error = %v", tError)
  }
  
  tText := tHistory.FrontText()
  
  tRawText := `
  hoge-line1
  hoge-line2
  `
  
  if tText != tRawText {
    t.Fatalf("Test_TextWithNewLine() failed. text does not match. \n" +
    "tText = %s, tRawText = %s \n",
    tText,
    tRawText)
  }
}