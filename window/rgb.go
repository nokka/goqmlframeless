package window

import (
	"fmt"
)

// RGB represents a color.
type RGB struct {
	R uint16
	G uint16
	B uint16
}

// Hex returns the hex value of the color.
func (c *RGB) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", uint8(c.R), uint8(c.G), uint8(c.B))
}
