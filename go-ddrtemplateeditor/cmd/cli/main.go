package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"models"
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

	replaceItems(template2, items)

	err = saveTemplateToFile(template2, *outputfile)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Operation completed successfully. Result saved to " + *outputfile)
}

func createItems(template *models.Template, itemTypes string) []*models.Item {
	items := []*models.Item{}
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

func createItem(template *models.Template, itemType string) *models.Item {
	switch itemType {
	case "hammer":
		return models.Hammer(template)
	case "sword":
		return models.Sword(template)
	case "shotgun":
		return models.Shotgun(template)
	case "shotgun_crshr":
		return models.ShotgunCrosshair(template)
	case "shotgun_bllt":
		return models.ShotgunBullet(template)
	case "pistol":
		return models.Pistol(template)
	case "pistol_crshr":
		return models.PistolCrosshair(template)
	case "pistol_bllt":
		return models.PistolBullet(template)
	default:
		fmt.Println("Error: unknown item type '" + itemType + "'")
		return nil
	}
}

func replaceItems(template *models.Template, items []*models.Item) {
	for _, item := range items {
		template.ReplaceItem(item)
	}
}

func loadTemplateFromFile(path string) (*models.Template, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return &models.Template{Img: img}, nil
}

func saveTemplateToFile(template *models.Template, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, template.Img)
}
