package kubectl

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetCurrentCluster() (string, error) {
	current := exec.Command("kubectl", "config", "current-context")
	currentOutput, err := current.Output()
	if err != nil {
		return "", err
	}
	currentContext := strings.TrimSpace(string(currentOutput))
	if currentContext == "" {
		return "", fmt.Errorf("no current context found")
	}
	cmd := exec.Command("kubectl",
		"config",
		"view",
		"-o",
		fmt.Sprintf("jsonpath={.contexts[?(@.name==%q)].context.cluster}", currentContext))
	cmd.Env = os.Environ()
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	cluster := strings.TrimSpace(string(output))
	if cluster == "" {
		return "", fmt.Errorf("no cluster found for the current context")
	}

	return cluster, nil
}
