package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/moby/moby/client"
	"golang.org/x/net/context"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	check(err)

	cwd, _ := os.Getwd()
	file, err := os.Open(cwd + "/Dockerfile.tar.gz")
	defer file.Close()

	imageName := fmt.Sprintf("%v", time.Now().Unix())
	res, err := cli.ImageBuild(ctx, file, types.ImageBuildOptions{
		Tags:        []string{imageName},
		ForceRemove: true,
	})
	check(err)
	defer res.Body.Close()

	fmt.Printf("OSType: %s\n", res.OSType)

	b, err := ioutil.ReadAll(res.Body)
	check(err)
	fmt.Println(string(b))

	fmt.Printf("Build Image. Image's name is %v\n", imageName)
}
