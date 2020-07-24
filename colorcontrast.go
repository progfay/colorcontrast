package colorcontrast

import (
	"image/color"
	"math"
)

// CalcContrastRatio calculate WCAG contrast ratio
func CalcContrastRatio(foreground, background color.Color) float64 {
	l1 := getRelativeLuminance(background)
	l2 := getRelativeLuminance(alphaBlend(foreground, background))

	// https://www.w3.org/TR/2008/REC-WCAG20-20081211/#contrast-ratiodef
	return (math.Max(l1, l2) + 0.05) / (math.Min(l1, l2) + 0.05)
}

func alphaBlend(foreground, background color.Color) color.RGBA {
	fr, fg, fb, fa := foreground.RGBA()
	br, bg, bb, _ := background.RGBA()

	fr >>= 8
	fg >>= 8
	fb >>= 8
	fa >>= 8
	br >>= 8
	bg >>= 8
	bb >>= 8

	return color.RGBA{
		R: uint8(math.Round(float64(fr*fa+br*(0xFF-fa)) / 0xFF)),
		G: uint8(math.Round(float64(fg*fa+bg*(0xFF-fa)) / 0xFF)),
		B: uint8(math.Round(float64(fb*fa+bb*(0xFF-fa)) / 0xFF)),
		A: 0xFF,
	}
}

func getRelativeLuminance(c color.Color) float64 {
	// https://www.w3.org/TR/2008/REC-WCAG20-20081211/#relativeluminancedef
	cr, cg, cb, _ := c.RGBA()

	r := float64(cr>>8) / 0xFF
	g := float64(cg>>8) / 0xFF
	b := float64(cb>>8) / 0xFF

	if r <= 0.03928 {
		r = r / 12.92
	} else {
		r = math.Pow((r+0.055)/1.055, 2.4)
	}

	if g <= 0.03928 {
		g = g / 12.92
	} else {
		g = math.Pow((g+0.055)/1.055, 2.4)
	}

	if b <= 0.03928 {
		b = b / 12.92
	} else {
		b = math.Pow((b+0.055)/1.055, 2.4)
	}

	return 0.2126*r + 0.7152*g + 0.0722*b
}
