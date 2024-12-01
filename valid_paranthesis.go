package main

import "fmt"



func isValid(s string) bool {
	var stack []rune
	var dict = map[rune]rune{'(':')','{':'}','[':']'}

	for i:=0;i<len(s);i++{
		switch{
		case s[i] == '(' || s[i] == '{' || s[i] == '[':
			stack = append(stack,rune(s[i]))
		case len(stack)>0 && dict[rune(stack[len(stack)-1])] == rune(s[i]):
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}


func main(){
	var word string
	fmt.Scanln(&word)
	fmt.Println(isValid(word))
}