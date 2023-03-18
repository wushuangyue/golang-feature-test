/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 22:37:56
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 22:40:44
 * @FilePath: \golang-feature-test\goroutine .go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import (
         "fmt"

 )
 
 func fibonacci(n int, c chan int) {
         x, y := 0, 1
         for i := 0; i < n; i++ {
                 c <- x
                 x, y = y, x+y
         }
         close(c)
 }
 
 func main() {
         c := make(chan int, 10)
         go fibonacci(cap(c), c)
         // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
         // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
         // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
         // 会结束，从而在接收第 11 个数据的时候就阻塞了。
         for i := range c {
                 fmt.Println(i)
         }
 }