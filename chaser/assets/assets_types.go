package assets

import "github.com/kemokemo/ebiten-sketchbook/chaser/assets/images"

type ImageType int

const (
	PlayerImage ImageType = iota
	EnemyImage
)

var imageMap map[ImageType][]byte

func init() {
	imageMap = make(map[ImageType][]byte)
	imageMap[PlayerImage] = images.PlayerImage
	imageMap[EnemyImage] = images.EnemyImage
}
