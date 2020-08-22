# Goメモ
## slice
### makeで初期化する場合とそうでない場合
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
