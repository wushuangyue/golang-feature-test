/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 23:01:29
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 23:01:57
 * @FilePath: \golang-feature-test\statementLoopForDead.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++ // 无限循环下去
	}
	fmt.Println(sum) // 无法输出
}
