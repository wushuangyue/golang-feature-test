/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 22:41:16
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 22:41:44
 * @FilePath: \golang-feature-test\functionClosure.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
