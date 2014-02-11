// off-screen demonstrates how to perform hardware accelerated off-screen
// rendering.
package main

import (
	"image"
	"log"
	"path"

	"github.com/mewkiz/pkg/goutil"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewmew/sfml/texture"
)

func main() {
	err := render()
	if err != nil {
		log.Fatalln(err)
	}
}

// render demonstrates how to perform hardware accelerated off-screen rendering.
func render() (err error) {
	srcDir, err := goutil.SrcDir("github.com/mewmew/sfml/examples/off-screen")
	if err != nil {
		return err
	}

	// Load background texture.
	bg, err := texture.Load(path.Join(srcDir, "data/bg.png"))
	if err != nil {
		return err
	}
	defer bg.Free()

	// Load foreground texture.
	fg, err := texture.Load(path.Join(srcDir, "data/fg.png"))
	if err != nil {
		return err
	}
	defer fg.Free()

	// DrawRect draws a subset of the foreground texture, as defined by the
	// source rectangle (90, 90, 250, 250), onto the background texture starting
	// at the destination point (20, 90).
	dp := image.Pt(20, 90)
	sr := image.Rect(90, 90, 250, 250)
	err = bg.DrawRect(dp, fg, sr)
	if err != nil {
		return err
	}

	// Convert the background texture into an image.Image and write it to a file.
	result, err := bg.Image()
	if err != nil {
		return err
	}
	err = imgutil.WriteFile("result.png", result)
	if err != nil {
		return err
	}

	return nil
}
