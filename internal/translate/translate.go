package translate

import (
	"image/color"
)

func Color(theme string) color.NRGBA{
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