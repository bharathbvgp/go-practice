package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี";
	fmt.Println("Len : " , len(s));
	fmt.Println(s);
	// printing all bytes stored for the string s
	for i := 0 ; i < len(s) ; i++ {
		fmt.Printf("%x \n" , s[i]);
	}

	// getting rune count 
	fmt.Println("Rune count : " , utf8.RuneCountInString(s));

	// printing rune value and idx 
	for idx , runeValue := range s {
		fmt.Printf("%#U starts at %d\n" , runeValue , idx);
	}


	for i,w := 0,0 ; i < len(s) ; i += w {
		runeValue , width := utf8.DecodeLastRuneInString(s[i:]);
		fmt.Printf("%U starts at %d\n" , runeValue , i);
		w = width;
	}


}