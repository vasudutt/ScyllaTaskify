package uuid

import "github.com/google/uuid"

type generator struct{}

func (g *generator) Generate() string {
	return uuid.New().String()
}

func (g *generator) Parse(s string) error {
	_, err := uuid.Parse(s)
	return err
}
