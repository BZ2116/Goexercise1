/*练习 5.7 Fizz-Buzz 问题：fizzbuzz.go

写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。
遇到 5 的倍数时，打印 Buzz 而不是相应的数字。对于同时为 3 和 5 的倍数的数，打印 FizzBuzz（提示：使用 switch 语句）。*/
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 != 0:
			fmt.Println("Fizz")
		case i%3 != 0 && i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)

		}
	}
}
