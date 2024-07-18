package controllers

import (
	"dvtool/services"
	"dvtool/types"
	"net/http"

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
func (c *IndexController) Output(ctx *gin.Context) {
	var actionValidator types.ActionValidator
	var action types.Action

	c.IndexService.ParseActionFile(&action)

	if err := ctx.ShouldBind(&actionValidator); err != nil {
		ctx.HTML(http.StatusBadRequest, "index", gin.H{
			"title":   "Home",
			"actions": action.Actions,
		})

		return
	}

	c.IndexService.Output(ctx)
}
