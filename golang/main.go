package main

import(
	"fmt"
	"isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Println("Valova Alyona Andreevna")
	fmt.Println("Task А:\n", lab4.TaskA(0.95, 2.0, 1.25, 2.75, 0.3))
	fmt.Println("Task В:\n", lab4.TaskB(2.0, 0.95, [5]float64{2.2, 3.78, 4.51, 6.58, 1.2}))
}
