package main

import (
	"github.com/ieee0824/emoji-go"
	"fmt"
)

func main(){
	e := emoji.New("hoge")

	fmt.Println(e.Generate())
}
