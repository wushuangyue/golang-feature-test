/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 14:29:11
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 14:29:33
 * @FilePath: \golang-feature-test\iota.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}