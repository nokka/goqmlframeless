package window

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/svg"
	"github.com/therecipe/qt/widgets"
)

// SVGButton toolbar for everything but darwin.
type SVGButton struct {
	f       *QFramelessWindow
	Widget  *widgets.QWidget
	IconBtn *svg.QSvgWidget
	isHover bool
}

// SetObjectName ...
func (b *SVGButton) SetObjectName(name string) {
	b.IconBtn.SetObjectName(name)
}

// Hide ...
func (b *SVGButton) Hide() {
	b.Widget.Hide()
}

// Show ...
func (b *SVGButton) Show() {
	b.Widget.Show()
}

// SetStyle ...
func (b *SVGButton) SetStyle(color *RGB) {
	backgroundColor := "background-color:none;"
	// Override background color if it's set.
	if color != nil {
		hoverColor := RGB{R: 0, B: 0, G: 0}
		backgroundColor = fmt.Sprintf("background-color: rgba(%d, %d, %d, %f);", hoverColor.R, hoverColor.G, hoverColor.B, 1)
	}

	b.Widget.SetStyleSheet(fmt.Sprintf(`
	.QWidget { 
		%s;
		border:none;
	}
	`, backgroundColor))
}

// NewSVGButton ...
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
	layout.SetAlignment(icon, core.Qt__AlignCenter)

	return &SVGButton{
		Widget:  widget,
		IconBtn: icon,
	}
}
