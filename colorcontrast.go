package colorcontrast

import (
	"image/color"
	"math"
)

// CalcContrastRatio calculate WCAG contrast ratio
func CalcContrastRatio(foreground, background color.Color) float64 {
	l1 := getRelativeLuminance(background.RGBA())
	l2 := getRelativeLuminance(alphaBlend(foreground, background))

	// https://www.w3.org/TR/2008/REC-WCAG20-20081211/#contrast-ratiodef
	return (math.Max(l1, l2) + 0.05) / (math.Min(l1, l2) + 0.05)
}

func alphaBlend(foreground, background color.Color) (r, g, b, a uint32) {
	fr, fg, fb, fa := foreground.RGBA()
	br, bg, bb, _ := background.RGBA()

	r = uint32((fr*fa + br*(0xFFFF-fa)) >> 16)
	g = uint32((fg*fa + bg*(0xFFFF-fa)) >> 16)
	b = uint32((fb*fa + bb*(0xFFFF-fa)) >> 16)
	a = 0xFFFF
	return
}

func getRelativeLuminance(cr, cg, cb, _ uint32) float64 {
	// https://www.w3.org/TR/2008/REC-WCAG20-20081211/#relativeluminancedef
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
