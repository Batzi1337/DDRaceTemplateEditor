package models

import (
	"image"
	"image/draw"
)

type Item struct {
	startPoint    image.Point
	width, height int
	Img           image.RGBA64Image
}

func create(t *Template, startPoint image.Point, width, height int) *Item {
	item := &Item{
		startPoint: startPoint,
		width:      width,
		height:     height,
	}
	setImage(item, t)

	return item
}

func setImage(item *Item, t *Template) {
	rect := image.Rect(item.startPoint.X, item.startPoint.Y, item.startPoint.X+item.width, item.startPoint.Y+item.height)
	subImage := image.NewRGBA(rect)
	draw.Draw(subImage, subImage.Bounds(), t.Img, image.Point{item.startPoint.X, item.startPoint.Y}, draw.Src)
	item.Img = subImage
}

func Hammer(t *Template) *Item {
	return create(t, image.Point{64, 32}, 128, 96)
}

func Shotgun(t *Template) *Item {
	return create(t, image.Point{64, 192}, 256, 64)
}

func ShotgunCrosshair(t *Template) *Item {
	return create(t, image.Point{0, 192}, 64, 64)
}

func ShotgunBullet(t *Template) *Item {
	return create(t, image.Point{320, 192}, 64, 64)
}

func Pistol(t *Template) *Item {
	return create(t, image.Point{64, 128}, 128, 64)
}

func PistolCrosshair(t *Template) *Item {
	return create(t, image.Point{0, 128}, 64, 64)
}

func PistolBullet(t *Template) *Item {
	return create(t, image.Point{192, 128}, 64, 64)
}

func Sword(t *Template) *Item {
	return create(t, image.Point{64, 320}, 256, 64)
}
