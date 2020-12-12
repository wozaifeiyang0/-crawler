package fileutil

import (
	thy "github.com/wozaifeiyang0/thylog"
	"os"
	"path/filepath"
	"time"
)

// CreateDateDir 根据当前日期来创建文件夹
func CreateDateDir(Path string) string {
	folderName := time.Now().Format("2006-01-02 15:04:05")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, os.ModePerm)
		os.Chmod(folderPath, os.ModePerm)
		thy.Info.Println("创建文件目录：" + folderPath)

	}
	return folderPath
}
