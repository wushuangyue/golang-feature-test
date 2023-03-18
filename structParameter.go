/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 21:33:24
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 21:55:14
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
    printBook(Book1)
 
    /* 打印 Book2 信息 */
    printBook(Book2)
 }
 
 func printBook( book Books ) {
    fmt.Printf( "Book title : %s\n", book.title)
    fmt.Printf( "Book author : %s\n", book.author)
    fmt.Printf( "Book subject : %s\n", book.subject)
    fmt.Printf( "Book book_id : %d\n", book.book_id)
 }