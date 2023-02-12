// @Author: Ciusyan 2023/2/7
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

type Handler struct {
	service video.ServiceServer
	l       logger.Logger

	// 提供一个空结构体，用于默认实现方法
	ioc.GinDefault
}

func (h *Handler) Init() error {
	h.l = zap.L().Named(video.AppName)
	h.service = ioc.GetGrpcDependency(video.AppName).(video.ServiceServer)
	return nil
}

func (h *Handler) Name() string {
	return video.AppName
}

func (h *Handler) Version() string {
	return ""
}

func (h *Handler) Registry(r gin.IRoutes) {
	r.GET("/feed/", h.feed)
}

func (h *Handler) RegistryWithMiddle(r gin.IRoutes) {
	r.POST("/publish/action/", h.publishAction)
	r.GET("/publish/list/", h.publishList)
}

func init() {
	ioc.GinDI(&Handler{})
}