package wallpaper

import (
	"webserver/app/service/response"
	"webserver/app/service/wallpaper"

	"github.com/gogf/gf/net/ghttp"
)

/********************上传图片********************/

//UploadPicRequest 上传图片请求结构体
type UploadPicRequest struct {
	ID      string
	PicName string
	Content string
}

// Controller  用户API管理对象
type Controller struct{}

type PicType uint8

const (
	PicTypeJPG PicType = 0
	PicTypePNG PicType = 1
	PicTypeGIF PicType = 2
	PicTypeBMP PicType = 3
)

func (c *Controller) UploadPic(r *ghttp.Request) {
	var data *UploadPicRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := wallpaper.EncodePic(data.Content, data.PicName); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := wallpaper.EncodePic(data.Content, data.PicName); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", UploadPicResponse{
		data.ID,
	})
}

type UploadPicResponse struct {
	ID string
}

/********************获取图片********************/
//GetPicRequest 获取图片请求结构体
type GetPicRequest struct {
	ID string
}

func (c *Controller) GetPic(r *ghttp.Request) {
	var data *GetPicRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

}
