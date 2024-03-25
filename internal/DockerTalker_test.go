package internal

import "testing"

func TestBuildImage(t *testing.T) {
	doc := Dockter{}
	doc.Init()
	defer doc.Close()
	// doc.CreateImage()
}
