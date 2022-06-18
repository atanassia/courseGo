package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().Unix())
	var n int = 10
	var m [10]int

	for i:=0; i<=n-1; i++{
		m[i] =  rand.Intn(15) + 1
		
	}

	summary := 0	
	fmt.Println(m)
	
	for _, value := range m{
		for _, value2 := range m{
			if value == value2{
				summary++
			}
		}
	}

	if summary - n > 0{
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}

}