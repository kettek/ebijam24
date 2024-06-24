package render

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/ebijam24/assets"
)

type Sprite struct {
	Positionable
	Originable
	Rotateable
	image *ebiten.Image
	Scale float64
}

func (s *Sprite) Size() (float64, float64) {
	return float64(s.image.Bounds().Dx()) * s.Scale, float64(s.image.Bounds().Dy()) * s.Scale
}

func NewSprite(name string) (*Sprite, error) {
	dataSprite, err := assets.LoadSprite(name)
	if err != nil {
		return nil, err
	}
	sprite := &Sprite{
		Scale: 1,
	}
	sprite.image = dataSprite.Image
	return sprite, nil
}

func NewSubSprite(dataSprite *Sprite, x, y, w, h int) (*Sprite, error) {
	sprite := &Sprite{
		Scale: 1,
	}
	sprite.image = dataSprite.image.SubImage(image.Rect(x, y, x+w, y+h)).(*ebiten.Image)
	return sprite, nil
}

func (s *Sprite) Draw(o *Options) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(s.Origin())
	opts.GeoM.Translate(s.Position())

	ox, oy := s.Origin()
	opts.GeoM.Translate(-ox, -oy)
	opts.GeoM.Rotate(s.Rotation())
	opts.GeoM.Translate(ox, oy)

	opts.GeoM.Scale(s.Scale, s.Scale)

	opts.GeoM.Concat(o.DrawImageOptions.GeoM)

	o.Screen.DrawImage(s.image, opts)
}

func (s *Sprite) Update() {
	// ???
}
