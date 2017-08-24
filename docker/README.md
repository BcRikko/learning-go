Docker Remote API
====

Docker Remote APIを使ってDockerを操作する。

## Get started

### イメージの取得
```shell
$ docker pull [name]
```

```go
_, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
if err != nil {
    panic(err)
}
```

### コンテナの作成
```shell
$ docker build [path | url]
```

```go
resp, err := cli.ContainerCreate(ctx, &container.Config{}, nil, nil, "")
if err != nil {
    panic(err)
}
```

### コンテナの実行
```shell
$ docker start [container]
```

```go
if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
    panic(err)
}
```

### コンテナの実行が終わるまで待機
```shell
$ docker wait [container]
```

```go
if _, err = cli.ContainerWait(ctx, resp.ID); err != nil {
    panic(err)
}
```

### コンテナのログを取得

```shell
$ docker logs [container]
```

```go
out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
if err != nil {
    panic(err)
}

io.Copy(os.Stdout, out)
```


### Dockerfileからイメージを作成

Dockerfileのままだと使えない？ので、tar.gzで圧縮する。

```shell
$ docker build -t tags .
```

```shell
$ ls
Dockerfile main.go

$ tar -zcvf Dockerfile.tar.gz Dockerfile
$ ls
Dockerfile.tar.gz Dockerfile main.go
```

```go
file, err := os.Open("Dockerfile.tar.gz")
image := "test"

res, err := cli.ImageBuild(ctx, file, types.ImageBuildOptions{
    Tags:        []string{image},
    ForceRemove: true,
})
```

`ForceRemove:true`にすることでイメージをビルドするときの中間コンテナを削除する。