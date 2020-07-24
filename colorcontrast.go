package colorcontrast

import (
	"image/color"
	"math"
)

// CalcContrastRatio calculate WCAG contrast ratio
func CalcContrastRatio(foreground, background color.Color) float64 {
	bgOnWhite := alphaBlend(background, color.White)
	bgOnBlack := alphaBlend(background, color.Black)

	lWhite := getRelativeLuminance(bgOnWhite)
	lBlack := getRelativeLuminance(bgOnBlack)
	lFg := getRelativeLuminance(foreground)

	if lWhite < lFg {
		return getContrastRatioOpaque(foreground, bgOnWhite)
	} else if lBlack > lFg {
		return getContrastRatioOpaque(foreground, bgOnBlack)
	} else {
		return 1
	}
}

func alphaBlend(foreground, background color.Color) color.RGBA {
	fr, fg, fb, fa := foreground.RGBA()
	br, bg, bb, _ := background.RGBA()

	return color.RGBA{
		R: uint8((fr*fa + br*(0xFFFF-fa)) / 0xFFFF),
		G: uint8((fg*fa + bg*(0xFFFF-fa)) / 0xFFFF),
		B: uint8((fb*fa + bb*(0xFFFF-fa)) / 0xFFFF),
		A: 0xFF,
	}
}

func getContrastRatioOpaque(foreground, background color.Color) float64 {
	l1 := getRelativeLuminance(background)
	l2 := getRelativeLuminance(alphaBlend(foreground, background))

	// https://www.w3.org/TR/2008/REC-WCAG20-20081211/#contrast-ratiodef
	return (math.Max(l1, l2) + 0.05) / (math.Min(l1, l2) + 0.05)
}

func getRelativeLuminance(c color.Color) float64 {
	// https://www.w3.org/TR/2008/REC-WCAG20-20081211/#relativeluminancedef
	cr, cg, cb, _ := c.RGBA()

	r := float64(cr) / 0xFFFF
	g := float64(cg) / 0xFFFF
	b := float64(cb) / 0xFFFF

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
