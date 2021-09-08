package commands

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
)

var (
	imageName = "trizlybear/discless:latest"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Discless background daemon",
	RunE:  StartDaemon,
}

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the Discless background daemon",
	RunE:  StopDaemon,
}

func StartDaemon(c *cobra.Command, args []string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	_, err = cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	// io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"8080/tcp": struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"8080/tcp": []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
					HostPort: "8080",
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		if errdefs.Conflict(err) != nil {
			fmt.Println("Docker daemon is already running.")
			return nil
		} else {
			return err
		}

	}
	fmt.Println("Succesfully started Docker daemon.")
	return nil
}

func StopDaemon(c *cobra.Command, args []string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == imageName {
			if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
				return err
			}
			fmt.Println("Succesfully stopped docker daemon.")
			return nil
		}
	}
	fmt.Println("Couldn't find a running daemon.")
	return nil
}

