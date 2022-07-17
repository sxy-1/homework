package main

import "fmt"

var mp [15][15]byte
var a, b, c, d int
var dx = [4]int{-1,0,1,0}
var dy = [4]int{0,1,0,-1}
var d1 int = 0
var d2 int = 0

func move1() {
	for i := 0; i < 4; i++ {
		d1 %= 4
	//	fmt.Println(d1,a+dx[d1] >= 0 && a+dx[d1] <= 9 && b+dy[d1] >= 0 && b+dy[d1] <= 9 && mp[a+dx[d1]][b+dy[d1]] != '*')
		if a+dx[d1] >= 0 && a+dx[d1] <= 9 && b+dy[d1] >= 0 && b+dy[d1] <= 9 && mp[a+dx[d1]][b+dy[d1]] != '*' {
			a += dx[d1]
			b += dy[d1]
			break
		}
		d1++
	}
}
func move2() {
	for i := 0; i < 4; i++ {
		d2 %= 4
		if c+dx[d2] >= 0 && c+dx[d2] <= 9 && d+dy[d2] >= 0 && d+dy[d2] <= 9 && mp[c+dx[d2]][d+dy[d2]] != '*' {
			c += dx[d2]
			d += dy[d2]
			break
		}
		d2++
	}
}
func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Scanf("%c", &mp[i][j])
			if mp[i][j] == 'C' {
				a = i
				b = j
			}
			if mp[i][j] == 'F' {
				c = i
				d = j
			}
		}
		fmt.Scanf("\n")
	}
	step := 0
	for ; step < 200000; step++ {
		if a == c && b == d {
			break
		}
		move1()
		move2()
	//	fmt.Println(a,b,c,d,step)
	}
	if step >= 200000 {
		fmt.Println("0")
	} else {
		fmt.Print(step-1)
	}
	
}
