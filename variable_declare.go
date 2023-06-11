/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-06-11 22:16:09
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-06-11 22:21:32
 * @FilePath: \golang-feature-test\variable_declare.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	var d = true
	fmt.Println(d)

	/*
	* 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	* intVal := 1 相等于：
	* var intVal int
	* intVal =1
	*
	 */

	f := "Runoob" // var f string = "Runoob"

	fmt.Println(f)
}
