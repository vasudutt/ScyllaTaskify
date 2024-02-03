package uuid

import "testing"

func TestGenerator(t *testing.T) {
	g := New()
	uuid := g.Generate()
	if err := g.Parse(uuid); err != nil {
		t.Errorf("parsing generated UUID: %v", err)
	}

	if err := g.Parse("invalid-uuid"); err == nil {
		t.Errorf("parsing invalid UUID: %v", err)
	}

	if err := g.Parse(""); err == nil {
		t.Errorf("parsing empty UUID: %v", err)
	}

	if err := g.Parse("00000000-0000-0000-0000-000000000000"); err != nil {
		t.Errorf("parsing zero UUID: %v", err)
	}
}
