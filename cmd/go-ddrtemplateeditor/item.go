package main

import (
	"image"
	"image/draw"
)

type Replacement interface {
	Image() image.RGBA64Image
	StartPoint() *image.Point
}

type Item struct {
	startPoint    image.Point
	width, height int
	img           image.RGBA64Image
	srcTemplate   *Template
}

func (h *Item) Image() image.RGBA64Image {
	return h.img
}

func (h *Item) StartPoint() *image.Point {
	return &h.startPoint
}

func NewItem(t *Template, startPoint image.Point, width, height int) *Item {
	item := &Item{
		startPoint:  startPoint,
		width:       width,
		height:      height,
		srcTemplate: t,
	}
	SetImage(item)

	return item
}

func SetImage(item *Item) {
	rect := image.Rect(item.startPoint.X, item.startPoint.Y, item.startPoint.X+item.width, item.startPoint.Y+item.height)
	subImage := image.NewRGBA(rect)
	draw.Draw(subImage, subImage.Bounds(), item.srcTemplate.img, image.Point{item.startPoint.X, item.startPoint.Y}, draw.Src)
	item.img = subImage
}

func NewHammer(t *Template) *Item {
	return NewItem(t, image.Point{64, 32}, 128, 96)
}

func NewShotgun(t *Template) *Item {
	return NewItem(t, image.Point{64, 192}, 256, 64)
}

func NewShotgunCrosshair(t *Template) *Item {
	return NewItem(t, image.Point{0, 192}, 64, 64)
}

func NewShotgunBullet(t *Template) *Item {
	return NewItem(t, image.Point{320, 192}, 64, 64)
}

func NewPistol(t *Template) *Item {
	return NewItem(t, image.Point{64, 128}, 128, 64)
}

func NewPistolCrosshair(t *Template) *Item {
	return NewItem(t, image.Point{0, 128}, 64, 64)
}

func NewPistolBullet(t *Template) *Item {
	return NewItem(t, image.Point{192, 128}, 64, 64)
}

func NewSword(t *Template) *Item {
	return NewItem(t, image.Point{64, 320}, 256, 64)
}
