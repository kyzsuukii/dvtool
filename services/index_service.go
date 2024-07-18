package services

import (
	"dvtool/types"
	"dvtool/utils"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type IndexServiceInterface interface {
	Index(ctx *gin.Context)
	Output(ctx *gin.Context)
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

func (s *IndexService) Output(ctx *gin.Context) {
	shell := ctx.PostForm("shell")

	for key, values := range ctx.Request.PostForm {
		if key != "shell" {
			placeholder := fmt.Sprintf("{{ %s }}", key)
			shell = strings.ReplaceAll(shell, placeholder, values[0])
		}
	}

	cmdOutput, err := utils.RunCommand(shell)

	utils.CheckError(err)

	ctx.HTML(http.StatusOK, "output", gin.H{
		"title":  "Output",
		"output": cmdOutput,
	})
}
