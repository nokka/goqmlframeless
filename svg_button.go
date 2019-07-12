package goqmlframeless

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/svg"
	"github.com/therecipe/qt/widgets"
)

// SVGButton toolbar for everything but darwin.
type SVGButton struct {
	Widget  *widgets.QWidget
	IconBtn *svg.QSvgWidget
	isHover bool
}

// SetObjectName will set the object name.
func (b *SVGButton) SetObjectName(name string) {
	b.IconBtn.SetObjectName(name)
}

// Hide will hide the button.
func (b *SVGButton) Hide() {
	b.Widget.Hide()
}

// Show will show the button.
func (b *SVGButton) Show() {
	b.Widget.Show()
}

// SetStyle will set the style for the button.
func (b *SVGButton) SetStyle(color *RGB) {
	backgroundColor := "background-color:none;"

	// Override background color if it's set.
	if color != nil {
		backgroundColor = fmt.Sprintf("background-color: rgba(%d, %d, %d, %f);", color.R, color.G, color.B, 1.0)
	}

	b.Widget.SetStyleSheet(fmt.Sprintf(`
	.QWidget { 
		%s;
		border:none;
	}
	`, backgroundColor))
}

// NewSVGButton creates a new SVG button.
func NewSVGButton(parent widgets.QWidget_ITF) *SVGButton {
	iconSize := 15
	marginTB := iconSize / 6
	marginLR := 1

	// Widget for button.
	widget := widgets.NewQWidget(parent, 0)
	widget.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)

	// Layout used for button.
	layout := widgets.NewQVBoxLayout2(widget)
	layout.SetContentsMargins(marginLR, marginTB, marginLR, marginTB)

	// Create the SVG widget.
	icon := svg.NewQSvgWidget(nil)
	icon.SetFixedSize2(iconSize, iconSize)

	// Add icon to the layout.
	layout.AddWidget(icon, 0, 0)
	layout.SetAlignment(icon, core.Qt__AlignRight)

	return &SVGButton{
		Widget:  widget,
		IconBtn: icon,
	}
}
