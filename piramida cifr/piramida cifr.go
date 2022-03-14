package main

import "fmt"

/*программка выводит:
0
0 1
0 1 2
0 1 2 3
и так далее до 0 1 2 3 4 5 6 7 8 9
пирамида цифр
*/
func main() {
	fmt.Println(0)
	fmt.Println(0, 1)
	fmt.Println(0, 1, 2)
	fmt.Println(0, 1, 2, 3)
	fmt.Println(0, 1, 2, 3, 4)
	fmt.Println(0, 1, 2, 3, 4, 5)
	fmt.Println(0, 1, 2, 3, 4, 5, 6)
	fmt.Println(0, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(0, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
}
