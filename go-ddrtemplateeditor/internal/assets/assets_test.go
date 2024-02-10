package assets

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceItems(t *testing.T) {
	srcImage := getImage("../../samples/templates/purple_haze.png")
	replacementImage := getImage("../../samples/items/item1.png")

	src := Template(srcImage)
	replacement := Item(replacementImage)

	result := ReplaceItems(src, replacement)

	// Verify that the replaced item is present in the result
	assert.True(t, isItemPresent(result, replacement))
}

func isItemPresent(src Template, item Item) bool {
	// Convert the images to RGBA for pixel-level comparison
	srcRGBA := convertToRGBA(src)
	itemRGBA := convertToRGBA(item)

	// Get the dimensions of the item
	itemBounds := itemRGBA.Bounds()

	// Iterate over the pixels of the source image
	for y := srcRGBA.Bounds().Min.Y; y < srcRGBA.Bounds().Max.Y; y++ {
		for x := srcRGBA.Bounds().Min.X; x < srcRGBA.Bounds().Max.X; x++ {
			// Check if the current pixel matches the item
			if matchItem(srcRGBA, itemRGBA, x, y, itemBounds) {
				return true
			}
		}
	}

	return false
}

func matchItem(src, item *image.RGBA, x, y int, itemBounds image.Rectangle) bool {
	// Iterate over the pixels of the item
	for itemY := itemBounds.Min.Y; itemY < itemBounds.Max.Y; itemY++ {
		for itemX := itemBounds.Min.X; itemX < itemBounds.Max.X; itemX++ {
			// Get the corresponding pixel in the source image
			srcPixel := src.RGBAAt(x+itemX, y+itemY)
			// Get the corresponding pixel in the item image
			itemPixel := item.RGBAAt(itemX, itemY)

			// Compare the RGB values of the pixels
			if srcPixel.R != itemPixel.R || srcPixel.G != itemPixel.G || srcPixel.B != itemPixel.B {
				return false
			}
		}
	}

	return true
}

func convertToRGBA(img image.Image) *image.RGBA {
	// Create a new RGBA image with the same bounds as the input image
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	// Draw the input image onto the RGBA image
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	return rgba
}

func getImage(filepath string) image.Image {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	return img
}
