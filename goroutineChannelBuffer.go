/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 22:37:56
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 22:39:56
 * @FilePath: \golang-feature-test\goroutine .go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import "fmt"
 
 func main() {
     // 这里我们定义了一个可以存储整数类型的带缓冲通道
         // 缓冲区大小为2
         ch := make(chan int, 2)
 
         // 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
         // 而不用立刻需要去同步读取数据
         ch <- 1
         ch <- 2
 
         // 获取这两个数据
         fmt.Println(<-ch)
         fmt.Println(<-ch)
 }