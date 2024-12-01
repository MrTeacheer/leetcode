package main

import "fmt"


func romanToInt(s string) int{
	var dict = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result int;
	runes:=[]rune(s)
	for i:=0;i<len(runes);{
		switch{
		case i<len(runes)-1 && dict[runes[i]]<dict[runes[i+1]]:
			result+=dict[runes[i+1]]-dict[runes[i]]
			i+=2
		default:
			result+=dict[runes[i]]
			i++
		}
	}
	return result
}

func main(){
	var roman string
	fmt.Scanln(&roman)
	fmt.Println(romanToInt(roman))
}