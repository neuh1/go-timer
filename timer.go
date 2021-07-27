package main

import (
	"fmt"
	"time"
)

func main() {
	var sum int
	var first int
	first = 0
	fmt.Println("Podaj liczbe")
	fmt.Scanln(&sum)
	for sum < 2147483647 {
		if first == 0 {
			fmt.Println("Pozostały czas : ", sum)
			first = first + 1
		}
		time.Sleep(1 * time.Second)
		sum = sum - 1
		if sum != -1 {
			fmt.Println("Pozostały czas : ", sum)
		}
		if sum == 0 {
			fmt.Println("Czas dobiegł końca")
		}
		if sum == -1 {
			time.Sleep(5 * time.Second)
			main()
		}
	}
}
