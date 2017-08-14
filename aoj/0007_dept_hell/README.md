[Debt Hell \| Aizu Online Judge](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0007)
====

某国に住んでいる友達がお金に困って、あるヤミ金融業者から 10 万円の借金をしたまま、全く返済していないといいます。この業者は、一週間ごとに 5% の利子を借金に加え、さらに借金の 1,000 円未満を切り上げます。  
整数 n を入力し、n 週間後の借金の残高を出力するプログラムを作成して下さい。

## input
```
整数nが１行に与えられる。
```

## inputの条件
* 0 <= n <= 100

## output
n 週間後の返済額を１行に出力する。

----
## 学んだこと

### 数値の切り捨て、切り上げ

```go
import (
    "fmt"
    "math"
)

func main() {
    num := 123.45

    floor := math.Floor(num)
    fmt.Println(floor)
    // 123

    ceil := math.Ceil(num)
    fmt.Println(ceil)
    // 124
}
```


## 参考サイト

* [math#Ceil \- The Go Programming Language](https://golang.org/pkg/math/#Ceil)
* [math#Floor \- The Go Programming Language](https://golang.org/pkg/math/#Floor)