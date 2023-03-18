/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 12:56:15
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 14:03:24
 * @FilePath: \golang-feature-test\hello.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
    //variable declare
	// %d 表示整型数字，%s 表示字符串
	var stockcode=123
	var enddate="2020-12-31"
	var url="Code=%d&endDate=%s"
	var target_url=fmt.Sprintf(url,stockcode,enddate)
	
	//仔细观察输出， 蛮有意思，线程调度导致输出的顺序不一样
	fmt.Println(target_url);
	fmt.Printf(url,stockcode,enddate)
	
	fmt.Println("Hello, World!")

	//字符串连接
	fmt.Println("Google " + "Runoob")

	fmt.Println("============================================")
	var a string = "Runoob"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

	fmt.Println("===============================================")
	// 声明一个变量并初始化
    var d = "RUNOOB"
    fmt.Println(d)

    // 没有初始化就为零值
    var e int
    fmt.Println(e)

    // bool 零值为 false
    var f bool
    fmt.Println(f)

	//float and string
	var i int
    var g float64
    var k bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, g, k, s)

	//第二种变量声明形式：根据值自行判断变量类型
	var d1 = true
    fmt.Println(d1)
	//第三种变量声明形式
	/*
	intVal := 1相当于：
	var intVal int 
    intVal =1 
	*/
	m := "Runoob" // var f string = "Runoob"

    fmt.Println(m)
}
