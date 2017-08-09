[List of Top 3 Hills \| Aizu Online Judge](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0001)
====

山の高さをメートル単位の整数で表した 10 個のデータがあります。その 10 個のデータを読み込み、その中で、高い順から３つ出力するプログラムを作成して下さい。  
高さが同じ山が複数ある場合でも、異なる順位を付けるものとします。

## input
```
山の高さ1
山の高さ2
     .
     .
山の高さ10
```

## inputの条件
```
0 ≤ 山の高さ ≤ 10,000
```

## output
```
１番目に高い山の高さ
２番目に高い山の高さ
３番目に高い山の高さ
```

----
## 学んだこと

### 数値の配列を含む構造体の定義

```go
test := struct {
    ary []int
}{
    []int{1,2,3}
}
```

### 配列のソート（整数型）

#### 昇順

```go
import (
    "fmt"
    "sort"
)

ary := []int{5,2,4,3,1}
sort.Ints(ary)

fmt.Println(ary)
// [1, 2, 3, 4, 5]
```

#### 降順

```go
import (
    "fmt"
    "sort"
)

ary := []int{5,2,4,3,1}
sort.Sort(sort.Reverse(sort.IntSlice(ary)))

fmt.Println(ary)
// [5, 4, 3, 2, 1]
```



### 配列同士の比較

```go
import (
    "fmt"
    "reflect"
)

x := []int[1,2,3]
y := []int[1,2,3]

if reflect.DeepEqual(x, y) {
    fmt.Println("Same")
} else {
    fmt.Println("Different")
}
```


## 参考サイト

* [sort \- The Go Programming Language](https://golang.org/pkg/sort/)
* [reflect \- The Go Programming Language](https://golang.org/pkg/reflect/#DeepEqual)