package clipboard

import (
  "os/exec"
  "fmt"
  "os"
  "runtime"
)

func GetClipBoard() ([]byte, error) {
  
  tCommandLiteral := ""
  switch runtime.GOOS {
  case "darwin":
    tCommandLiteral = "/usr/bin/pbpaste"
  case "linux":
    tCommandLiteral = "xsel --clipboard --output"
  }
  
  tPasteCommand := exec.Command(tCommandLiteral)
  tText, tError := tPasteCommand.Output()
  
  if tError != nil {
    return nil, fmt.Errorf("GetClipBoard(): コマンド実行でエラー %v", tError)
  }
  return tText, nil
}

func PasteToClipBoard(aRawBytes []byte) error {
  tTempFilePath := "/tmp/clip.txt"
  tTempFile, tError := os.Create(tTempFilePath)
  
  if tError != nil {
    return fmt.Errorf("PasteToClipBoard(): TempFile作成できず")
  }
  
  // 引数のバイト列を一時ファイルに書き込む→中身を後でクリップボードに貼り付ける
  tTempFile.Write(aRawBytes)
  tTempFile.Sync()
  tTempFile.Close()
  
  tCommandLiteral := ""
  switch runtime.GOOS {
  case "darwin":
    tCommandLiteral = "/usr/bin/pbcopy"
  case "linux":
    tCommandLiteral = "xsel --clipboard --input"
  }
  
  // このあたりを参考にするとパイプ経由でコマンドの実行結果を得られる
  // http://qiita.com/yuroyoro/items/9358cd25b5f7fe9dd37f
  tCopyCommand := exec.Command(
    "/bin/bash",
    "-c",
    fmt.Sprintf("/bin/cat \"%s\" | %s", tTempFilePath, tCommandLiteral),
  )
  _, tError = tCopyCommand.Output()
  if tError != nil {
    return fmt.Errorf("PasteToClipBoard(): コマンド実行でエラー %v", tError)
  }
  
  os.Remove(tTempFilePath)
  return nil
}
