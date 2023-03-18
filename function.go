/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 20:43:53
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 20:44:39
 * @FilePath: \golang-feature-test\function.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
   var ret int

   /* 调用函数并返回最大值 */
   ret = max(a, b)

   fmt.Printf( "最大值是 : %d\n", ret )

   c, d := swap("Google", "Runoob")
   fmt.Println(c, d)
}

func swap(x, y string) (string, string) {
	return y, x
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 定义局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}