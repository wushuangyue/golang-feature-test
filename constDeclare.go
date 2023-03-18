/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 14:23:13
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 14:27:43
 * @FilePath: \golang-feature-test\constDeclare.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
import "unsafe"
const (
    d = "abc"
    e = len(d)
    f = unsafe.Sizeof(d)
)

func main() {
   const LENGTH int = 10
   const WIDTH int = 5  
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)
   println()
   println(a, b, c)  
   println(d, e, f)
}