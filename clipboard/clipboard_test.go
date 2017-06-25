package clipboard

import (
  "testing"
)

var mExpectedClipboardText string = "ainomamaniwagamamanibokuhakimidakewokizutukenai"

func TestGetClipBoard(t *testing.T) {

  tActualClipBoardText, tError := GetClipBoard()
  if tError != nil {
    t.Fatalf("TestGetClipBoard(): %v", tError)
  }
  if string(tActualClipBoardText) != mExpectedClipboardText {
    t.Fatalf("TestGetClipBoard(): actual = %s", tActualClipBoardText)
  }
}

var mWillBePastedToClipBoard string = "oosamumachiniromanhashizumu"
func TestPasteToClipBoard(t *testing.T) {
  tError := PasteToClipBoard([]byte(mWillBePastedToClipBoard))
  if tError != nil {
    t.Fatalf("TestPasteToClipBoard(): %v", tError)
  }
  
  tActualClipBoard, tError := GetClipBoard()
  if tError != nil {
    t.Fatalf("TestPasteClipBoard(): actual = %s, error = %v", tActualClipBoard, tError)
  }
}