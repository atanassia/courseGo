package main

import (
	"fmt"
	"unicode/utf8"
	_ "unicode/utf8"
)

func main(){
	nihongo := []byte("日本語")
	// fmt.Println(string(nihongo[2]))
	// for index, runeValue := range nihongo {
	// 	fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	// }

	fmt.Println(utf8.RuneCount(nihongo))
	// for i := 0; i<len(nihongo); i++  {
	// 	fmt.Printf("%#U starts at byte position %d\n", nihongo[i], i)
	// }
}