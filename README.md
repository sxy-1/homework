## golang第一周任务题解

### 1、凯撒密码   [P1914 小书童——凯撒密码 - 洛谷 ](https://www.luogu.com.cn/problem/P1914)

```go
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

```

傻逼题 无题解。不确定声明和make是否需要。

### 2、计算机的改良 [P1022 计算器的改良 - 洛谷](https://www.luogu.com.cn/problem/P1022)

```go
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

```

题解是抄的  [计算器的改良 - 孤 - 洛谷博客 (luogu.com.cn)](https://www.luogu.com.cn/blog/109358/ji-suan-qi-di-gai-liang)

### 3、USACO的牛 [P1518 [USACO2.4\]两只塔姆沃斯牛 The Tamworth Two - 洛谷](https://www.luogu.com.cn/problem/P1518)

```go
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

```

tnnd 一直WA 到现在还不知道啥子情况 等会实名制下一下