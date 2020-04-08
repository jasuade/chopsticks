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
	messageTextures   MessageTexture
}

type MessagesPrinter interface {
	printMessage(r *sdl.Renderer, resources *Resources, textId string) error
}

type MessageTexture map[string]*sdl.Texture

var Messages = []struct {
	id   string
	text string
}{
	{id: "MSG_START", text: "Start the game with %d payers in mode <%s> and level <%s>"},
	{id: "MSG_TURN", text: "Player %d is your turn"},
	{id: "MSG_HAND_TO_ATTACK", text: "With which hand (left(l) or right(r))do you want to attack:"},
	{id: "MSG_HAND_TO_ATTACK_WITH", text: "And which hand (left(l) or right(r))do you want to attack:"},
	{id: "MSG_SPLIT", text: "How many chopsticks you want to transfere?"},
	{id: "MSG_ACTION", text: "Are you goint to attack(a) or to split(s):"},
	{id: "MSG_WIN", text: "Player P%d WINS the game"},
	{id: "MSG_LOSE", text: "Player Px LOSES the game"},
	{id: "ERR_INVALID_ATTACK", text: "Invalid attack, the hand does not exist"},
	{id: "ERR_SPLIT_HAND", text: "Unable to slpit, you cannot kill a hand"},
	{id: "ERR_SPLIT", text: "That is not a meaninful split"},
	{id: "ERR_ACTION", text: "Invalid action: please introduce a valid action"},
}

func InitSDL() (*sdl.Window, *sdl.Renderer, *Resources, error) {
	w, r, err := createWinGame()
	if err != nil {
		return nil, nil, nil, err
	}

	resources, err := loadResources(r)
	if err != nil {
		return nil, nil, nil, err
	}

	r.Clear()

	err = printBackground(r, resources)
	if err != nil {
		return nil, nil, nil, err
	}

	err = printTitle(r, resources)
	if err != nil {
		return nil, nil, nil, err
	}

	i := &MessageTexture{}
	err = i.printMessage(r, resources, "ERR_SPLIT")
	if err != nil {
		return nil, nil, nil, err
	}
	return w, r, resources, nil
}

func GameLoop(w *sdl.Window, r *sdl.Renderer) error {
	r.Present()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				endGame(w, r)
				return nil
			}
		}
	}
}

func endGame(w *sdl.Window, r *sdl.Renderer) {
	sdl.Quit()
	ttf.Quit()
	w.Destroy()
	r.Destroy()
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

func printBackground(r *sdl.Renderer, resources *Resources) error {
	dstRect := &sdl.Rect{X: 0, Y: 0, W: 1000, H: 800}

	if err := r.Copy(resources.backgroundTexture, nil, dstRect); err != nil {
		return err
	}
	return nil
}

func printTitle(r *sdl.Renderer, resources *Resources) error {
	dstRect := &sdl.Rect{X: 100, Y: 5, W: 800, H: 100}

	if err := r.Copy(resources.titleTexture, nil, dstRect); err != nil {
		return err
	}
	return nil
}

func (m *MessageTexture) printMessage(r *sdl.Renderer, resources *Resources, textId string) error {
	dstRect := &sdl.Rect{X: (1000 - 700) / 2, Y: 430, W: 700, H: 30}

	if err := r.Copy(resources.messageTextures[textId], nil, dstRect); err != nil {
		return err
	}
	return nil
}

func PrintPlayers(players []game.PlayerI, r *sdl.Renderer, resources *Resources) error {
	var positionX int32 = 100
	var positionY int32 = 800 - 300
	var angle float64 = 0

	for _, player := range players {
		//TODO take the logic out of here and call a function just to print hands in the correct position
		leftHand := player.GetPlayer().LeftHand
		rightHand := player.GetPlayer().RightHand
		textureLeftHand := resources.handTextures[leftHand]
		textureRightHand := resources.handTextures[rightHand]

		srcRect := &sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
		dstRect := &sdl.Rect{X: positionX, Y: positionY, W: 370, H: 300}
		if err := r.CopyEx(textureLeftHand, srcRect, dstRect, angle, nil, 0); err != nil {
			return err
		}

		srcRect = &sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
		dstRect = &sdl.Rect{X: positionX + 400, Y: positionY, W: 370, H: 300}
		// CopyEx(texture *Texture, src, dst *Rect, angle float64, center *Point, flip RendererFlip) error
		if err := r.CopyEx(textureRightHand, srcRect, dstRect, angle, nil, 5); err != nil {
			return err
		}
		angle += 180
		positionY -= 400
	}
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
	if err := loadTextsTextures(resources, r); err != nil {
		return nil, err
	}
	if err := loadBackground(resources, r); err != nil {
		return nil, err
	}
	return resources, nil
}

func loadTextsTextures(resources *Resources, r *sdl.Renderer) error {
	resources.messageTextures = make(map[string]*sdl.Texture)
	font, err := ttf.OpenFont("ui/font/pixeboy.ttf", 10)
	if err != nil {
		return err
	}
	defer font.Close()
	for i := 0; i < len(Messages); i++ {
		s, err := font.RenderUTF8Solid(Messages[i].text, sdl.Color{
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
		resources.messageTextures[Messages[i].id] = texture
	}
	return nil
}

func loadBackground(resources *Resources, r *sdl.Renderer) error {
	var err error
	if resources.backgroundTexture, err = loadPNG("ui/img/background.png", r); err != nil {
		return err
	}
	return nil
}

func loadImageTextures(resources *Resources, r *sdl.Renderer) error {
	for i := 0; i < 5; i++ {
		path := fmt.Sprintf("ui/img/%d_finger.png", i)
		texture, err := loadPNG(path, r)
		if err != nil {
			return err
		}
		resources.handTextures = append(resources.handTextures, texture)
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
