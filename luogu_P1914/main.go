package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	str := ""
	fmt.Scan(&str)
	n %= 26
	for i := 0; i < len(str); i++ {
		if rune(int(str[i])+n%26) <= rune(int('z')) {
			fmt.Printf(string(rune(int(str[i]) + n%26)))
		} else {
			fmt.Printf(string(rune(int(str[i]) + n%26 - 26)))
		}

	}
}
