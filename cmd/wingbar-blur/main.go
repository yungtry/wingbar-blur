package main

import (
	"log"
	"os"
	"image"
	"image/color"
	"github.com/disintegration/imaging"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Path string `long:"path" description:"Path to the wallpaper." required:"true"`
	Theme string `short:"t" long:"theme" description:"Choose theme: none, light, dark." default:"none"`
	Opacity float64 `short:"o" long:"opacity" description:"Theme opacity" default:"0.4"`
	Blur float64 `short:"b" long:"blur" description:"Choose blur intensity" default:"10"`
	Size int `short:"s" long:"size" description:"Choose size" default:"20"`
	Change bool `short:"c" long:"change" description:"Set the output as wallpaper."`
}

func main() {
	flags.ParseArgs(&opts, os.Args[1:])

	src, err := imaging.Open(opts.Path)
	if err != nil {
		log.Fatalf("Please specify '--path' to the wallpaper. \nError: %v", err)
	}

	Process(src, opts.Theme, opts.Opacity, opts.Blur, opts.Size, opts.Change)
}

func Process(src image.Image, theme string, opacity float64, blur float64, size int, change bool) {
	img1 := imaging.Blur(src, blur)
	img2 := imaging.CropAnchor(img1, src.Bounds().Dx(), size, imaging.TopLeft)

	overlay := imaging.New(src.Bounds().Dx(), size, TranslateColor(theme))
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

func TranslateSize(size string) int{
	switch (size) {
	case "small":
		return 5
	case "medium":
		return 10
	case "big":
		return 15
	default:
		return 10
	}
}

func TranslateColor(theme string) color.NRGBA{
	switch(theme) {
	case "none":
		return color.NRGBA{0, 0, 0, 0}
	case "light":
		return color.NRGBA{255, 255, 255, 255}
	case "dark":
		return color.NRGBA{0, 0, 0, 255}
	default:
		return color.NRGBA{0, 0, 0, 0}
	}
}