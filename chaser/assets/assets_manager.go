package assets

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
)

type AssetsManager struct {
	imageMap map[ImageType]*ebiten.Image
}

func NewAssetsManager() *AssetsManager {
	return &AssetsManager{
		imageMap: make(map[ImageType]*ebiten.Image),
	}
}

func (am *AssetsManager) LoadAssets() error {
	for k, v := range imageMap {
		img, _, err := image.Decode(bytes.NewReader(v))
		if err != nil {
			return err
		}
		am.imageMap[k], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		if err != nil {
			return err
		}
	}
	return nil
}

func (am *AssetsManager) GetImage(at ImageType) *ebiten.Image {
	// TODO: check exist and return error as needed
	return am.imageMap[at]
}
