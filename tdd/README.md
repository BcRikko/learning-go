TDD（テスト駆動開発）をする
====

[Go言語をTDDで学んでみた話 – 時を超えたプログラミングの道](https://twop.agile.esm.co.jp/learning-go-lang-by-tdd-8326723d9362)を参考にすすめる。


## ステップ1: Red(失敗)になることを確認する

* roman.go(./roman.go) ← 機能
* roman_test.go(./roman_test.go) ← テストコード

テストには`testing`パッケージを使う。
```go
import (
    "testing"
)

// Testのあとにはテスト対象を書く（ただし[A-Z]ではじめる）
func TestXxx(t *testing.T) {
    // "actual"の部分でテスト対象を実行 → gotに代入
    if got, want := "actual", "expect"; got != want {
        // gotとwantが違ったらエラー
        t.Errorf("エラーメッセージ")
    }
}
```

JavaScriptでは`expect`, `actual`を使うが、goでは`want`, `got`を使うのが一般的らしい。


[commit:ea70fcb](https://github.com/BcRikko/learning-go/commit/ea70fcb429bfb806cb04ee545af2e3d0ac985109)を実行すると、以下のようになる。  
テストの実行は、VSCodeのエディタ上に表示される`run package tests`や`run test`をクリックするだけでOK。
```
Running tool: /usr/local/bin/go test -timeout 30s -tags 

--- FAIL: TestToRoman (0.00s)
	$GOPATH/src/github.com/BcRikko/learning-go/tdd/roman_test.go:9: ToRoman(1): got _ want I
FAIL
exit status 1
FAIL	github.com/BcRikko/learning-go/tdd	0.009s
Error: Tests failed.
```


## ステップ2: 仮実装でテストを通す


[commit:33b7062](https://github.com/BcRikko/learning-go/commit/33b706269621daf69bd16bca5b59b2f4045c9d26)で仮実装して実行する。

```
Running tool: /usr/local/bin/go test -timeout 30s -tags 

PASS
ok  	github.com/BcRikko/learning-go/tdd	0.008s
Success: Tests passed.
```

公開する関数はパスカルケースで定義する。そのときにコメントがついていないと`severity: 'Warning' message: 'exported function ToRoman should have comment or be unexported'`という警告が表示されるので、関数の直前に`// {Method} hoge`のようなコメントを付ける。

```go
// ToRoman は数値をローマ数字に変換する
func ToRoman (in int16) string {
    // hoge
}
```


## ステップ3: リファクタリング（APIの見直し、関数→メソッドに変更）

独自型`Arabic`を定義して、メソッド`Arabic.ToRoman`を追加する。戻り値にも独自型`Roman`に変更する。

golangにはclass構文がないので、タイプ`type`や構造体`struct`にメソッドを追加できる。

```go
package main

import (
    "fmt"
)

// type
type MyNumber int16

// func (レシーバ) メソッド名() 戻り値
func (num MyNumber) ToString() string {
    return fmt.Sprintf("%v", num)
}


// struct
type MySquare struct {
    height, width float64
}

func (sq MySquare) GetArea() float64 {
    return sq.height * sq.width
}

func main() {
    num := MyNumber(10)
    str := num.ToString()
    fmt.Printf("%T: %s \n", str, str)

    sq := MySquare{3, 4}
    area := sq.GetArea()
    fmt.Printf("%T: %f \n", area, area)
}
```

[commit:0beb2d7](https://github.com/BcRikko/learning-go/commit/0beb2d71ae06417c756bcb636d487154ba56f9d6)を実行し、テストが通ることを確認する。
```
Running tool: /usr/local/bin/go test -timeout 30s -tags 

PASS
ok  	github.com/BcRikko/learning-go/tdd	0.008s
Success: Tests passed.
```


## ステップ4: データ駆動型のテストに書き換える

`struct`の配列を使い、データ駆動型のテストに書き換える。

```go
testCase := []struct {
    desc string
    in   "入力値"
    want "期待値"
}{
    {"テストの説明", "入力値", "期待値"}
}

for _, test := range testCase {
    t.Run(test.desc, func(t *testing.T) {
        if got := "テストしたいメソッド"; got != test.want {
            t.Errorf("エラーメッセージ")
        }
    })
}
```


[commit:f1129cc](https://github.com/BcRikko/learning-go/commit/f1129cc304ddf1556b225ddcff3c7f439c0c9016)を実行し、テストが通ることを確認する。

```
# 失敗した場合
Running tool: /usr/local/bin/go test -timeout 30s -tags  -run ^TestToRoman$

--- FAIL: TestToRoman (0.00s)
    --- FAIL: TestToRoman/1=>Iに変換できること (0.00s)
    	$GOPATH/src/github.com/BcRikko/learning-go/tdd/roman_test.go:19: Arabic.ToRoman(): got _ want I
FAIL
exit status 1
FAIL	github.com/BcRikko/learning-go/tdd	0.009s
Error: Tests failed.
```

```
#成功した場合
Running tool: /usr/local/bin/go test -timeout 30s -tags  -run ^TestToRoman$

PASS
ok  	github.com/BcRikko/learning-go/tdd	0.009s
Success: Tests passed.
```


## ステップ5: テストを追加する

テストケースに1つ追加する。
```go
testCase := []struct {
    desc string
    in   Arabic
    want Roman
}{
    {"1=>Iに変換できること", 1, "I"},
    {"2=>IIに変換できること", 2, "II"},
}
```

[commit:0629165](https://github.com/BcRikko/learning-go/commit/0629165c79b543cd70fdf3a616cfa725f39bd654)を実行し、失敗することを確認する。

```
Running tool: /usr/local/bin/go test -timeout 30s -tags  -run ^TestToRoman$

--- FAIL: TestToRoman (0.00s)
    --- FAIL: TestToRoman/2=>IIに変換できること (0.00s)
    	$GOPATH/src/github.com/BcRikko/learning-go/tdd/roman_test.go:20: Arabic.ToRoman(): got I want II
FAIL
exit status 1
FAIL	github.com/BcRikko/learning-go/tdd	0.008s
Error: Tests failed.
```


## ステップ6: 完成させる

[todo]()を実行し、テストが通ることを確認する。

```
Running tool: /usr/local/bin/go test -timeout 30s -tags  -run ^TestToRoman$

PASS
ok  	github.com/BcRikko/learning-go/tdd	0.008s
Success: Tests passed.
```


## おまけ

ローマ数字の条件

* 表現できる範囲は`0 < num < 4000`
* 使う記号は以下の通り
   * I → 1
   * V → 5
   * X → 10
   * L → 50
   * C → 100
   * D → 500
   * M → 1000
* ただし4つ以上同じ文字を並べてはいけない（○: III、 ×: IIII → IV）
* 文字より小さい数字を表すときは左側、大きい数字を表すときは右側に書く
  * 例 4:IV, 5:V, 6:VI, 9:IX, 10:X, 11:XI