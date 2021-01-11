package geometry

import (
	"image/color"
)

// Color stores RGB color values and a thickness value.
type Color struct {
	r         uint8
	g         uint8
	b         uint8
	thickness int
}

// NewColor constructs a new Color.
func NewColor(r, g, b, t int) *Color {
	return &Color{
		r:         uint8(r),
		g:         uint8(g),
		b:         uint8(b),
		thickness: t,
	}
}

// ToRGBA returns Color in color.RGBA.
func (c *Color) ToRGBA() color.RGBA {
	return color.RGBA{
		R: c.r,
		G: c.g,
		B: c.b,
		A: 0,
	}
}

// T returns thickness of the color line.
func (c *Color) T() int {
	return c.thickness
}
