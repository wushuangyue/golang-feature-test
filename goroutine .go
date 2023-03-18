/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 22:37:56
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 22:38:30
 * @FilePath: \golang-feature-test\goroutine .go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
        "fmt"
        "time"

)

func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}

func main() {
        go say("world")
        say("hello")
}