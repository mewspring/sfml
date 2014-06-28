// many demonstrates how to create and handle more than one window at once.
package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path"

	"github.com/mewkiz/pkg/goutil"
	"github.com/mewmew/sfml/texture"
	"github.com/mewmew/sfml/window"
	"github.com/mewmew/we"
)

// dataDir is the absolute path to the example source directory.
var dataDir string

func init() {
	// Locate the absolute path to the example source directory.
	var err error
	dataDir, err = goutil.SrcDir("github.com/mewmew/sfml/examples/data")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	err := many()
	if err != nil {
		log.Fatalln(err)
	}
}

// many demonstrates how to create and handle more than one window at once.
func many() (err error) {
	errc := make(chan error)
	go spawn("window one", errc)
	go spawn("window two", errc)
	for i := 0; i < 2; i++ {
		err = <-errc
		if err != nil {
			return err
		}
	}
	return nil
}

// spawn spawns a new window with the specified title.
func spawn(title string, errc chan error) {
	errc <- tiny(title)
}

// tiny demonstrates how to render images onto the window using the Draw and
// DrawRect methods. It also gives an example of a basic event loop.
func tiny(title string) (err error) {
	// Open a window with the specified dimensions.
	win, err := window.Open(640, 480)
	if err != nil {
		return err
	}
	defer win.Close()
	win.SetTitle(title)

	// Load background texture.
	bg, err := texture.Load(path.Join(dataDir, "bg.png"))
	if err != nil {
		return err
	}
	defer bg.Free()

	// Load foreground texture.
	fg, err := texture.Load(path.Join(dataDir, "fg.png"))
	if err != nil {
		return err
	}
	defer fg.Free()

	// Drawing and event loop.
	for {
		// Poll events until the event queue is empty.
		for e := win.PollEvent(); e != nil; e = win.PollEvent() {
			fmt.Printf("%v - %T: %v\n", title, e, e)
			switch e.(type) {
			case we.Close:
				// Close the window.
				return nil
			}
		}

		// Fill the window with white color.
		win.Fill(color.White)

		// Draw the entire background texture onto the window.
		err = win.Draw(image.ZP, bg)
		if err != nil {
			return err
		}

		// Draw a subset of the foreground texture, as defined by the source
		// rectangle (90, 90, 225, 225), onto the window starting at the
		// destination point (10, 10).
		dp := image.Pt(10, 10)
		sr := image.Rect(90, 90, 225, 225)
		err = win.DrawRect(dp, fg, sr)
		if err != nil {
			return err
		}

		// Display what has been rendered so far to the window.
		win.Display()
	}
}
