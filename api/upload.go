package api

import (
	"github.com/gin-gonic/gin"
	"mall/global"
	"mall/response"
)

// FileUpload 图片文件上传
func FileUpload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response.Failed("上传图片出错", c)
	}
	image := global.Config.Upload
	err = c.SaveUploadedFile(file, image.SavePath+file.Filename)
	if err != nil {
		response.Failed("上传图片出错", c)
	}
	imageURL := image.AccessUrl + file.Filename
	response.Success("上传图片成功", imageURL, c)
}
