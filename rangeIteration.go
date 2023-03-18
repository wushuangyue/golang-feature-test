/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-18 22:13:24
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-18 22:13:39
 * @FilePath: \golang-feature-test\range copy.1go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main
 import "fmt"
 func main() {
     //这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
     nums := []int{2, 3, 4}
     sum := 0
     for _, num := range nums {
         sum += num
     }
     fmt.Println("sum:", sum)
     //在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
     for i, num := range nums {
         if num == 3 {
             fmt.Println("index:", i)
         }
     }
     //range 也可以用在 map 的键值对上。
     kvs := map[string]string{"a": "apple", "b": "banana"}
     for k, v := range kvs {
         fmt.Printf("%s -> %s\n", k, v)
     }
 
     //range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
     for i, c := range "go" {
         fmt.Println(i, c)
     }
 }