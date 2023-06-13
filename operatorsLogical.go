/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-06-11 22:34:18
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-06-11 22:34:39
 * @FilePath: \golang-feature-test\operators_logical.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}
