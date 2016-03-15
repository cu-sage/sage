package models

import "log"

type Sprite struct {
	Name    string
	Scripts []string
}

func NewSprite(name string) *Sprite {
	return &Sprite{
		Name:    name,
		Scripts: []string{},
	}
}

func (s *Sprite) AddScript(script string) {
	s.Scripts = append(s.Scripts, script)
}

func (s *Sprite) Log() {
	log.Println(s.Name)
	for _, script := range s.Scripts {
		log.Println(script)
	}
}
