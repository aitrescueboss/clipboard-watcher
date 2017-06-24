package history

import (
  "testing"
  "fmt"
)

var mHistoryFilePath01 = "./history_test01.json" //ただのテキスト
var mHistoryFilePath02 = "./history_test02.json" //改行を含むテキスト

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
  
  tExpectedFrontText := "hoge"
  // tFront := tHistory.FrontString()
  // tFront := tHistory.FrontBytes()
  tFront, tIsValid := tHistory.Front()
  
  if tIsValid {
    if tFront != tExpectedFrontText {
      t.Fatalf("Test_FronElementIsValid() failed. Actual front text = %s \n", tFront)
    }
  } else {
    t.Fatal("Test_FrontElementIsValid() failed. Front element is not valid. why?: ", tFront)
  }
}

// 途中保持.
func Test_InterElementIsValid(t *testing.T) {
  tHistory, tError := New().ImportFromFile(mHistoryFilePath01)
  if tError != nil {
    t.Fatalf("Test_InterElementIsValid() failed. error = %v", tError)
  }

  tIndex := 2 // "piyo" exptected.
  _, tError = tHistory.Element(tIndex)
  if tError != nil {
    t.Fatalf("Test_InterElementIsValid() failed. error = %v", tError)
  }

  tHistory, tError = tHistory.MoveUp(tIndex)
  if tError != nil {
    t.Fatalf("Test_InterElementIsValid() failed. error = %v", tError)
  }
  
  tFront, tIsValid := tHistory.Front()
  if tIsValid != true || tFront != "piyo" {
    t.Fatalf("Test_InterElementIsValid() failed. exptected = %s, actual = %s", "piyo", tFront)
  }
  
  fmt.Print("MoveUp()後: ")
  for _, tElement := range tHistory.mElements {
    fmt.Print(tElement.mContents, ", ")
  }
  fmt.Println()
}

// 改行を含むテキスト.
func Test_TextWithNewLine(t *testing.T) {
  tHistory, tError := New().ImportFromFile(mHistoryFilePath02)
  if tError != nil {
    t.Fatalf("Test_TextWithNewLine() failed. error = %v", tError)
  }

  tText, _ := tHistory.Front()

  tRawText := `hoge-line1
hoge-line2`

  if tText != tRawText {
    t.Fatalf("Test_TextWithNewLine() failed. text does not match. \n" +
    "tText = %s \n, tRawText = %s \n",
    tText,
    tRawText)
  }
}

// ファイル出力.
func TestHistory_ExportToFile(t *testing.T) {
  tHistory, tError := New().ImportFromFile(mHistoryFilePath02)
  if tError != nil {
    t.Fatalf("TestHistory_ExportToFile() failed. error = %v", tError)
  }
  
  tHistory, tError = tHistory.ExportToFile("./history_test03.json")
  if tError != nil {
    t.Fatalf("TestHistory_ExportToFile() failed. error = %v", tError)
  }
}