/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 22:31:39
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 22:33:59
 * @FilePath: \golang-feature-test\functionAsParamCallback.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

// 声明一个函数类型
type cb func(int) int

func main() {
	testCallBack(1, callBack) //执行函数---testCallBack
}

func testCallBack(x int, f cb) { //定义了一个函数 testCallBack
	f(x) //由于传进来的是callBack函数，该函数执行需要传入一个int类型参数，因此传入x
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}
