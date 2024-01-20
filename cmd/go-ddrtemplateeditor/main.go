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

	items := []*Item{}
	for _, itemType := range strings.Split(*itemTypes, ",") {

		switch itemType {
		case "hammer":
			hammer := NewHammer(template1)
			items = append(items, hammer)
			fmt.Println("Replacing hammer")
		case "sword":
			sword := NewSword(template1)
			items = append(items, sword)
			fmt.Println("Replacing sword")
		case "shotgun":
			shotgun := NewShotgun(template1)
			items = append(items, shotgun)
			fmt.Println("Replacing shotgun")
		case "shotgun_crshr":
			shotgunCrshr := NewShotgunCrosshair(template1)
			items = append(items, shotgunCrshr)
			fmt.Println("Replacing shotgun crosshair")
		case "shotgun_bllt":
			shotgunBllt := NewShotgunBullet(template1)
			items = append(items, shotgunBllt)
			fmt.Println("Replacing shotgun bullet")
		case "pistol":
			pistol := NewPistol(template1)
			items = append(items, pistol)
			fmt.Println("Replacing pistol")
		case "pistol_crshr":
			pistolCrshr := NewPistolCrosshair(template1)
			items = append(items, pistolCrshr)
			fmt.Println("Replacing pistol crosshair")
		case "pistol_bllt":
			pistolBllt := NewPistolBullet(template1)
			items = append(items, pistolBllt)
			fmt.Println("Replacing pistol bullet")
		default:
			fmt.Println("Error: unknown item type '" + itemType + "'")
			return
		}
	}

	// Replace the items
	for _, item := range items {
		template2.ReplaceItem(item)
	}

	// Save the result
	err = template2.Save(*outputfile)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Operation completed successfully. Result saved to " + *outputfile)
}
