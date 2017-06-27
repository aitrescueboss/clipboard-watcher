package clipboard

import (
  "os/exec"
  "fmt"
  "os"
  "runtime"
)

func GetClipBoard() ([]byte, error) {
  
  tCommandLiteral := ""
  tCommandArgs := make([]string, 0, 8)
  switch runtime.GOOS {
  case "darwin":
    tCommandLiteral = "/usr/bin/pbpaste"
  case "linux":
    tCommandLiteral = "xsel"
    tCommandArgs = append(tCommandArgs, "--clipboard")
    tCommandArgs = append(tCommandArgs, "--output")
  }
  
  tPasteCommand := exec.Command(tCommandLiteral, tCommandArgs...)
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
  
  tCommandLiteral := "/bin/bash"
  tCommandArgs := make([]string, 0, 8)
  tCommandArgs = append(tCommandArgs, "-c")
  switch runtime.GOOS {
  case "darwin":
    tCommandArgs = append(tCommandArgs, fmt.Sprintf("/bin/cat \"%s\" | %s", tTempFilePath, "/usr/bin/pbcopy"))
    //tCommandArgs = append(tCommandArgs, "/usr/bin/pbcopy")
    //tCommandLiteral = "/usr/bin/pbcopy"
  case "linux":
    tCommandLiteral = "xsel"
    tCommandArgs = append(tCommandArgs, fmt.Sprintf("/bin/cat \"%s\" | %s %s %s", tTempFilePath, "xsel", "--clipboard", "--input"))
    //tCommandArgs = append(tCommandArgs, "--clipboard")
    //tCommandArgs = append(tCommandArgs, "--input")
  }
  
  // このあたりを参考にするとパイプ経由でコマンドの実行結果を得られる
  // http://qiita.com/yuroyoro/items/9358cd25b5f7fe9dd37f
  tCopyCommand := exec.Command(
    tCommandLiteral,
    tCommandArgs...,
  )
  _, tError = tCopyCommand.Output()
  if tError != nil {
    return fmt.Errorf("PasteToClipBoard(): コマンド実行でエラー %v", tError)
  }
  
  os.Remove(tTempFilePath)
  return nil
}
