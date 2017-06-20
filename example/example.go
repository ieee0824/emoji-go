package main

import (
	"fmt"
	"github.com/ieee0824/emoji-go/fonts"
)

func main(){
	fmt.Println(fonts.New().Get())
}
