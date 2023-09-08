package filex

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type ImageInfo struct {
	Width  int
	Height int
}

func GetImageInfo(imagePath string) (*ImageInfo, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}

	c, _, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}

	return &ImageInfo{
		Width:  c.Width,
		Height: c.Height,
	}, nil
}

// imaging.Blur(srcImage, 0.5)
