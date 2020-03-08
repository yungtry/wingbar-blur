package process

import (
	"log"
	"image"
	"image/color"
	"github.com/disintegration/imaging"
	"github.com/yungtry/wingbar-blur/internal/translate"
)

func Process(src image.Image, theme string, opacity float64, blur float64, size int, change bool) {
	img1 := imaging.Blur(src, blur)
	img2 := imaging.CropAnchor(img1, src.Bounds().Dx(), size, imaging.TopLeft)

	overlay := imaging.New(src.Bounds().Dx(), size, translate.Color(theme))
	img3 := imaging.Overlay(img2, overlay, image.Pt(0, 0), opacity)

	dst := imaging.New(src.Bounds().Dx(), src.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, src, image.Pt(0, 0))
	dst = imaging.Paste(dst, img3, image.Pt(0, 0))
	err := imaging.Save(dst, "blur.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	//set-wallpaper ~/Downloads/coolpicture.png
}