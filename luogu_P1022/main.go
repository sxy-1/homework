package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	var a byte
	f := 1
	now := 1
	k, b, x := 0, 0, 0
	r := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '-' {
			b += now * f * x
			x = 0
			f = -1
			r = 0
		}
		if str[i] == '+' {
			b += now * f * x
			x = 0
			f = 1
			r = 0
		}
		if str[i] == '=' {
			b += now * f * x
			x = 0
			f = 1
			now = -1
			r = 0
		}
		if str[i] >= 'a' && str[i] <= 'z' {
			if r != 0 {
				k += now * f * x
				x = 0
			} else {
				k += now * f
			}
			a = str[i]
			r = 0
		}
		if str[i] >= '0' && str[i] <= '9' {
			x = x*10 + int(str[i]-'0')
			r = 1
		}
	}
	b += now * f * x
	var ans float32
	
	ans = float32(b) / float32(k)
	ans = -ans
	if ans == -0.0 {
		ans = 0
	}

	fmt.Printf("%c=%.3f\n", a, ans)

}
