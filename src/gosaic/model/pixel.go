package model

import (
	"encoding/json"
	"errors"
)

type Pixel interface {
	GetData() []byte
	GetPixels() []*Lab
	SetData([]byte)
	SetPixels([]*Lab)
}

// PixelEncode encodes slice of Pixels to
// json-encoded []byte and stores in Data.
func PixelEncode(p Pixel) error {
	b, err := json.Marshal(p.GetPixels())
	if err != nil {
		return err
	}
	p.SetData(b)
	return nil
}

// PixelDecode decodes []byte of Data to
// slice of *Lab and stores in Pixels.
func PixelDecode(p Pixel) error {
	var pixels []*Lab
	err := json.Unmarshal(p.GetData(), &pixels)
	if err != nil {
		return err
	}
	p.SetPixels(pixels)
	return nil
}

func PixelDiff(p1, p2 Pixel) (float64, error) {
	if len(p1.GetPixels()) != len(p2.GetPixels()) {
		return nil, errors.New("Pixel slice not the same length")
	}

	result := float64(0.0)

	for i := 0; i < len(p1.GetPixels()); i++ {
		lab1 := p1.GetPixels()[i]
		lab2 := p2.GetPixels()[i]
		result += lab1.dist(lab2)
	}

	return result
}