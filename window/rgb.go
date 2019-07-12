package window

import (
	"fmt"
	"math"
)

// RGB represents a color.
type RGB struct {
	R uint16
	G uint16
	B uint16
}

// Hex ...
func (c *RGB) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", uint8(c.R), uint8(c.G), uint8(c.B))
}

// Brend ...
func (c *RGB) Brend(color *RGB, alpha float64) *RGB {
	if color == nil {
		return &RGB{0, 0, 0}
	}
	return &RGB{
		R: uint16((float64(c.R) * float64(1-alpha)) + (float64(color.R) * float64(alpha))),
		G: uint16((float64(c.G) * float64(1-alpha)) + (float64(color.G) * float64(alpha))),
		B: uint16((float64(c.B) * float64(1-alpha)) + (float64(color.B) * float64(alpha))),
	}
}

func (c *RGB) fade() *RGB {
	r := (float64)(c.R)
	g := (float64)(c.G)
	b := (float64)(c.B)
	disp := (math.Abs(128-r) + math.Abs(128-g) + math.Abs(128-b)) / 3 * 1 / 4
	var newColor [3]float64
	for i, color := range []float64{
		r, g, b,
	} {
		if color > 128 {
			newColor[i] = color - disp
		} else {
			newColor[i] = color + disp
		}
		if newColor[i] < 0 {
			newColor[i] = 0
		} else if newColor[i] > 255 {
			newColor[i] = 255
		}
	}

	return &RGB{
		R: (uint16)(newColor[0]),
		G: (uint16)(newColor[1]),
		B: (uint16)(newColor[2]),
	}
}
