// Package file 文件操作辅助函数
package file

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"yftxhub/pkg/app"
	"yftxhub/pkg/auth"
	"yftxhub/pkg/helpers"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
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

func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var avatar string
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/avatars/%s/%s/", app.TimenowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	os.MkdirAll(publicPath+dirName, 0755)

	//保存文件
	fileName := randomNameFromUploadFile(file)
	// pulbic/uploads/avatars/2024/04/03/1/sadasddasd.png
	avatarPath := publicPath + dirName + fileName

	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	// 裁剪图片
	img, err := imaging.Open(avatarPath, imaging.AutoOrientation(true))
	if err != nil {
		return avatar, err
	}
	resizeAvatar := imaging.Thumbnail(img, 256, 256, imaging.Lanczos)
	resizeAvatarName := randomNameFromUploadFile(file)
	resizeAvatarPath := publicPath + dirName + resizeAvatarName
	err = imaging.Save(resizeAvatar, resizeAvatarPath)
	if err != nil {
		return avatar, err
	}
	// 删除老文件
	err = os.Remove(avatarPath)
	if err != nil {
		return avatar, err
	}

	return dirName + resizeAvatarName, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
