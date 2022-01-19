package compiler

import "testing"

func TestGenerate(t *testing.T) {
	g, err := NewPullGenerator()
	if err != nil {
		t.Fatal(err)
	}

	err = g.Run()
	if err != nil {
		t.Fatal(err)
	}
}
