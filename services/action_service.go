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

type ActionServiceInterface interface {
	Index(ctx *gin.Context)
	Output(ctx *gin.Context)
	ParseActionFile(action *types.Action)
}

type ActionService struct{}

func NewActionService() *ActionService {
	return &ActionService{}
}

func (s *ActionService) ParseActionFile(action *types.Action) {
	actionFile, err := os.ReadFile(path.Join("config", "action.yaml"))

	utils.CheckError(err)

	utils.CheckError(yaml.Unmarshal(actionFile, &action))
}

func (s *ActionService) IsShellCommandAllowed(action *types.Action, command string) bool {
	for _, ActionCommand := range action.Actions {
		actualCommand := ActionCommand.Shell
		for _, arg := range ActionCommand.Arguments {
			placeholder := fmt.Sprintf("{{ %s }}", arg.Name)
			actualCommand = strings.ReplaceAll(actualCommand, placeholder, arg.Default)
		}

		if actualCommand == command {
			return true
		}
	}

	return false
}

func (s *ActionService) Index(ctx *gin.Context) {
	var action types.Action

	s.ParseActionFile(&action)

	for i, actionCommand := range action.Actions {
		encryptedShell, err := utils.Encrypt(actionCommand.Shell)
		utils.CheckError(err)
		action.Actions[i].Shell = encryptedShell
	}

	ctx.HTML(http.StatusOK, "action", gin.H{
		"title":   "Home",
		"actions": action.Actions,
	})
}

func (s *ActionService) Output(ctx *gin.Context) {
	shell := ctx.PostForm("shell")

	shell, err := utils.Decrypt(shell)

	utils.CheckError(err)

	for key, values := range ctx.Request.PostForm {
		if key != "shell" {
			placeholder := fmt.Sprintf("{{ %s }}", key)
			shell = strings.ReplaceAll(shell, placeholder, values[0])
		}
	}

	var action types.Action

	s.ParseActionFile(&action)

	if !s.IsShellCommandAllowed(&action, shell) {
		ctx.HTML(http.StatusBadRequest, "action", gin.H{
			"title":   "Home",
			"actions": action.Actions,
		})
		return
	}

	cmdOutput, err := utils.RunCommand(shell)

	utils.CheckError(err)

	ctx.HTML(http.StatusOK, "output", gin.H{
		"title":  "Output",
		"output": cmdOutput,
	})
}
