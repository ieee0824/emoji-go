# emoji-go

emoji generator api for golang.  
https://emoji.pine.moe


# example

```
var img image.Image

img, _ = emoji.New("hoge").Generate()
img, _ = emoji.New("red").SetHexColor("FF0000FF").Generate()
img, _ = emoji.New("red").SetColor(color.RGBA{0xff,0x00,0x00,0xff}).Generate()
```