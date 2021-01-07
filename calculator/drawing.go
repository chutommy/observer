package calculator

import (
	"image/color"

	"gocv.io/x/gocv"
)

// Draw draws an Object on the image matrices.
func (o *Object) Draw(img *gocv.Mat, c color.RGBA, thickness int) {
	gocv.Rectangle(img, o.rect, c, thickness)
}

// Draw draws Objects on the image matrices.
func (oo Objects) Draw(img *gocv.Mat, c color.RGBA, thickness int) {
	for _, o := range oo {
		gocv.Rectangle(img, o.rect, c, thickness)
	}
}
