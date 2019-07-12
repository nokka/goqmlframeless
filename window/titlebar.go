// +build !darwin

package window

import (
	"fmt"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

const (
	iconSize = 15
)

func addTitleBarButtons(f *QFramelessWindow) {
	f.titleBarLayout.SetSpacing(2)

	f.iconMinimize = NewSVGButton(nil)
	f.iconMinimize.IconBtn.SetFixedSize2(iconSize, iconSize)
	f.iconMinimize.SetObjectName("IconMinimize")
	f.iconMinimize.SetStyle(nil)
	f.iconMinimize.Hide()

	f.iconClose = NewSVGButton(nil)
	f.iconClose.IconBtn.SetFixedSize2(iconSize, iconSize)
	f.iconClose.SetObjectName("IconClose")
	f.iconClose.SetStyle(nil)
	f.iconClose.Hide()

	f.titleBarLayout.SetAlignment(nil, core.Qt__AlignRight)

	// Add stretch to push buttons to the right side.
	f.titleBarLayout.AddStretch(0)
	f.titleBarLayout.AddWidget(f.iconMinimize.Widget, 0, 0)
	f.titleBarLayout.AddWidget(f.iconClose.Widget, 0, 0)
}

// styleTitlebarButtons ...
func styleTitlebarButtons(f *QFramelessWindow) {
	color := &RGB{R: 255, G: 255, B: 255}
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

	f.iconMinimize.IconBtn.Load2(core.NewQByteArray2(SvgMinimize, len(SvgMinimize)))
	f.iconClose.IconBtn.Load2(core.NewQByteArray2(SvgClose, len(SvgClose)))

	f.iconMinimize.Show()
	f.iconClose.Show()
}

// setupTitleBarActions ...
func setupTitleBarActions(f *QFramelessWindow) {
	// Setup minimize button actions.
	f.iconMinimize.Widget.ConnectEnterEvent(func(event *core.QEvent) {
		// Set style when hovering the icon.
		f.iconMinimize.SetStyle(&RGB{
			R: 0,
			G: 162,
			B: 232,
		})
	})

	f.iconMinimize.Widget.ConnectLeaveEvent(func(event *core.QEvent) {
		// Reset style when leaving the icon.
		f.iconMinimize.SetStyle(nil)
	})

	f.iconMinimize.Widget.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.isTitleBarPressed = false
	})

	f.iconMinimize.Widget.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		isContain := f.iconMinimize.Widget.Rect().Contains(e.Pos(), false)
		if !isContain {
			return
		}
		f.SetWindowState(core.Qt__WindowMinimized)
		f.Widget.Hide()
		f.Widget.Show()
	})

	// Setup close button actions.
	f.iconClose.Widget.ConnectEnterEvent(func(event *core.QEvent) {
		// Set style when hovering the icon.
		f.iconClose.SetStyle(&RGB{
			R: 0,
			G: 162,
			B: 232,
		})
	})

	f.iconClose.Widget.ConnectLeaveEvent(func(event *core.QEvent) {
		// Reset style when leaving the icon.
		f.iconClose.SetStyle(nil)
	})

	f.iconClose.Widget.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.isTitleBarPressed = false
	})

	f.iconClose.Widget.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		isContain := f.iconClose.Widget.Rect().Contains(e.Pos(), false)
		if !isContain {
			return
		}
		f.Close()
	})

	// Setup movable window.
	f.titleBar.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.Widget.Raise()
		f.isTitleBarPressed = true
		f.titleBarMousePos = e.GlobalPos()
		f.position = f.Pos()
	})

	f.titleBar.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		f.isTitleBarPressed = false
	})

	f.titleBar.ConnectMouseMoveEvent(func(e *gui.QMouseEvent) {
		if !f.isTitleBarPressed {
			return
		}
		x := f.position.X() + e.GlobalPos().X() - f.titleBarMousePos.X()
		y := f.position.Y() + e.GlobalPos().Y() - f.titleBarMousePos.Y()
		newPos := core.NewQPoint2(x, y)
		f.Move(newPos)
	})
}
