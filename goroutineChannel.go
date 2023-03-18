/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 22:37:56
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 22:39:17
 * @FilePath: \golang-feature-test\goroutine .go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import "fmt"
 
 func sum(s []int, c chan int) {
         sum := 0
         for _, v := range s {
                 sum += v
         }
         c <- sum // 把 sum 发送到通道 c
 }
 
 func main() {
         s := []int{7, 2, 8, -9, 4, 0}
 
         c := make(chan int)
         go sum(s[:len(s)/2], c)
         go sum(s[len(s)/2:], c)
         x, y := <-c, <-c // 从通道 c 中接收
 
         fmt.Println(x, y, x+y)
 }