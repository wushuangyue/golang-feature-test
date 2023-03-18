/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 21:20:02
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 21:27:16
 * @FilePath: \golang-feature-test\pointerPointer.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {

   var a int
   var ptr *int
   var pptr **int

   a = 3000

   /* 指针 ptr 地址 */
   ptr = &a

   /* 指向指针 ptr 地址 */
   pptr = &ptr

   /* 获取 pptr 的值 */
   fmt.Printf("变量 a = %d\n", a )
   fmt.Printf("指针变量 *ptr = %d\n", *ptr )
   fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
   fmt.Printf("指针ptr=%s\n",ptr);
   fmt.Printf("指针ptr=%s\n",pptr);
}