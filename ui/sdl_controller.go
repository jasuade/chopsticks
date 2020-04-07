package ui

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func InitSDL() error {
	w, r, err := createWinGame()
	if err != nil {
		return err
	}
	defer sdl.Quit()
	defer w.Destroy()
	defer r.Destroy()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return nil
			}
		}
	}
}

func createWinGame() (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, err
	}
	// func CreateWindowAndRenderer(w, h int32, flags uint32) (*Window, *Renderer, error)
	w, r, err := sdl.CreateWindowAndRenderer(1000, 800, sdl.WINDOW_OPENGL)
	if err != nil {
		return nil, nil, err
	}
	texture, err := loadPNG("ui/img/fourfingers.png", r)
	if err != nil {
		return nil, nil, err
	}
	//Draw
	r.SetDrawColor(206, 39, 39, 255)
	r.Clear()
	//Create one src and dst rect
	srcRect := &sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
	dstRect := &sdl.Rect{200, 800 - 300, 370, 300}
	if err := r.Copy(texture, srcRect, dstRect); err != nil {
		return nil, nil, err
	}

	//Create another src and dst rect
	srcRect2 := &sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
	dstRect2 := &sdl.Rect{500, 800 - 300, 370, 300}
	if err := r.Copy(texture, srcRect2, dstRect2); err != nil {
		return nil, nil, err
	}

	r.Present()
	return w, r, nil
}

func loadPNG(filePath string, r *sdl.Renderer) (*sdl.Texture, error) {
	file := sdl.RWFromFile(filePath, "rb")
	defer file.Close()
	img, err := img.LoadPNGRW(file)
	defer img.Free()
	if err != nil {
		return nil, err
	}
	sdlTexture, err := r.CreateTextureFromSurface(img)
	if err != nil {
		return nil, err
	}
	return sdlTexture, nil
}

// func InitSDL() error {
// 	err := sdl.Init(sdl.INIT_EVERYTHING)
// 	if err != nil {
// 		fmt.Printf("Error Initializing sdl %v", err)
// 		return err
// 	}
// 	defer sdl.Quit()
// 	w, _, err := sdl.CreateWindowAndRenderer(3000, 3000, sdl.WINDOW_SHOWN)
// 	if err != nil {
// 		fmt.Printf("Error creating window %v", err)
// 		return err
// 	}
// 	defer w.Destroy()

// 	return nil
// }
