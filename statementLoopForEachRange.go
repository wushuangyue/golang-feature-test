/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 23:02:14
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 23:02:47
 * @FilePath: \golang-feature-test\statementLoopForEachRange.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}
