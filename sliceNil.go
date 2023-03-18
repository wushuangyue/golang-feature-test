/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 22:01:54
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 22:07:07
 * @FilePath: \golang-feature-test\slice.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import "fmt"
 
 func main() {
    var numbers []int
 
    printSlice(numbers)
 
    if(numbers == nil){
       fmt.Printf("切片是空的")
    }
 }
 
 func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
 }