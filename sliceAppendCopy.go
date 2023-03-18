/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 22:01:54
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 22:11:05
 * @FilePath: \golang-feature-test\slice.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import "fmt"
 
 func main() {
    var numbers []int
    printSlice(numbers)
 
    /* 允许追加空切片 */
    numbers = append(numbers, 0)
    printSlice(numbers)
 
    /* 向切片添加一个元素 */
    numbers = append(numbers, 1)
    printSlice(numbers)
 
    /* 同时添加多个元素 */
    numbers = append(numbers, 2,3,4)
    printSlice(numbers)
 
    /* 创建切片 numbers1 是之前切片的两倍容量*/
    numbers1 := make([]int, len(numbers), (cap(numbers))*2)
 
    /* 拷贝 numbers 的内容到 numbers1 */
    copy(numbers1,numbers)
    printSlice(numbers1)  
 }
 
 func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
 }