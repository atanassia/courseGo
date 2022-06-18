package main

import "fmt"

func main(){
	words := make(map[rune]int)
	
	state := "съешь ещё этих мягких французских булок, да выпей чаю"
	for _, value := range state{
		summary := 0
		if _, ok := words[value]; ok{
			continue
		}else{
			for _, value2 := range state{
				if value == value2{
					summary++
				}
			}
			words[value] = summary
		}
	}

	for key, val := range words{
		fmt.Println(string(key), ":", val)
	}
}