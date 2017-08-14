[Digit Number \| Aizu Online Judge](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0002)
====

与えられた２つの整数 a と b の和の桁数を出力するプログラムを作成して下さい。

## input
```
複数のデータセットが与えられます。各データセットは １ 行に与えられます。各データセットは２つの整数 a と b が１つのスペースで区切られて与えられます。入力の終わりまで処理して下さい。
```

## inputの条件
* 0 ≤ a, b ≤ 1,000,000
* データセットの数 ≤ 200

## output
各データセットごとに、a+b の桁数を出力して下さい。


----
## 学んだこと

### 文字列の分割

```go
import (
    "fmt"
    "strings"
)

func main() {
    str := "1 2 3 4 5"
    for _, s := range strings.Split(str, " ") {
        fmt.Println(s)
        // 1, 2, 3, 4, 5
    }
}
```


### 文字列のトリム

```go
import (
    "fmt"
    "strings"
)

func main() {
    // スペースのトリム
    str := "   a   "
    trimmed := strings.TrimSpace(str)
    fmt.Println(trimmed)
    // a

    // スペース以外のトリム
    str2 := "@@@a@@@"
    trimmed2 := strings.Trim(str2, "@")
    fmt.Println(trimmed2)
    // a
}
```


### 文字列 ⇔ 整数の変換

```go
import (
    "fmt"
    "strconv"
)

func main() {
    // string -> int
    str1 := "1"
    if num1, err := strconv.Atoi(str1); err == nil {
        fmt.Printf("num1: %v %T", num1, num1)
        // num1: 1 int
    }

    // int -> string
    num2 := 1
    str2 := strconv.Itoa(num2)
    fmt.Printf("str2: %v %T", str2, str2)
    // str2: 1 string

    // Atoi = ParseInt(val, 10, 0)
    str3 := "1"
    if num3, err := strconv.ParseInt(str3, 10, 0); err == nil {
        fmt.Printf("num3: %v %T", num3, num3)
        // num2: 1 int64
    }
}
```


## 参考サイト
* strings
    * [strings#Split \- The Go Programming Language](https://golang.org/pkg/strings/#Split)
    * [strings#TrimSpace \- The Go Programming Language](https://golang.org/pkg/strings/#TrimSpace)
* strconv
    * [strconv#Atoi \- The Go Programming Language](https://golang.org/pkg/strconv/#Atoi)
    * [strconv#Itoa \- The Go Programming Language](https://golang.org/pkg/strconv/#Itoa)
    * [strconv#ParseInt \- The Go Programming Language](https://golang.org/pkg/strconv/#ParseInt)