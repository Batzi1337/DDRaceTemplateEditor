package models

import (
	"image"
	"image/draw"
)

type Template struct {
	ID   int
	Name string
	Img  image.Image
}

func (t *Template) ReplaceItem(replacement *Item) {
	rect := replacement.Img.Bounds()
	dstImage := image.NewRGBA(t.Img.Bounds())
	draw.Draw(dstImage, dstImage.Bounds(), t.Img, image.Point{}, draw.Src)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			dstImage.Set(x, y, replacement.Img.At(x, y))
		}
	}

	t.Img = dstImage
}
