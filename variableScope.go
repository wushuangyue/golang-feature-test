package main

import "fmt"

/* 声明全局变量 */
var g int

func main() {
   /* 声明局部变量 */
   var a, b, c int

   /* 初始化参数 */
   a = 10
   b = 20
   c = a + b
   g = a + b 

   fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)
   fmt.Printf ("结果： a = %d, b = %d and g = %d\n", a, b, g)

   /* 声明局部变量 */
   var g int = 10
   fmt.Printf ("结果： g = %d\n",  g)
}