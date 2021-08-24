// fonts demonstrates how to render text using TTF fonts.
package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/mewkiz/pkg/goutil"
	"github.com/mewspring/sfml/font"
	"github.com/mewspring/sfml/texture"
	"github.com/mewspring/sfml/window"
	"github.com/mewspring/we"
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
	err := fonts()
	if err != nil {
		log.Fatalln(err)
	}
}

// fonts demonstrates how to render text using TTF fonts.
func fonts() (err error) {
	// Some operating systems require that the main thread is used for both
	// window creation and event handling. Therefore we lock the goroutine to an
	// OS thread.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Open a window with the specified dimensions.
	win, err := window.Open(640, 480)
	if err != nil {
		return err
	}
	defer win.Close()

	// Load background texture.
	bg, err := texture.Load(path.Join(dataDir, "bg2.png"))
	if err != nil {
		return err
	}
	defer bg.Free()

	// Load the text TTF font.
	textFont, err := font.Load(path.Join(dataDir, "Exocet.ttf"))
	if err != nil {
		return err
	}
	defer textFont.Free()

	// Create a new graphical text entry based on the Exocet TTF font and
	// initialize its text to "TTF fonts", its font size to 32 (the default is
	// 12) and its color to white (the default is black).
	text, err := font.NewText(textFont, "TTF fonts", 32, color.White)
	if err != nil {
		return err
	}
	defer text.Free()

	// Load the fps TTF font.
	fpsFont, err := font.Load(path.Join(dataDir, "DejaVuSansMono.ttf"))
	if err != nil {
		return err
	}
	defer fpsFont.Free()

	// Create a graphical FPS text entry. The text of this graphical text entry
	// will be updated repeatedly using SetText.
	fps, err := font.NewText(fpsFont, 14, color.White)
	if err != nil {
		return err
	}
	defer fps.Free()

	// start and frames will be used to calculate the average FPS of the
	// application.
	start := time.Now()
	frames := 0.0

	// 60 FPS
	ticker := time.NewTicker(time.Second / 60)

	// Drawing and event loop.
	for {
		// Cap the FPS.
		<-ticker.C

		// Fill the window with white color.
		win.Fill(color.White)

		// Draw the entire background texture onto the window.
		err = win.Draw(image.ZP, bg)
		if err != nil {
			return err
		}

		// Draw the entire text onto the window starting the destination point
		// (420, 12).
		dp := image.Pt(420, 12)
		err = win.Draw(dp, text)
		if err != nil {
			return err
		}

		// Update the text of the FPS text entry.
		fps.SetText(getFPS(start, frames))

		// Draw the entire FPS text entry onto the screen starting at the
		// destination point (8, 4).
		dp = image.Pt(8, 4)
		err = win.Draw(dp, fps)
		if err != nil {
			return err
		}

		// Display what has been rendered so far to the window.
		win.Display()
		frames++

		// Poll events until the event queue is empty.
		for e := win.PollEvent(); e != nil; e = win.PollEvent() {
			fmt.Printf("%T: %v\n", e, e)
			switch e.(type) {
			case we.Close:
				// Close the window.
				return nil
			}
		}
	}
}

// getFPS returns the average FPS as a string, based on the provided start time
// and frame count.
func getFPS(start time.Time, frames float64) (text string) {
	// Average FPS.
	fps := frames / time.Since(start).Seconds()
	return fmt.Sprintf("FPS: %.2f", fps)
}
