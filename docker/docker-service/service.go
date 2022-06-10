package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type ContainerData struct {
	Username    string `json:"username,omitempty"`
	Name        string `json:"name,omitempty"`
	Port        string `json:"port,omitempty"`
	ImageName   string `json:"imageName,omitempty"`
	NetWorkMode string `json:"netWorkMode,omitempty"`
	ContainerID string `json:"containerID,omitempty"`
}

// type ContainerStart struct {
// 	ContainerID string `json:"containerID,omitempty"`
// }

func (c *ContainerData) service(ctx context.Context, cli *client.Client) (string, error) {
	dir, _ := os.Getwd()
	dirList := strings.Split(dir, "/")
	dir = strings.Join(dirList[0:len(dirList)-2], "/")
	cConfig := container.Config{
		ExposedPorts: nat.PortSet{"8084/tcp": {}},
		Env:          []string{"USERNAME=" + c.Username},
		Cmd:          strslice.StrSlice{"/home/simtek/docker/start.sh"},
		Image:        c.ImageName,
	}
	hostConfig := container.HostConfig{
		Binds:         []string{dir + ":/home/simtek/code", dir + "/omlibrary/:/usr/bin/../lib/omlibrary"},
		NetworkMode:   container.NetworkMode(c.NetWorkMode),
		PortBindings:  nat.PortMap{"8084/tcp": []nat.PortBinding{{HostIP: "", HostPort: c.Port}}},
		RestartPolicy: container.RestartPolicy{Name: "always", MaximumRetryCount: 0},
	}
	fmt.Println("配置初始化完成，准备创建容器")
	create, err := cli.ContainerCreate(ctx, &cConfig, &hostConfig, nil, nil, c.Name)
	if err != nil {
		fmt.Println("create err: ", err)
		return create.ID, err
	}
	fmt.Printf("容器创建完成，准备启动， 容器ID: %s \n", create.ID)
	err = cli.ContainerStart(ctx, create.ID, types.ContainerStartOptions{})
	if err != nil {
		return create.ID, err
	}

	fmt.Println("容器启动完成。")
	return create.ID, nil
}

func (c *ContainerData) verify(ctx context.Context, cli *client.Client) error {
	list, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}
	for i := 0; i < len(list); i++ {
		if strings.HasSuffix(list[i].Names[0], c.Name) {
			return errors.New("name字段与其他容器冲突, " + "容器名：" + list[i].Names[0] + ", 需要创建的名称" + c.Name)
		}
		for p := 0; p < len(list[i].Ports); p++ {
			if strconv.Itoa(int(list[i].Ports[p].PublicPort)) == c.Port {
				return errors.New("port字段与其他容器冲突")
			}
		}
	}
	return nil
}

func (c *ContainerData) containerList(ctx context.Context, cli *client.Client) ([]types.Container, error) {
	list, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		fmt.Println("verify err: ", err)
		return []types.Container{}, err
	}
	return list, nil
}

func (c *ContainerData) containerStart(ctx context.Context, containerID string, cli *client.Client) error {
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	return err
}

func (c *ContainerData) containerStop(ctx context.Context, containerID string, cli *client.Client) error {
	err := cli.ContainerStop(ctx, containerID, nil)
	return err
}

func (c *ContainerData) containerRemove(ctx context.Context, containerID string, cli *client.Client) error {
	list, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	for i := 0; i < len(list); i++ {
		if strings.HasPrefix(list[i].ID, containerID) {
			if list[i].State == "running" {
				return errors.New("请停止容器后进行删除")
			}
		}
	}
	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	return err
}
