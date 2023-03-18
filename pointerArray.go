/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 21:19:09
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 21:19:55
 * @FilePath: \golang-feature-test\pointerArray.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

const MAX int = 3

func main() {
   a := []int{10,100,200}
   var i int
   var ptr [MAX]*int;

   for  i = 0; i < MAX; i++ {
      ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
   }

   for  i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
   }
}