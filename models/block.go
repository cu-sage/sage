package models

import "log"

type Block struct {
	Type  string   `xml:"type,attr"`
	Field []string `xml:"field"`
	Value *Block   `xml:"value>block"`
	Next  *Block   `xml:"next>block"`
}

func (b *Block) Log() {
	log.Printf("Type: %s\n", b.Type)
	log.Printf("Field: %s\n", b.Field)

	if b.Value != nil {
		log.Printf("Value:\n")
		b.Value.Log()
	}

	if b.Next != nil {
		log.Printf("Next:\n")
		b.Next.Log()
	}
}
