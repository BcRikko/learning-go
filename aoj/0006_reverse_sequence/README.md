[Reverse Sequence \| Aizu Online Judge](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0006)
====

文字列 str を入力し、その文字列を逆順に出力するプログラムを作成して下さい。文字は半角英数字のみで、20 文字以内とします。


## input
```
文字列 str が１行に与えられる。
```

## inputの条件
```
```

## output
str を逆順にした文字列を１行に出力する。


----
## 学んだこと

### stringとrune

C#だと、stringはEnumarableなオブジェクト、それを1文字ずつ取得するとchar型になる。
```c#
var str = "abc";
foreach (char s in str) {
    System.Console.WriteLine(s);
}
```

JavaScriptの場合は、for..ofで1文字ずつ取得できるが、すべてがstring型になる。
```js
const str = "abc";
for (let s of str) {
    console.log(s);
    // a, b, c
}
```

golangの場合、stringはbyte配列なので、1文字ずつ取得しようとするとbyte型の値が取得されてしまう。`string(rune)`をすると文字として取得できる。
```go
str := "abcあいう"
for _, s := range str {
    fmt.Printf("%v %T\n", s, s)
}
// 97 int32
// 98 int32
// 99 int32
// 12354 int32
// 12356 int32
// 12358 int32

for _, s := range str {
    fmt.Println(string(s))
}
// a
// b
// c
// あ
// い
// う
```


## 参考サイト

* [String と Rune — プログラミング言語 Go \| text\.Baldanders\.info](http://text.baldanders.info/golang/string-and-rune/)