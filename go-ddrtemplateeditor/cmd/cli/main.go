package main

import (
	"flag"
	"fmt"
	"go-ddrtemplateeditor/internal/assets"
	"image"
	"image/png"
	"os"
	"strings"
)

func main() {
	// Parse command line arguments
	flag.Usage = func() {
		fmt.Println("go-ddrtemplateeditor is a tool to replace items in a DDR template image.")
		fmt.Println("Example: go-ddrtemplateeditor -item hammer,sword -src template1.png -dst template2.png -out output.png")
		flag.PrintDefaults()
	}
	itemTypes := flag.String("item", "", "Use -item <hammer|sword|shotgun|shotgun_crshr|shotgun_bllt|pistol|pistol_crshr|pistol_bllt> to set the comma separated items to replace")
	srcFile := flag.String("src", "", "Use -src <path_to_template_png> to set the source template file")
	dstFile := flag.String("dst", "", "Use -dst <path_to_template_png> to set the destination template file")
	outputfile := flag.String("out", "new_template.png", "Use -out <path_to_output_png> to set the output file")
	flag.Parse()

	template1, err := loadTemplateFromFile(*srcFile)
	if err != nil {
		fmt.Println("Error loading image1:", err)
		return
	}

	template2, err := loadTemplateFromFile(*dstFile)
	if err != nil {
		fmt.Println("Error loading image2:", err)
		return
	}

	items := createItems(template1, *itemTypes)
	if items == nil {
		return
	}

	template2 = assets.ReplaceItems(template2, items...)

	err = saveTemplateToFile(template2, *outputfile)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Operation completed successfully. Result saved to " + *outputfile)
}

func createItems(template assets.Template, itemTypes string) []assets.Item {
	items := []assets.Item{}
	for _, itemType := range strings.Split(itemTypes, ",") {
		item := createItem(template, itemType)
		if item == nil {
			return nil
		}
		items = append(items, item)
		fmt.Println("Replacing", itemType)
	}
	return items
}

func createItem(template assets.Template, itemType string) assets.Item {
	switch itemType {
	case "hammer":
		return hammer(template)
	case "sword":
		return sword(template)
	case "shotgun":
		return shotgun(template)
	case "shotgun_crshr":
		return shotgunCrosshair(template)
	case "shotgun_bllt":
		return shotgunBullet(template)
	case "pistol":
		return pistol(template)
	case "pistol_crshr":
		return pistolCrosshair(template)
	case "pistol_bllt":
		return pistolBullet(template)
	default:
		fmt.Println("Error: unknown item type '" + itemType + "'")
		return nil
	}
}

func hammer(t assets.Template) assets.Item {
	x, y := 64, 32
	width, height := 128, 96
	return assets.NewItem(x, y, width, height, t)
}

func shotgun(t assets.Template) assets.Item {
	x, y := 64, 192
	width, height := 256, 64
	return assets.NewItem(x, y, width, height, t)
}

func shotgunCrosshair(t assets.Template) assets.Item {
	x, y := 0, 192
	width, height := 64, 64
	return assets.NewItem(x, y, width, height, t)
}

func shotgunBullet(t assets.Template) assets.Item {
	x, y := 320, 192
	width, height := 64, 64
	return assets.NewItem(x, y, width, height, t)
}

func sword(t assets.Template) assets.Item {
	x, y := 64, 320
	width, height := 256, 64
	return assets.NewItem(x, y, width, height, t)
}

func pistol(t assets.Template) assets.Item {
	x, y := 64, 128
	width, height := 128, 64
	return assets.NewItem(x, y, width, height, t)
}

func pistolCrosshair(t assets.Template) assets.Item {
	x, y := 0, 128
	width, height := 64, 64
	return assets.NewItem(x, y, width, height, t)
}

func pistolBullet(t assets.Template) assets.Item {
	x, y := 192, 128
	width, height := 64, 64
	return assets.NewItem(x, y, width, height, t)
}

func loadTemplateFromFile(path string) (assets.Template, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return assets.Template(img), nil
}

func saveTemplateToFile(template assets.Template, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, template)
}
