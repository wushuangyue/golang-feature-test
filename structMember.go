/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 21:33:24
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 21:54:15
 * @FilePath: \golang-feature-test\struct1.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import "fmt"
 
 type Books struct {
    title string
    author string
    subject string
    book_id int
 }
 
 func main() {
    var Book1 Books        /* 声明 Book1 为 Books 类型 */
    var Book2 Books        /* 声明 Book2 为 Books 类型 */
 
    /* book 1 描述 */
    Book1.title = "Go 语言"
    Book1.author = "www.runoob.com"
    Book1.subject = "Go 语言教程"
    Book1.book_id = 6495407
 
    /* book 2 描述 */
    Book2.title = "Python 教程"
    Book2.author = "www.runoob.com"
    Book2.subject = "Python 语言教程"
    Book2.book_id = 6495700
 
    /* 打印 Book1 信息 */
    fmt.Printf( "Book 1 title : %s\n", Book1.title)
    fmt.Printf( "Book 1 author : %s\n", Book1.author)
    fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
    fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)
 
    /* 打印 Book2 信息 */
    fmt.Printf( "Book 2 title : %s\n", Book2.title)
    fmt.Printf( "Book 2 author : %s\n", Book2.author)
    fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
    fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
 }