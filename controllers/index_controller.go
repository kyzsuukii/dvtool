package controllers

import (
	"dvtool/services"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	IndexService services.IndexServiceInterface
}

func NewIndexController(indexService services.IndexServiceInterface) *IndexController {
	return &IndexController{IndexService: indexService}
}

func (c *IndexController) Index(ctx *gin.Context) {
	c.IndexService.Index(ctx)
}
