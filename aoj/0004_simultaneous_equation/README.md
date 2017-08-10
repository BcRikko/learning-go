[Simultaneous Equation \| Aizu Online Judge](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0004)
====

連立方程式
```
ax+by=c
dx+ey=f
```
の解、x, y を出力するプログラムを作成して下さい。a, b, c, d, e, f はそれぞれ、 -1,000 以上 1,000 以下の実数とし、連立方程式の解が一意に存在するように与えられるものとします。



## input
```
複数のデータセットが与えられます。入力の最後まで処理して下さい。１つのデータセットが１行に与えられます。１つのデータセットに a, b, c, d, e, f が１つのスペースで区切られて与えられます。
```

## inputの条件
* -1,000 <= value <= 1,000
* 実数

## output
各データセットに対して、x, y を１つのスペースで区切って１行に出力して下さい。各値は小数点以下第３位まで出力して下さい。小数点以下第４位を四捨五入して下さい。


----
## 学んだこと

### フォーマット時の小数点以下の指定

このあたりはC言語などとほとんど同じ
```go
// float
val := 1.23456

// 小数点以下3桁、小数第4位で四捨五入
fmt.Printf("%.3f", val)
```


## 参考サイト

* [fmt#hdr-Printing \- The Go Programming Language](https://golang.org/pkg/fmt/#hdr-Printing)