Hello World!!
====

VSCodeでHello Worldするまでについて

----

go開発環境構築手順については、[setup.md](./setup.md)を参照。

----


## プロジェクトディレクトリをつくる

golangでは、一般的に`$GOPATH/src/github.com/{user}/{repository}`のように、`＄GOPATH`配下で開発を行う。`go get 〜`をしたときも`$GOPATH/src/github.com/...`にダウンロードされる。

※ ここでいう`$GOPATH`は`$HOME/go` (`/Users/{user}/go`)を指す


```txt
$GOPATH
└── src
    └── github.com
        └── BcRikko
           └── learning-go
             └── hello.go
```


## Hello world

サンプルプログラムを書く。

```go
// hello.go
package main

import (
    "fmt"
)

func main () {
    fmt.Println("Hello world!!")
}
```


## VSCodeから実行する

`hello.go`を開いて、`F5` 、またはコマンドパレットから`Debug: Start Debugging`を実行する。

Shift+Command+Yで`DEBUG CONSOLE`を開くと、以下のような結果が出力される。
※ 初回実行時はパスワードの入力を求められる。

```txt
2017/08/01 15:22:04 server.go:73: Using API v1
2017/08/01 15:22:04 debugger.go:97: launching process with args: [$GOPATH/src/github.com/BcRikko/learning-go/debug]
API server listening at: 127.0.0.1:2345
2017/08/01 15:22:04 debugger.go:505: continuing
Hello world!!
```