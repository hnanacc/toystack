package main

import (
	"encoding/json"
	"os"
)

type Image struct {
	Cmd      string
	MaxMem   string
	MaxPids  string
	Hostname string
}

func NewImage(p string) (img *Image, err error) {
	r, err := os.ReadFile(p)
	if err != nil {
		return
	}
	img = &Image{}
	err = json.Unmarshal(r, img)
	return
}
