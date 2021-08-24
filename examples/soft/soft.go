// soft demonstrates how to combine software and hardware rendering.
package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path"

	"github.com/mewkiz/pkg/goutil"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewspring/sfml/texture"
)

// dataDir is the absolute path to the example source directory.
var dataDir string

func init() {
	// Locate the absolute path to the example source directory.
	var err error
	dataDir, err = goutil.SrcDir("github.com/mewspring/sfml/examples/data")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	err := soft()
	if err != nil {
		log.Fatalln(err)
	}
}

// soft demonstrates how to combine software and hardware rendering.
func soft() (err error) {
	// Create a native Go image.
	softBorder := image.NewRGBA(image.Rect(0, 0, 640, 480))

	// Draw a border using software rendering.
	rect := softBorder.Rect
	for x := rect.Min.X; x < rect.Max.X; x++ {
		softBorder.Set(x, rect.Min.Y, color.Black)
		softBorder.Set(x, rect.Max.Y-1, color.Black)
	}
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		softBorder.Set(rect.Min.X, y, color.Black)
		softBorder.Set(rect.Max.X-1, y, color.Black)
	}

	// Load background texture.
	bg, err := texture.LoadDrawable(path.Join(dataDir, "bg.png"))
	if err != nil {
		return err
	}
	defer bg.Free()

	// Convert the Go image to a GPU texture.
	border, err := texture.Read(softBorder)
	if err != nil {
		return err
	}
	defer border.Free()

	// Draw the border onto the background.
	bg.Draw(image.ZP, border)

	// Convert the background texture into an image.Image and write it to a file.
	result, err := bg.Image()
	if err != nil {
		return err
	}
	outPath := "result.png"
	err = imgutil.WriteFile(outPath, result)
	if err != nil {
		return err
	}
	fmt.Println("Created:", outPath)

	return nil
}
