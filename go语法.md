```
Go语言中一行代表一个语句结束，不需要;结尾。


package main
Go语言类型在变量名之后
var x, y int
var (    // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

g, h := 123, "hello"
//:=这种不带声明格式的只能在函数体中出现
//函数外的每个语句都必须以关键字开始（var，func等等）
func main(){
    g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h)
}

```

### 变量var

 - 局部变量：在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
 - 全局变量：在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

### 切片

 - nil：空切片
 - 函数：len(), cap(), append(), copy()

### 3.method and interface

go没有类，









