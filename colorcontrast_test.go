package colorcontrast_test

import (
	"fmt"
	"image/color"
	"math"
	"testing"

	"github.com/progfay/colorcontrast"
)

var testcases = []struct {
	foreground color.Color
	background color.Color
	contrast   float64
}{
	{
		color.White,
		color.Black,
		21,
	},
	{
		color.Black,
		color.White,
		21,
	},
	{
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		color.RGBA{R: 0, G: 0, B: 255, A: 255},
		2.14,
	},
}

func TestColorContrast(t *testing.T) {
	for _, testcase := range testcases {
		fr, fg, fb, fa := testcase.foreground.RGBA()
		f := fmt.Sprintf("#%02X%02X%02X%02X", fr>>8, fg>>8, fb>>8, fa>>8)
		br, bg, bb, ba := testcase.background.RGBA()
		b := fmt.Sprintf("#%02X%02X%02X%02X", br>>8, bg>>8, bb>>8, ba>>8)
		t.Run(fmt.Sprintf("foreground: %s, background: %s", f, b), func(t *testing.T) {
			got := colorcontrast.CalcContrastRatio(testcase.foreground, testcase.background)
			got = math.Floor(got*100) / 100
			if got != testcase.contrast {
				t.Errorf("got %v, want %v", got, testcase.contrast)
			}
		})
	}
}
