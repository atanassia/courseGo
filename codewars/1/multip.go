package main

import "fmt"

func main(){
	number := 10
	count :=0
	for i:=1; i<number; i++{
		if i % 3 == 0 || i % 5 == 0{
			count+=i
		}
	}

	fmt.Println(count)
}