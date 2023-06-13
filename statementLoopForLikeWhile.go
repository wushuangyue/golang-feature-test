/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 23:00:22
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 23:01:14
 * @FilePath: \golang-feature-test\statementLoopForLikeWhile.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
}
