package main

import (
	"io"
	"log"
	"os"

	// typeだけdockerを参照しているので`cannot use ... as type ...`というエラーになる
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/moby/moby/client"

	"golang.org/x/net/context"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func runImage(id string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	check(err)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: id,
	}, nil, nil, "")
	check(err)

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	check(err)

	_, err = cli.ContainerWait(ctx, resp.ID)
	check(err)

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
	})
	check(err)

	io.Copy(os.Stdout, out)

	err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
	check(err)
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("requires at least 1 argument(s). Image IDs")
	}

	ids := os.Args[1:]
	for _, id := range ids {
		runImage(id)
	}
}
