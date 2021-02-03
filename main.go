package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}

}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("Couldn't initialize sdl: %v", err)
	}
	defer sdl.Quit()

	err = ttf.Init()
	if err != nil {
		return fmt.Errorf("Couldn't initialize ttf: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("Couldn't create a window: %v", err)
	}
	defer w.Destroy()

	err = drawText(r)
	if err != nil {
		return fmt.Errorf("Couldn't draw the texture: %v", err)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func drawText(r *sdl.Renderer) error {
	r.Clear()

	f, err := ttf.OpenFont("resource/fonts/FunnyTeca.ttf", 20)
	if err != nil {
		return fmt.Errorf("Couldn't open the Font: %v", err)
	}
	defer f.Close()

	s, err := f.RenderUTF8Solid("Go Flappy", sdl.Color{R: 255, G: 50, B: 120, A: 255})
	if err != nil {
		return fmt.Errorf("Couldn't render the Font: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("Couldn't create the texture: %v", err)
	}
	defer t.Destroy()

	err = r.Copy(t, nil, nil)
	if err != nil {
		return fmt.Errorf("Couldn't copy the texture: %v", err)
	}

	r.Present()
	return nil
}
