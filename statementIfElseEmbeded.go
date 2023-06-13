/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 22:56:45
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 22:57:22
 * @FilePath: \golang-feature-test\statementIfElseEmbeded.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}
