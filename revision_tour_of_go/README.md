# Goメモ
## Slice
A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).

### make
When called, make allocates an array and returns a slice that refers to that array

```
var s []byte
s = make([]byte, 5, 5)
// s == []byte{0, 0, 0, 0, 0}
```

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

### length and capacity
- sliceは参照型。capacityはsliceの最初から数えて、元となる配列の要素数。lengthはsliceの要素数を表す。
- Goでは、容量オーバしたスライスに対して新しくメモリ領域を確保してコピーする。
- sliceの値を変えるとベースとなっていた配列の値も書き換わるので注意

## Channel
Closing channel is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop

## Interface
interfaceを引数にしたfunctionは作ることできるけど、interfaceをレシーバーにしたmethodは作れない

## Error
When trying to output error to the console, go looks for Error method. Always implement Error method when creating a custom error

```
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// returning error type
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		// when trying to output error to the console, go looks for Error method
		// always implement Error method when creating a custom error
		fmt.Println(err)
	}
}
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

