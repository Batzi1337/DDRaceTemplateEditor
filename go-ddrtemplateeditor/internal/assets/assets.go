package assets

import (
	"image"
	"image/draw"
)

type Item image.Image

type Template image.Image

func ReplaceItem(src Template, replacement Item) Template {
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

func Hammer(t Template) Item {
	x, y := 64, 32
	width, height := 128, 96
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t, image.Point{x, y}, draw.Src)

	return subImage
}

func Shotgun(t Template) Item {
	x, y := 64, 192
	width, height := 256, 64
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t, image.Point{x, y}, draw.Src)

	return subImage
}

func ShotgunCrosshair(t Template) Item {
	x, y := 0, 192
	width, height := 64, 64
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t, image.Point{x, y}, draw.Src)

	return subImage
}

func ShotgunBullet(t Template) Item {
	x, y := 320, 192
	width, height := 64, 64
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t, image.Point{x, y}, draw.Src)

	return subImage
}

func Sword(t Template) Item {
	x, y := 64, 320
	width, height := 256, 64
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t, image.Point{x, y}, draw.Src)

	return subImage
}

func Pistol(t Template) Item {
	x, y := 64, 128
	width, height := 128, 64
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t, image.Point{x, y}, draw.Src)

	return subImage
}

func PistolCrosshair(t Template) Item {
	x, y := 0, 128
	width, height := 64, 64
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t, image.Point{x, y}, draw.Src)

	return subImage
}

func PistolBullet(t Template) Item {
	x, y := 192, 128
	width, height := 64, 64
	rect := image.Rect(x, y, x+width, y+height)
	subImage := image.NewRGBA64(rect)
	draw.Draw(subImage, subImage.Bounds(), t, image.Point{x, y}, draw.Src)

	return subImage
}
