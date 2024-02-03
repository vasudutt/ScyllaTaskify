package uuid

type Generator interface {
	Generate() string
	Parse(string) error
}

func New() Generator {
	return &generator{}
}
