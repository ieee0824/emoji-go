package emoji

import (
	"testing"
	"image/color"
)

func TestColorToHex(t *testing.T) {
	tests := []struct {
		input color.RGBA
		want hexColor
	}{
		{color.RGBA{0xff, 0xff, 0xff, 0xff},
			"FFFFFFFF",
		},
		{
			color.RGBA{0x01, 0x1a, 0xfa, 0xfB},
			"011AFAFB",
		},
		{
			color.RGBA{0x00, 0xff,0x00,0xff},
			"00FF00FF",
		},
	}

	for _, test := range tests {
		hex := colorToHex(test.input)
		if *hex != test.want {
			t.Fatalf("want %v, but %v:", test.want, *hex)
		}
	}
}

func compareColor(c0, c1 color.Color) bool {
	r0, g0, b0, a0 := c0.RGBA()
	r1, g1, b1, a1 := c0.RGBA()

	if r0 != r1 {
		return false
	} else if g0 != g1 {
		return false
	} else if b0 != b1 {
		return false
	} else if a0 != a1 {
		return false
	}
	return true
}

func TestHexColor_RGBA(t *testing.T) {
	tests := []struct {
		input hexColor
		want *color.RGBA
		err bool
	}{
		{
			"FFF",
			nil,
			true,
		},
		{
			"FFFFFFFF",
			&color.RGBA{0xff, 0xff, 0xff, 0xff},
			false,
		},
	}

	for _, test := range tests {
		got, err := test.input.RGBA()
		if !test.err && err != nil {
			t.Fatalf("should not be error for %v but: %v", test.input, err.Error())
		}
		if test.err && err == nil {
			t.Fatalf("should be error for %v but not:", test.input)
		}
		if test.want == nil && got == nil {
			continue
		}
		if !compareColor(test.want, got) {
			t.Fatalf("want %v, but %v:", test.want, got)
		}
	}
}