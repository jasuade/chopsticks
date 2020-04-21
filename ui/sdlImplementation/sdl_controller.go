package ui

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"
)

type Resources struct {
	titleTexture      *sdl.Texture
	handTextures      []*sdl.Texture
	backgroundTexture *sdl.Texture
	fontMessage       *ttf.Font
}

type SDLController struct {
	Renderer  *sdl.Renderer
	window    *sdl.Window
	Resources *Resources
}

func InitSDL() (*SDLController, error) {
	w, r, err := createWinGame()
	if err != nil {
		return nil, err
	}

	resources, err := loadResources(r)
	if err != nil {
		return nil, err
	}

	sdlController := &SDLController{r, w, resources}
	return sdlController, nil
}

func createWinGame() (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, err
	}

	if err := ttf.Init(); err != nil {
		return nil, nil, err
	}

	w, r, err := sdl.CreateWindowAndRenderer(1000, 800, sdl.WINDOW_OPENGL)
	if err != nil {
		return nil, nil, err
	}

	return w, r, err
}

func (controller *SDLController) EndGame() {
	controller.Resources.fontMessage.Close()
	sdl.Quit()
	ttf.Quit()
	controller.window.Destroy()
	controller.Renderer.Destroy()
}

func (controller *SDLController) PrintBackground() error {
	dstRect := &sdl.Rect{X: 0, Y: 0, W: 1000, H: 800}

	if err := controller.Renderer.Copy(controller.Resources.backgroundTexture, nil, dstRect); err != nil {
		return err
	}
	return nil
}

func (controller *SDLController) PrintTitle() error {
	dstRect := &sdl.Rect{X: 100, Y: 5, W: 800, H: 100}

	if err := controller.Renderer.Copy(controller.Resources.titleTexture, nil, dstRect); err != nil {
		return err
	}
	return nil
}

func (controller *SDLController) PrintMessage(text string) error {

	surface, err := controller.Resources.fontMessage.RenderUTF8Solid(text, sdl.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})
	if err != nil {
		return err
	}
	defer surface.Free()
	messageTexture, err := controller.Renderer.CreateTextureFromSurface(surface)
	dstRect := &sdl.Rect{X: (1000 - 700) / 2, Y: 430, W: 700, H: 30}

	if err := controller.Renderer.Copy(messageTexture, nil, dstRect); err != nil {
		return err
	}
	return nil
}

func (controller *SDLController) PrintPlayers(players []game.PlayerI) error {
	var positionX int32 = 100
	var positionY int32 = 800 - 300
	var angle float64 = 0
	srcRectLeft := &sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
	srcRectRight := &sdl.Rect{X: 0, Y: 0, W: 512, H: 512}

	for _, player := range players {
		leftHand := player.GetPlayer().LeftHand
		rightHand := player.GetPlayer().RightHand
		textureLeftHand := controller.Resources.handTextures[leftHand]
		textureRightHand := controller.Resources.handTextures[rightHand]

		if leftHand == 0 {
			srcRectLeft = &sdl.Rect{X: 0, Y: 0, W: 1600, H: 1600}
		}
		if rightHand == 0 {
			srcRectRight = &sdl.Rect{X: 0, Y: 0, W: 1600, H: 1600}
		}

		dstRect := &sdl.Rect{X: positionX, Y: positionY, W: 370, H: 300}
		if err := controller.Renderer.CopyEx(textureLeftHand, srcRectLeft, dstRect, angle, nil, 0); err != nil {
			return err
		}

		dstRect = &sdl.Rect{X: positionX + 400, Y: positionY, W: 370, H: 300}
		if err := controller.Renderer.CopyEx(textureRightHand, srcRectRight, dstRect, angle, nil, 5); err != nil {
			return err
		}
		angle += 180
		positionY -= 400
	}
	return nil
}

func loadResources(renderer *sdl.Renderer) (*Resources, error) {
	resources := &Resources{}
	if err := loadImageTextures(resources, renderer); err != nil {
		return nil, err
	}
	if err := loadFontTextures(resources, renderer); err != nil {
		return nil, err
	}
	if err := loadBackground(resources, renderer); err != nil {
		return nil, err
	}
	return resources, nil
}

func loadBackground(resources *Resources, r *sdl.Renderer) error {
	var err error
	if resources.backgroundTexture, err = loadPNG("ui/sdlImplementation/img/background.png", r); err != nil {
		return err
	}
	return nil
}

func loadImageTextures(resources *Resources, r *sdl.Renderer) error {
	for i := 0; i < 5; i++ {
		path := fmt.Sprintf("ui/sdlImplementation/img/%d_finger.png", i)
		texture, err := loadPNG(path, r)
		if err != nil {
			return err
		}
		resources.handTextures = append(resources.handTextures, texture)
	}
	return nil
}

func loadFontTextures(resources *Resources, r *sdl.Renderer) error {
	font, err := ttf.OpenFont("ui/sdlImplementation/font/pixeboy.ttf", 20)
	if err != nil {
		return err
	}
	resources.fontMessage = font
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
