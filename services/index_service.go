package services

import (
	"dvtool/types"
	"dvtool/utils"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type IndexServiceInterface interface {
	Index(ctx *gin.Context)
	ParseActionFile(action *types.Action)
}

type IndexService struct{}

func NewIndexService() *IndexService {
	return &IndexService{}
}

func (s *IndexService) ParseActionFile(action *types.Action) {
	actionFile, err := os.ReadFile(path.Join("config", "action.yaml"))

	utils.CheckError(err)

	utils.CheckError(yaml.Unmarshal(actionFile, &action))
}

func (s *IndexService) Index(ctx *gin.Context) {
	var action types.Action

	s.ParseActionFile(&action)

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":   "Home",
		"actions": action.Actions,
	})
}
