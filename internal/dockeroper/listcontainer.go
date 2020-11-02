package dockeroper

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

func ListContainer() ([]string, error) {
	tmpSli := make([]string, 0, 10)
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Printf("err : %v\n", err)
		return nil, err
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Printf("err : %v\n", err)
		return nil, err
	}
	for _, container := range containers {
		containerName := container.Names[0]
		if strings.HasPrefix(containerName, "/") {
			containerName = strings.Trim(containerName, "/")
			tmpSli = append(tmpSli, containerName)
		}

	}
	return tmpSli, nil
}
