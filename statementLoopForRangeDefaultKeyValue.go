/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 23:03:05
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 23:03:45
 * @FilePath: \golang-feature-test\statementLoopForRangeDefault.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
}
