/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 22:30:43
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 22:31:19
 * @FilePath: \golang-feature-test\functionReturnMultiValue.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}
