package emoji

import (
	"image/color"
	"net/http"
	"strconv"
	"fmt"
	"encoding/json"
	"image"
	"image/png"
	"errors"
)

const (
	API = "https://emoji.pine.moe/emoji"
)

func colorToHex(c color.RGBA) *hexColor {
	s := fmt.Sprintf("%02X%02X%02X%02X", c.R, c.G, c.B, c.A)
	h := hexColor(s)
	return &h
}

type hexColor string

func (h hexColor) String() string {
	return string(h)
}

func (h hexColor) RGBA() (*color.RGBA, error) {
	if len(h.String()) != 8 {
		return nil, errors.New("Incorrect format")
	}
	intColor, err := strconv.ParseUint(h.String(), 16, 32)
	if err != nil {
		return nil, err
	}
	r := uint8((intColor&0xff000000) >> 24)
	g := uint8((intColor&0x00ff0000) >> 16)
	b := uint8((intColor&0x0000ff00) >> 8)
	a := uint8((intColor&0x000000ff) >> 0)

	return &color.RGBA{r, g, b, a}, nil
}

type Emoji struct {
	Body *string `json:"body"`
	Color *color.RGBA `json:"color"`
	BackColor *color.RGBA `json"back_color"`
	api string
	client *http.Client
}

func New(s ...string) *Emoji {
	if len(s) == 0 {
		return &Emoji{
			nil,
			&color.RGBA{0x00, 0x00, 0x00, 0xff},
			&color.RGBA{0xff, 0xff, 0xff, 0x00},
			API,
			&http.Client{},
		}
	}
	return &Emoji{
		&s[0],
		&color.RGBA{0x00, 0x00, 0x00, 0xff},
		&color.RGBA{0xff, 0xff, 0xff, 0x00},
		API,
		&http.Client{},
	}
}

func (e Emoji) String() string {
	m := map[string]interface{}{}
	m["body"] = e.Body
	m["color"] = colorToHex(*e.Color)
	m["back_color"] = colorToHex(*e.BackColor)

	bin, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return ""
	}
	return string(bin)
}

func (e *Emoji) SetBody(s string) *Emoji {
	e.Body = & s
	return e
}

func (e *Emoji) SetHexColor(h hexColor) *Emoji {
	c, err := h.RGBA()
	if err != nil {
		return e
	}
	e.Color = c
	return e
}

func (e *Emoji) SetColor(c color.RGBA) *Emoji {
	e.Color = &c
	return e
}

func (e *Emoji) SetBackHexColor(h hexColor) *Emoji {
	c, err := h.RGBA()
	if err != nil {
		return e
	}
	e.BackColor = c
	return e
}

func (e *Emoji) SetBackColor(c color.RGBA) *Emoji {
	e.BackColor = &c
	return e
}

func (e *Emoji) Generate() (image.Image, error) {
	req, err := http.NewRequest("GET", e.api, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("text", *e.Body)
	q.Add("color", string(*colorToHex(*e.Color)))
	q.Add("back_color", string(*colorToHex(*e.BackColor)))
	req.URL.RawQuery = q.Encode()

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	img, err := png.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}