package ui

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"
)

type Resources struct {
	titleTexture        *sdl.Texture
	onefingerTexture    *sdl.Texture
	twofingersTexture   *sdl.Texture
	threefingersTexture *sdl.Texture
	fourfingersTexture  *sdl.Texture
	fistTexture         *sdl.Texture
}

func InitSDL(players []game.PlayerI) error {
	w, r, err := createWinGame()
	if err != nil {
		return err
	}
	defer sdl.Quit()
	defer ttf.Quit()
	defer w.Destroy()
	defer r.Destroy()

	resources, err := loadResources(r)
	if err != nil {
		return err
	}
	//Draw
	r.SetDrawColor(206, 39, 39, 255)
	r.Clear()

	err = PrintTitle(r, resources)
	if err != nil {
		return err
	}

	err = printPlayers(players, r, resources)
	if err != nil {
		return err
	}

	r.Present()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return nil
			}
		}
	}

	return nil
}

func createWinGame() (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, err
	}
	//Init font
	if err := ttf.Init(); err != nil {
		return nil, nil, err
	}

	w, r, err := sdl.CreateWindowAndRenderer(1000, 800, sdl.WINDOW_OPENGL)
	if err != nil {
		return nil, nil, err
	}

	return w, r, err
}

func PrintTitle(r *sdl.Renderer, resources *Resources) error {
	dstRect := &sdl.Rect{X: 100, Y: 5, W: 800, H: 100}

	if err := r.Copy(resources.titleTexture, nil, dstRect); err != nil {
		return err
	}
	return nil
}

func printPlayers(players []game.PlayerI, r *sdl.Renderer, resources *Resources) error {
	var positionX int32 = 0
	var texture *sdl.Texture

	for _, player := range players {
		//TODO take the logic out of here and call a function just to print hands in the correct position
		hand := player.GetPlayer().LeftHand
		switch hand {
		case 1:
			{
				texture = resources.onefingerTexture
			}
		case 2:
			{
				texture = resources.twofingersTexture
			}
		case 3:
			{
				texture = resources.threefingersTexture
			}
		case 4:
			{
				texture = resources.fourfingersTexture
			}
		default:
			{
				texture = resources.fistTexture
			}

		}

		//Create one src and dst rect
		srcRect := &sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
		dstRect := &sdl.Rect{positionX + 100, 800 - 300, 370, 300}
		positionX += 400
		if err := r.Copy(texture, srcRect, dstRect); err != nil {
			return err
		}
	}
	//Create another src and dst rect
	// srcRect2 := &sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
	// dstRect2 := &sdl.Rect{500, 800 - 300, 370, 300}
	// if err := r.Copy(texture, srcRect2, dstRect2); err != nil {
	// 	return err
	// }
	return nil
}

func loadResources(r *sdl.Renderer) (*Resources, error) {
	resources := &Resources{}
	if err := loadImageTextures(resources, r); err != nil {
		return nil, err
	}
	if err := loadFontTextures(resources, r); err != nil {
		return nil, err
	}
	return resources, nil
}

func loadImageTextures(resources *Resources, r *sdl.Renderer) error {
	var err error
	if resources.onefingerTexture, err = loadPNG("ui/img/onefinger.png", r); err != nil {
		return err
	}
	if resources.twofingersTexture, err = loadPNG("ui/img/twofingers.png", r); err != nil {
		return err
	}
	if resources.threefingersTexture, err = loadPNG("ui/img/threefingers.png", r); err != nil {
		return err
	}
	if resources.fourfingersTexture, err = loadPNG("ui/img/fourfingers.png", r); err != nil {
		return err
	}
	if resources.fistTexture, err = loadPNG("ui/img/fist.png", r); err != nil {
		return err
	}
	return nil
}

func loadFontTextures(resources *Resources, r *sdl.Renderer) error {
	font, err := ttf.OpenFont("ui/font/pixeboy.ttf", 20)
	if err != nil {
		return err
	}
	defer font.Close()
	s, err := font.RenderUTF8Solid("Chopstick Game", sdl.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})
	if err != nil {
		return err
	}
	defer s.Free()
	texture, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return err
	}

	resources.titleTexture = texture
	return nil
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
