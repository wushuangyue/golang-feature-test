/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 21:13:33
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 21:15:56
 * @FilePath: \golang-feature-test\pointer
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )

   var  ptr *int

   fmt.Printf("ptr 的值为 : %x\n", ptr  )
   /*
   if(ptr != nil)      ptr 不是空指针
   if(ptr == nil)      ptr 是空指针
   */
}