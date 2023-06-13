/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-04-22 10:36:12
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 22:46:27
 * @FilePath: \golang-feature-test\rangeSimple.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
