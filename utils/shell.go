package utils

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func RunCommand(cmdStr string) (string, error) {
	log.Println("Running command: " + cmdStr)
	cmdArgs := strings.Fields(cmdStr)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
