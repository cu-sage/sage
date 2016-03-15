package models

import "fmt"

type App struct {
	Sprites []*Sprite
}

func NewApp() *App {
	return &App{
		Sprites: []*Sprite{},
	}
}

func (a *App) AddSprite(s *Sprite) {
	a.Sprites = append(a.Sprites, s)
}

func (a *App) Log() {
	for _, sprite := range a.Sprites {
		sprite.Log()
		fmt.Println()
	}
}
