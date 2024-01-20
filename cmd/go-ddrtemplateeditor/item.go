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

func NewHammer(t *Template) *Item {
	hmmr := &Item{
		startPoint:  image.Point{64, 32},
		width:       128,
		height:      96,
		srcTemplate: t,
	}
	SetImage(hmmr)

	return hmmr
}

func NewShotgun(t *Template) *Item {
	shotgun := &Item{
		startPoint:  image.Point{64, 192},
		width:       256,
		height:      64,
		srcTemplate: t,
	}
	SetImage(shotgun)

	return shotgun

}

func NewShotgunCrosshair(t *Template) *Item {
	crshr := &Item{
		startPoint:  image.Point{0, 192},
		width:       64,
		height:      64,
		srcTemplate: t,
	}
	SetImage(crshr)

	return crshr
}

func NewPistol(t *Template) *Item {
	pstl := &Item{
		startPoint:  image.Point{64, 128},
		width:       128,
		height:      64,
		srcTemplate: t,
	}
	SetImage(pstl)

	return pstl
}

func NewPistolCrosshair(t *Template) *Item {
	crshr := &Item{
		startPoint:  image.Point{0, 128},
		width:       64,
		height:      64,
		srcTemplate: t,
	}
	SetImage(crshr)

	return crshr
}

func NewPistolBullet(t *Template) *Item {
	bllt := &Item{
		startPoint:  image.Point{192, 128},
		width:       64,
		height:      64,
		srcTemplate: t,
	}
	SetImage(bllt)

	return bllt
}

func NewSword(t *Template) *Item {
	sword := &Item{
		startPoint:  image.Point{64, 320},
		width:       256,
		height:      64,
		srcTemplate: t,
	}
	SetImage(sword)

	return sword
}

func SetImage(item *Item) {
	rect := image.Rect(item.startPoint.X, item.startPoint.Y, item.startPoint.X+item.width, item.startPoint.Y+item.height)
	subImage := image.NewRGBA(rect)
	draw.Draw(subImage, subImage.Bounds(), item.srcTemplate.img, image.Point{item.startPoint.X, item.startPoint.Y}, draw.Src)
	item.img = subImage
}
