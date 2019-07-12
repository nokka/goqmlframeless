package window

import (
	"fmt"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/svg"
	"github.com/therecipe/qt/widgets"
)

// QToolButtonForNotDarwin toolbar for everything but darwin.
type QToolButtonForNotDarwin struct {
	f       *QFramelessWindow
	Widget  *widgets.QWidget
	IconBtn *svg.QSvgWidget
	isHover bool
}

// SetObjectName ...
func (b *QToolButtonForNotDarwin) SetObjectName(name string) {
	b.IconBtn.SetObjectName(name)
}

// Hide ...
func (b *QToolButtonForNotDarwin) Hide() {
	b.Widget.Hide()
}

// Show ...
func (b *QToolButtonForNotDarwin) Show() {
	b.Widget.Show()
}

// SetStyle ...
func (b *QToolButtonForNotDarwin) SetStyle(color *RGB) {
	var backgroundColor string
	if color == nil {
		backgroundColor = "background-color:none;"
	} else {
		hoverColor := color.Brend(b.f.frameColor, 0.75)
		backgroundColor = fmt.Sprintf("background-color: rgba(%d, %d, %d, %f);", hoverColor.R, hoverColor.G, hoverColor.B, b.f.colorAlpha)
	}

	b.Widget.SetStyleSheet(fmt.Sprintf(`
	.QWidget { 
		%s;
		border:none;
	}
	`, backgroundColor))
}

// SetupTitleBarColorForNotDarwin ...
func (f *QFramelessWindow) SetupTitleBarColorForNotDarwin(color *RGB) {
	if color == nil {
		color = &RGB{
			R: 128,
			G: 128,
			B: 128,
		}
	} else {
		color = color.fade()
	}
	var SvgMinimize, SvgClose string

	if runtime.GOOS == "windows" {
		SvgMinimize = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<path fill="%s" d="M20,14H4V10H20" />
		</svg>
		`, color.Hex())

		SvgClose = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<path fill="%s" d="M13.46,12L19,17.54V19H17.54L12,13.46L6.46,19H5V17.54L10.54,12L5,6.46V5H6.46L12,10.54L17.54,5H19V6.46L13.46,12Z" />
		</svg>
		`, color.Hex())
	} else {
		SvgMinimize = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<path fill="%s" d="M17,13H7V11H17M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" />
		</svg>
		`, color.Hex())

		SvgClose = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<g transform="translate(0,1)">
		<path fill="%s" d="M12 2C6.47 2 2 6.47 2 12s4.47 10 10 10 10-4.47 10-10S17.53 2 12 2zm5 13.59L15.59 17 12 13.41 8.41 17 7 15.59 10.59 12 7 8.41 8.41 7 12 10.59 15.59 7 17 8.41 13.41 12 17 15.59z"/><path d="M0 0h24v24H0z" fill="none"/></g></svg>
		`, "#e86032")
	}

	fmt.Println("ICON MINIMIZE BUTTON")
	fmt.Println(f.iconMinimize)

	f.iconMinimize.IconBtn.Load2(core.NewQByteArray2(SvgMinimize, len(SvgMinimize)))
	f.iconClose.IconBtn.Load2(core.NewQByteArray2(SvgClose, len(SvgClose)))

	f.iconMinimize.Show()
	f.iconClose.Show()
}

// TITLE BAR BUTTONS

// SetTitleBarButtons ...
func (f *QFramelessWindow) SetTitleBarButtons() {
	iconSize := 15
	f.titleBarLayout.SetSpacing(1)

	f.iconMinimize = NewQToolButtonForNotDarwin(nil)
	f.iconMinimize.f = f
	f.iconMinimize.IconBtn.SetFixedSize2(iconSize, iconSize)
	f.iconMinimize.SetObjectName("IconMinimize")

	f.iconClose = NewQToolButtonForNotDarwin(nil)
	f.iconClose.f = f
	f.iconClose.IconBtn.SetFixedSize2(iconSize, iconSize)
	f.iconClose.SetObjectName("IconClose")

	f.SetIconsStyle(nil)

	f.iconMinimize.Hide()
	f.iconClose.Hide()

	f.titleBarLayout.SetAlignment(f.TitleBarBtnWidget, core.Qt__AlignRight)
	f.titleBarLayout.AddWidget(f.titleLabel, 0, 0)
	f.titleBarLayout.AddWidget(f.iconMinimize.Widget, 0, 0)
	f.titleBarLayout.AddWidget(f.iconClose.Widget, 0, 0)
}

// NewQToolButtonForNotDarwin ...
func NewQToolButtonForNotDarwin(parent widgets.QWidget_ITF) *QToolButtonForNotDarwin {
	iconSize := 15
	marginTB := iconSize / 6
	marginLR := 1
	if runtime.GOOS == "linux" {
		iconSize = 18
		marginLR = int(float64(iconSize) / float64(3.5))
	} else {
		marginLR = int(float64(iconSize) / float64(2.5))
	}

	widget := widgets.NewQWidget(parent, 0)
	widget.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	layout := widgets.NewQVBoxLayout2(widget)
	layout.SetContentsMargins(marginLR, marginTB, marginLR, marginTB)
	icon := svg.NewQSvgWidget(nil)
	icon.SetFixedSize2(iconSize, iconSize)

	layout.AddWidget(icon, 0, 0)
	layout.SetAlignment(icon, core.Qt__AlignCenter)

	return &QToolButtonForNotDarwin{
		Widget:  widget,
		IconBtn: icon,
	}
}

// SetIconsStyle ...
func (f *QFramelessWindow) SetIconsStyle(color *RGB) {
	for _, b := range []*QToolButtonForNotDarwin{
		f.iconMinimize,
		f.iconClose,
	} {
		b.SetStyle(color)
	}
}
