package models

import "fmt"

type App struct {
	StudentID     string
	AssignmentID  string
	TimeSubmitted int64
	Sprites       []*Sprite
	Original      string
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
