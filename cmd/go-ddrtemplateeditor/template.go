package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

type Template struct {
	img image.Image
}

func NewTemplate(path string) (*Template, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return &Template{img: img}, nil
}

func (t *Template) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, t.img)
	if err != nil {
		return err
	}

	return nil
}

func (t *Template) ReplaceItem(replacement Replacement) {
	img := replacement.Image()
	rect := img.Bounds()
	dstImage := image.NewRGBA(t.img.Bounds())
	draw.Draw(dstImage, dstImage.Bounds(), t.img, image.Point{}, draw.Src)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			dstImage.Set(x, y, img.At(x, y))
		}
	}

	t.img = dstImage
}
