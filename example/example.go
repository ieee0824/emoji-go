package main

import (
	"fmt"
	"github.com/ieee0824/emoji-go/emoji"
)

func main(){
	e := emoji.New("hoge")

	fmt.Println(e.Generate())
}
