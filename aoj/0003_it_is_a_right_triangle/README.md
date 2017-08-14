[Is it a Right Triangle? \| Aizu Online Judge](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0003)
====

３つの正の整数を入力し、それぞれの長さを３辺とする三角形が直角三角形である場合には YES を、違う場合には NO と出力するプログラムを作成して下さい。

## input
```
複数のデータセットが与えられます。１行目にデータセット数の N が与えられます。続いて N 行の入力が与えれます。各行に３つの整数が１つのスペースで区切られて与えられます。
```

## inputの条件
* 1 ≤ 1辺の長さ ≤ 1,000
* N ≤ 1,000

## output
各データセットごとに、YES または NO を１行に出力して下さい。


----
## 学んだこと

### Arrayのmap

JavaScriptの`Array.prototype.map`のような機能はgolangにはないので自分で実装する。またジェネリクスの仕組みもないので型を決め打ちでつくる。

```go
func mapInt(items []int, f func(v int) int) []int {
    ret := make([]int, len(items))
    for i, v := range items {
        ret[i] = f(v)
    }
    return ret
}

func main() {
    ary := []int{1,2,3,4,5}
    pows := mapInt(ary, func (v int) int {
        return v * v
        // math.Powもあるけどfloat64限定
    })

    fmt.Println(pows)
    // [1 4 9 16 25]
}
```

### golangでもコレクションの操作をする

レシーバを使うと、もっと`Array.prototype.map`っぽくなる。
```go
type T string
type ArrayT []T
func (a ArrayT) Map(fn func(item T) T) []T {
    var ret []T
    for _, item := range a {
        ret = append(ret, fn(item))
    }
    return ret
}

func main() {
    ary := ArrayT{"hoge", "fuga", "piyo"}
    mapped := ary.Map(func(item T) T {
        return item + item
    })

    fmt.Println(mapped)
    // [hogehoge, fugafuga, piyopiyo]
}
```

### makeによる割当

`make`を使うことで、スライス、マップ、チャネルの初期化（割当）ができる。  
スライスと配列の違いは別途まとめる。
```go
// 要素数10個の整数型配列
ints := make([]int, 10)
// [0 0 0 0 0 0 0 0 0 0]
fmt.Printlf(ints)
```

### 配列に要素を追加する

```go
var ints []int
fmt.Println(ints)
// []

ints = append(ints, 10)
fmt.Println(ints)
// [10]

ints = append(ints, 5)
fmt.Println(ints)
// [10 5]
```

## 参考サイト

* [Go by Example: Collection Functions](https://gobyexample.com/collection-functions)
* [実践Go言語 \- golang\.jp](http://golang.jp/effective_go#allocation_make)
* [builtin#make \- The Go Programming Language](https://golang.org/pkg/builtin/#make)