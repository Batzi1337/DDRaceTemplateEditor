package main

import (
	"flag"
	"fmt"
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

	// Load images
	template1, err := NewTemplate(*srcFile)
	if err != nil {
		fmt.Println("Error loading image1:", err)
		return
	}

	template2, err := NewTemplate(*dstFile)
	if err != nil {
		fmt.Println("Error loading image2:", err)
		return
	}

	items := createItems(template1, *itemTypes)
	if items == nil {
		return
	}

	replaceItems(template2, items)

	err = template2.Save(*outputfile)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Operation completed successfully. Result saved to " + *outputfile)
}

func createItems(template *Template, itemTypes string) []*Item {
	items := []*Item{}
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

func createItem(template *Template, itemType string) *Item {
	switch itemType {
	case "hammer":
		return NewHammer(template)
	case "sword":
		return NewSword(template)
	case "shotgun":
		return NewShotgun(template)
	case "shotgun_crshr":
		return NewShotgunCrosshair(template)
	case "shotgun_bllt":
		return NewShotgunBullet(template)
	case "pistol":
		return NewPistol(template)
	case "pistol_crshr":
		return NewPistolCrosshair(template)
	case "pistol_bllt":
		return NewPistolBullet(template)
	default:
		fmt.Println("Error: unknown item type '" + itemType + "'")
		return nil
	}
}

func replaceItems(template *Template, items []*Item) {
	for _, item := range items {
		template.ReplaceItem(item)
	}
}
