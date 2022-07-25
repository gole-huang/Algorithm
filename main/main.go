package main

import (
	"Algorithm/Algo"
	"fmt"
)

func main() {
	// 41
	s, p := "b", "?*?"
	fmt.Printf("%#v,%#v, Result: %v\n", s, p, Algo.IsMatch(s, p))
}

/* 44
"b","?*?"
"",""
"ab","?*"
"aa","aa"
"ho","**ho"
"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb","**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb", Result: false
"aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba","*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*", Result: true
"mississippi","m??*ss*?i*pi", Result: false
"abcabczzzde","*abc???de*", Result: true
"","******", Result: true
"aa","a", Result: false
"aa","*", Result: true
"cb","?a", Result: false
"adceb","*a*b", Result: true
"acdcb","a*c?b", Result: false
"abcabczzzde","*abc???de*", Result: true
*/
/*
1,2,0
3,4,-1,1
7,8,9,11,12
*/
/*
{{1,1,1,1,-1,-1,-1,1,0,0},
{1,0,0,0,1,0,0,0,1,0},
{0,0,1,1,1,1,0,1,1,1},
{1,1,0,1,1,1,0,-1,1,1},
{0,0,0,0,1,-1,0,0,1,-1},
{1,0,1,1,1,0,0,-1,1,0},
{1,1,0,1,0,0,1,0,1,-1},
{1,-1,0,1,0,0,0,1,-1,1},
{1,0,-1,0,-1,0,0,1,0,0},
{0,0,-1,0,1,0,1,0,0,1}}
Answer: 22
{{0, 1, -1},
 {1, 0, -1},
 {1, 1,  1}}
Answer: 5
{{1,1,1,1,0,0,0},
{0,0,0,1,0,0,0},
{0,0,0,1,0,0,1},
{1,0,0,1,0,0,0},
{0,0,0,1,0,0,0},
{0,0,0,1,0,0,0},
{0,0,0,1,1,1,1}}
Answer: 15
{{0,0,1,0,0,1,0,1,1,-1,0,0,-1,-1,0,1,1,-1,0,-1},
{1,1,1,0,1,0,0,0,0,1,1,1,1,1,1,1,0,0,1,0},
{1,0,1,1,0,0,1,0,0,0,1,0,1,1,1,-1,0,1,1,0},
{0,1,1,0,0,0,1,0,1,1,0,-1,1,0,0,1,0,0,1,1},
{-1,0,-1,1,0,0,1,1,0,0,1,1,0,-1,1,0,0,0,1,1},
{0,0,1,0,1,1,0,0,1,0,0,1,0,1,1,1,1,1,1,0},
{0,0,0,1,0,1,1,0,0,1,1,-1,1,0,1,1,0,1,1,0},
{0,0,0,0,0,0,1,0,0,1,0,0,1,0,0,0,1,0,1,1},
{0,0,0,-1,1,0,0,1,0,0,1,1,1,1,0,0,0,1,1,0},
{1,0,1,1,1,0,0,1,1,0,1,0,0,0,-1,0,-1,0,1,0},
{0,1,-1,1,1,1,1,0,1,0,0,1,1,0,-1,1,0,0,-1,0},
{0,0,0,0,1,0,1,0,0,-1,0,1,0,-1,0,0,1,0,1,1},
{1,-1,-1,0,0,1,1,1,0,1,1,1,1,1,1,0,0,0,1,0},
{-1,0,1,1,1,1,1,1,0,1,1,1,1,1,0,0,1,0,1,0},
{0,1,-1,1,1,1,0,0,1,-1,1,1,0,1,0,1,0,-1,1,0},
{1,-1,1,0,1,1,1,0,0,0,1,1,1,1,-1,0,0,1,0,-1},
{-1,1,0,0,0,1,1,1,1,1,0,1,1,-1,0,1,0,0,1,0},
{0,0,0,-1,0,1,0,0,0,0,0,0,1,0,1,1,0,0,0,1},
{0,1,0,0,0,0,0,0,0,1,1,1,1,0,0,1,0,0,0,1},
{0,0,0,1,-1,0,-1,1,0,1,0,0,0,0,1,0,0,1,-1,0}}
*/
