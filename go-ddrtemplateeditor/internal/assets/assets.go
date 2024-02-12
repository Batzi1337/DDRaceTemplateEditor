package assets

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"log"
)

type Item struct {
	img image.Image
}

type Template struct {
	img image.Image
	bP  *[]byte
}

func NewTemplate(img image.Image) *Template {
	t := &Template{img: img}
	return t
}

func (t *Template) NewItem(x, y, width, height int) *Item {
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t.img, image.Point{x, y}, draw.Src)

	return &Item{img: subImage}
}

func (t *Template) ReplaceItems(items ...*Item) {
	for _, item := range items {
		rect := item.img.Bounds()
		dstImage := image.NewRGBA(t.img.Bounds())
		draw.Draw(dstImage, dstImage.Bounds(), t.img, image.Point{}, draw.Src)

		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			for x := rect.Min.X; x < rect.Max.X; x++ {
				dstImage.Set(x, y, item.img.At(x, y))
			}
		}

		t.img = dstImage
	}
}

func (t *Template) Bytes() *[]byte {
	if t.bP != nil {
		return t.bP
	}

	var buffer bytes.Buffer
	err := png.Encode(&buffer, t.img)
	if err != nil {
		log.Fatal(err)
	}

	b := buffer.Bytes()
	t.bP = &b

	return t.bP
}

func (t *Template) Image() image.Image { return t.img }
