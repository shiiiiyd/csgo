### 切片slice

切片在动态扩展时，需要相同的slice, 例如 `a = append(b, 1)` 则会出错。slice在修改值时需要取下标然后修改。

1、 切片是连续内存并且可以动态扩展，由此引发的问题？

- 不同切片之间不可以相互扩展数据，例如：

```go
a := []int
b := []int{1, 2, 3}
c := a
a = append(b ,1)
```

- 修改切片的值未成功，例如：

```go
mySlice := []int{10, 20 ,30}

for _, v := range mySlice{
  value *= 2
}
fmt.Printf("mySlice %+v\n", mySlice)

for idx, _ := range mySlice{
  mySlice[index] *= 2
}
fmt.Printf("mySlice %+v\n", mySlice)
```

循环体中的变量作用域只在循环体中，v 时领时变量（值传递）

### 