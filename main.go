/*
Задание 1. Подсчёт чисел в массиве
Заполните массив неупорядоченными числами на основе генератора случайных чисел.
Введите число. Программа должна найти это число в массиве и вывести, сколько чисел
находится в массиве после введённого. При отсутствии введённого числа в массиве —
вывести 0. Для удобства проверки реализуйте вывод массива на экран.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

var num int

func main() {
	//начиная с версии 1.20, Seed, который демонстрируется в видеоуроках курса, вывели из использования в пакете
	//math для генерации случайных чисел, вместо этого необходимо использовать NewSource
	//примеры:
	/*r := rand.New(rand.NewSource(seed))
	fmt.Println(r.Uint64())*/
	/*// Create a new random number generator with a custom seed (e.g., current time)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	// Generate a random number of minutes between 1 and 15
	randomMinutes := rng.Intn(15) + 1*/

	rand.NewSource(time.Now().UnixNano())
	var a [n]int
	for i := 1; i < n; i++ {
		a[i] = rand.Intn(n)
	}
	fmt.Println("Введите любое однозначное число")
	fmt.Scan(&num)
	if find(a) == 0 {
		fmt.Printf("Результат %v, введенного числа %v в массиве нет", find(a), num)
	} else {
		fmt.Printf("После введенного числа %v, в массиве имеется еще %v чисел", num, find(a))
	}
	fmt.Println(a)
}

func find(arr [n]int) int {
	result := 0
	for i := 0; i < n; i++ {
		if arr[i] == num {
			result = (n - 1) - i
			break
		}
	}
	return result
}
