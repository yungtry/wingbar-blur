package main

import (
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/jessevdk/go-flags"
	"github.com/yungtry/wingbar-blur/internal/process"
)

var opts struct {
	Path    string  `long:"path" description:"Path to the wallpaper." required:"true"`
	Theme   string  `short:"t" long:"theme" description:"Choose theme: none, light, dark." default:"none"`
	Opacity float64 `short:"o" long:"opacity" description:"Theme opacity" default:"0.15"`
	Blur    float64 `short:"b" long:"blur" description:"Choose blur intensity" default:"20"`
	Size    int     `short:"s" long:"size" description:"Choose size" default:"30"`
	//Change  bool    `short:"c" long:"change" description:"Set the output as wallpaper."`
}

func main() {
	flags.ParseArgs(&opts, os.Args[1:])

	src, err := imaging.Open(opts.Path)
	if err != nil {
		log.Fatalf("Please specify '--path' to the wallpaper. \nError: %v", err)
	}

	process.Process(src, opts.Theme, opts.Opacity, opts.Blur, opts.Size)
}
