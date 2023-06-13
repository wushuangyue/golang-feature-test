/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-06-13 22:28:06
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-06-13 22:28:38
 * @FilePath: \golang-feature-test\methdo.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"

)

/* 定义结构体 */
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
