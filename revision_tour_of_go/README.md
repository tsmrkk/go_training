# Goメモ
## slice
### append
By using append function, it is possible to add multiple values to a slice

```
var s []int
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// len=0 cap=0 []
s = append(s, 0)
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// len=1 cap=1 [0]
s = append(s, 1)
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// len=2 cap=2 [0 1]
s = append(s, 2, 3, 4)
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// len=5 cap=6 [0 1 2 3 4]
```

## 疑問点
### makeで初期化する場合とそうでない場合
makeがよくわかってないけど、これってなんでmainの中で一度makeする必要がある?

```
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

これはmが上のプログラムだとただ型を宣言した部分(`var m map[string]Vertex`)と実際にそれに値(新たに生成したmap)を入れているから? 以下のように変更しても動く

```
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//var m map[string]Vertex

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```
