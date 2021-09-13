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
	ip		string
	port	string
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

func IStart() {
	StartCmd.Flags().StringVarP(&ip,"ip","i","localhost","set the ip for the docker daemon to run on")
	StartCmd.Flags().StringVarP(&port,"port","p","8443","set the port for the docker daemon to run on")
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
			"8443/tcp": struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"8443/tcp": []nat.PortBinding{
				{
					HostIP: 	ip,
					HostPort:	port,
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		if errdefs.Conflict(err) != nil {
			fmt.Println("❌ Docker daemon is already running.")
			return nil
		} else {
			return err
		}

	}
	fmt.Println("✅ Succesfully started Docker daemon.")
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
			fmt.Println("✅ Succesfully stopped docker daemon.")
			return nil
		}
	}
	fmt.Println("❌ Couldn't find a running daemon.")
	return nil
}

