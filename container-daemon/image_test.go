package main

import "testing"

func TestNew(t *testing.T) {
	testFile := "./stub/image.json"
	target := Image{
		Cmd:      "./stub/hello",
		Hostname: "container",
		MaxMem:   "100000",
		MaxPids:  "10",
	}

	img, err := NewImage(testFile)
	if err != nil {
		t.Fatalf("some error occured: %v", err)
	}

	if *img != target {
		t.Errorf("received %v, want %v", img, target)
	}
}
