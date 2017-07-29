package main

import (
	"./em"
	"fmt"
	"log"
)

func main() {
	str1, err := fulltextreplace.StrReplace("ABCＤＥＦＧ１２３４５６７８９")
	str2, err := fulltextreplace.StrReplace("ＡＢＣDEFG１２３")
	if err != nil {
		return
	}
	log.Print("Date")
	fmt.Println(str1);
	fmt.Println("↑To be the same.↓");
	fmt.Println(str2)
}