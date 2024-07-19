package controllers

import (
	"dvtool/services"
	"dvtool/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActionController struct {
	ActionService services.ActionServiceInterface
}

func NewActionController(actionService services.ActionServiceInterface) *ActionController {
	return &ActionController{ActionService: actionService}
}

func (c *ActionController) Index(ctx *gin.Context) {
	c.ActionService.RefreshEncryptionKey()
	c.ActionService.Index(ctx)
}
func (c *ActionController) Output(ctx *gin.Context) {
	var actionValidator types.ActionValidator
	var action types.Action

	c.ActionService.ParseActionFile(&action)

	if err := ctx.ShouldBind(&actionValidator); err != nil {
		ctx.HTML(http.StatusBadRequest, "action", gin.H{
			"title":   "Home",
			"actions": action.Actions,
		})

		return
	}

	c.ActionService.Output(ctx)
}
