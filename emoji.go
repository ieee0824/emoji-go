package emoji

import (
	"image/color"
	"net/http"
	"github.com/pkg/errors"
	"strconv"
	"fmt"
	"encoding/json"
	"image"
	"image/png"
)

const (
	API = "https://emoji.pine.moe/emoji"
)

type Emoji struct {
	Body *string `json:"body"`
	Color *color.RGBA `json:"color"`
	BackColor *color.RGBA `json"back_color"`
	client *http.Client
}

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
	s := string(h)
	if len([]rune(s)) != 8 {
		return nil, errors.New("Incorrect format")
	}

	rString := s[:2]

	s = s[2:]
	gString := s[:2]

	s = s[2:]
	bString := s[:2]

	s = s[2:]
	aString := s[:2]

	r, err := strconv.ParseInt(rString, 16, 16)
	if err != nil {
		return nil, err
	}

	g, err := strconv.ParseInt(gString, 16, 16)
	if err != nil {
		return nil, err
	}

	b, err := strconv.ParseInt(bString, 16, 16)
	if err != nil {
		return nil, err
	}

	a, err := strconv.ParseInt(aString, 16, 16)
	if err != nil {
		return nil, err
	}

	return &color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, nil
}

func New(s ...string) *Emoji {
	if len(s) == 0 {
		return &Emoji{
			nil,
			&color.RGBA{0x00, 0x00, 0x00, 0xff},
			&color.RGBA{0xff, 0xff, 0xff, 0x00},
			&http.Client{},
		}
	}
	return &Emoji{
		&s[0],
		&color.RGBA{0x00, 0x00, 0x00, 0xff},
		&color.RGBA{0xff, 0xff, 0xff, 0x00},
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
	req, err := http.NewRequest("GET", API, nil)
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