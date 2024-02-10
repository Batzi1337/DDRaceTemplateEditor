package assets

import (
	"image"
	"image/draw"
	"io"
)

type Item image.Image

type Template image.Image

func NewTemplate(r io.Reader) (Template, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	return Template(img), nil
}

func NewItem(x, y, width, height int, src Template) Item {
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), src, image.Point{x, y}, draw.Src)

	return subImage
}

func ReplaceItems(template Template, items ...Item) (result Template) {
	result = template
	for _, item := range items {
		result = replaceItem(result, item)
	}

	return
}

func replaceItem(src Template, replacement Item) Template {
	rect := replacement.Bounds()
	dstImage := image.NewRGBA(src.Bounds())
	draw.Draw(dstImage, dstImage.Bounds(), src, image.Point{}, draw.Src)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			dstImage.Set(x, y, replacement.At(x, y))
		}
	}

	return dstImage
}
