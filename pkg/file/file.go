// Package file 文件操作辅助函数
package file

import (
	"os"
	"path/filepath"
	"strings"
)

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileToChcek string) bool {
	if _, err := os.Stat(fileToChcek); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileNameWithOutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
