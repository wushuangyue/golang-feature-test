/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 23:08:22
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 23:09:30
 * @FilePath: \golang-feature-test\statementLoopContinue.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}
