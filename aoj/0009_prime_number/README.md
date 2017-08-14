[Prime Number \| Aizu Online Judge](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0009)
====

6 桁以下の正の整数 n を入力し、n 以下の素数がいくつあるかを出力するプログラムを作成して下さい。ただし、素数とは 1 と自分自身でしか割り切れない正の整数のうち 1 をのぞいたものをいいます。例えば 10 以下の素数は、2, 3, 5, 7 です。


## input
```
複数のデータセットが与えられます。各データセットに n (1 ≤ n ≤ 999,999) が１行に与えられます。入力の最後まで処理して下さい。
データセットの数は 30 を越えません。
```

## inputの条件
```
```

## output
各データセットごとに、n 以下の素数の個数を１行に出力して下さい。



----
## 学んだこと

### 平方根の求め方

```go
import (
    "fmt"
    "math"
)

func main() {
    num := 16
    sq := math.Sqrt(float64(num))

    fmt.Println(sq)
}
```



## 参考サイト

* [math#Sqrt \- The Go Programming Language](https://golang.org/pkg/math/#Sqrt)