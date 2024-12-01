package main

import "fmt"



func longestCommonPrefix(strs []string) string {
	var result string
	min := strs[0]
	for _,val := range strs{
		switch{
		case len(val)<len(min):
			min = val
		}
	}

	for i:=0;i<len(min);i++{
		var ok int = 1
		for ind,_ := range strs{
			switch{
			case !(min[i] == strs[ind][i]):
				ok = 0
				break
			}
		}
		switch{
		case ok == 1:
			result+=string(min[i])
		case ok == 0:
			return result
		}
	}
	return result
}


func main(){
	var word string
	var array []string
	for i:=0;i<5;i++{
		fmt.Scanln(&word)
		array = append(array,word)
	}
	fmt.Println(longestCommonPrefix(array))
}