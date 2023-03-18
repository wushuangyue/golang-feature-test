/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 22:34:44
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 22:35:54
 * @FilePath: \golang-feature-test\interface.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import "fmt"
 
 type Shape interface {
     area() float64
 }
 
 type Rectangle struct {
     width  float64
     height float64
 }
 
 func (r Rectangle) area() float64 {
     return r.width * r.height
 }
 
 type Circle struct {
     radius float64
 }
 
 func (c Circle) area() float64 {
     return 3.14 * c.radius * c.radius
 }
 
 func main() {
     var s Shape
 
     s = Rectangle{width: 10, height: 5}
     fmt.Printf("矩形面积: %f\n", s.area())
 
     s = Circle{radius: 3}
     fmt.Printf("圆形面积: %f\n", s.area())
 }