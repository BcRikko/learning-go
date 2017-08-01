Go開発環境構築
====


## Macの場合

### gvmのインストール

goのバージョン管理ツールである[gvm](https://github.com/moovweb/gvm)をインストールする。
直接goをインストールする方法は後述する。

```bash
# gvmのインストール
$ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

Cloning from https://github.com/moovweb/gvm.git to $HOME/.gvm
No existing Go versions detected
Installed GVM v1.0.22

Please restart your terminal session or to get started right away run
 `source $HOME/.gvm/scripts/gvm`


# 設定ファイルの再読込
$ source $HOME/.gvm/scripts/gvm


# gvmのバージョン確認
$ gvm version
Go Version Manager v1.0.22 installed at $HOME/.gvm
```

### goのインストール

さきほどインストールしたgvmを使って、goをインストールする。

```bash
# 利用できるバージョンの確認
$ gvm listall

# go1.5以降をインストールするために、一旦go1.4をインストールする
# go1.4以上がないとコンパイルできず「ERROR: Cannot find $HOME/go1.4/bin/go.」というエラーになる
$ gvm install go1.4
Installing go1.4...
 * Compiling...
go1.4 successfully installed!

$ gvm use go1.4
Now using version go1.4

# 最新バージョンをインストールする（2017/08/01時点では1.8.3）
$ gvm install go1.8.3
Installing go1.8.3...
 * Compiling...
go1.8.3 successfully installed!

# インストールされたバージョンを確認
$ gvm list
gvm gos (installed)

=> go1.4
   go1.8.3

# go1.8.3を使うようにする
$ gvm use go1.8.3
Now using version go1.8.3

# バージョン確認
$ go version
go version go1.8.3 darwin/amd64
```

### Hello world!

最低限の環境はできたので、サンプルプログラムをつくって実行してみる。

```go
// main.go
package main

import (
    "fmt"
)

func main () {
    fmt.Println("Hello world!")
}
```

実行してみる。

```bash
$ go run main.go
Hello world!
```



## おまけ

### gvmを使わずインストールする(Macの場合)

```bash
# goをインストールする
$ brew install go

# バージョン確認
$ go version
go version go1.8.1 darwin/amd64

# パスを通す
$ cat << EOL >> ~/.bash_profile
## go
if [ -x "`which go`" ]; then
  export GOPATH=\$HOME/.go
  export PATH=\$GOPATH/bin:\$PATH
fi
EOL

# 設定ファイルの反映
$ source ~/.bash_profile
```

直接`.bask_profile`を編集する場合は、バックスラッシュ`\`を削除する。
```bash
if [ -x "`which go`" ]; then
  export GOPATH\$HOME/.go
  export PATH=$GOPATH/bin:$PATH
fi
```
