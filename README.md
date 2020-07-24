## color contrast

Package for calculating color contrast ratio implemented by Go.

### Installation

```
go get -v github.com/progfay/colorcontrast
```


### API

#### `func CalcContrastRatio(foreground, background color.Color) float64`

- `foreground` ( [`color.Color`](https://golang.org/pkg/image/color/#Color) ) : Foreground color
- `background` ( [`color.Color`](https://golang.org/pkg/image/color/#Color) ) : Background color
- Return Color contrast ratio defined by WCAG
