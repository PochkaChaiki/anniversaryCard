package imageprocessor

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

const (
	fontFile  = "./static/Xiomara-wWLw.ttf"
	imageFile = "./static/picture.png"
	size      = 40
)

func CreateImage(text string) (image.Image, error) {
	bgImage, err := gg.LoadImage(imageFile)
	imageWidth, imageHeight := bgImage.Bounds().Dx(), bgImage.Bounds().Dy()
	if err != nil {
		return nil, err
	}
	dc := gg.NewContext(imageWidth, imageHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(fontFile, size); err != nil {
		return nil, err
	}
	x := float64(imageWidth / 2)
	y := float64(imageHeight / 2)
	maxWidth := float64(imageWidth) - 60
	dc.SetColor(color.Black)
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	return dc.Image(), nil
}
