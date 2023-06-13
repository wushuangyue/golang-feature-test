/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-04-22 10:36:12
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 23:00:18
 * @FilePath: \golang-feature-test\statementLoopForSimple.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
