package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/ebijam24/internal/render"
)

type GameStateStart struct {
	newDudes []*Dude
	length   int
}

func (s *GameStateStart) Begin(g *Game) {
	if s.length == 0 {
		s.length = 3
	}
	// Give the player a reasonable amount of GOLD
	g.gold = 700

	professions := []ProfessionKind{Knight, Vagabond, Ranger, Cleric}
	dudeLimit := len(professions)
	for i := 0; i < dudeLimit; i++ {
		pk := professions[i%len(professions)]
		dude := NewDude(pk, 1)
		s.newDudes = append(s.newDudes, dude)
	}
	// Add some more randomized dudes.
	for i := 0; i < 3; i++ {
		pk := professions[rand.Intn(len(professions))]
		dude := NewDude(pk, 1)
		s.newDudes = append(s.newDudes, dude)
	}

	// Create a new tower, yo.
	tower := NewTower()
	for i := 0; i < 3+s.length*2; i++ {
		story := NewStory()
		if i == 0 {
			story.Open()
		}
		tower.AddStory(story)
	}

	g.tower = tower
	g.camera.SetMode(render.CameraModeTower)
}
func (s *GameStateStart) End(g *Game) {
	g.dudes = append(g.dudes, s.newDudes...)
	g.camera.SetMode(render.CameraModeStack)
}
func (s *GameStateStart) Update(g *Game) GameState {
	//return &GameStateWin{}
	return &GameStateBuild{}
}
func (s *GameStateStart) Draw(g *Game, screen *ebiten.Image) {
}
