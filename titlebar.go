// +build !darwin

package goqmlframeless

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

const (
	iconSize = 15
)

// addTitleBarButtons will create all the buttons on the title bar.
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

// styleTitlebarButtons will style the OS specific buttons.
func styleTitlebarButtons(f *QFramelessWindow) {
	svgMinimize := getMinimizeSVG(RGB{R: 255, G: 255, B: 255})
	svgClose := getCloseSVG(RGB{R: 255, G: 255, B: 255})

	f.iconMinimize.IconBtn.Load2(core.NewQByteArray2(svgMinimize, len(svgMinimize)))
	f.iconClose.IconBtn.Load2(core.NewQByteArray2(svgClose, len(svgClose)))

	f.iconMinimize.Show()
	f.iconClose.Show()
}

// setupTitleBarEvents will setup all the events for the title bar.
func setupTitleBarEvents(f *QFramelessWindow) {
	// Setup minimize button actions.
	f.iconMinimize.Widget.ConnectEnterEvent(func(event *core.QEvent) {
		// Replace with hovered state svg.
		minimizeHovered := getMinimizeSVG(RGB{R: 155, G: 155, B: 155})
		f.iconMinimize.IconBtn.Load2(core.NewQByteArray2(minimizeHovered, len(minimizeHovered)))

		// Change cursor.
		cursor := gui.NewQCursor()
		cursor.SetShape(core.Qt__PointingHandCursor)
		f.SetCursor(cursor)
	})

	f.iconMinimize.Widget.ConnectLeaveEvent(func(event *core.QEvent) {
		// Replace with normal state svg.
		svgMinimize := getMinimizeSVG(RGB{R: 255, G: 255, B: 255})
		f.iconMinimize.IconBtn.Load2(core.NewQByteArray2(svgMinimize, len(svgMinimize)))

		// Reset style when leaving the icon.
		f.iconMinimize.SetStyle(nil)
		cursor := gui.NewQCursor()
		cursor.SetShape(core.Qt__ArrowCursor)
		f.SetCursor(cursor)
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
		// Replace with hovered state svg.
		closeHovered := getCloseSVG(RGB{R: 155, G: 155, B: 155})
		f.iconClose.IconBtn.Load2(core.NewQByteArray2(closeHovered, len(closeHovered)))

		cursor := gui.NewQCursor()
		cursor.SetShape(core.Qt__PointingHandCursor)
		f.SetCursor(cursor)
	})

	f.iconClose.Widget.ConnectLeaveEvent(func(event *core.QEvent) {
		// Replace with normal state svg.
		svgClose := getCloseSVG(RGB{R: 255, G: 255, B: 255})
		f.iconClose.IconBtn.Load2(core.NewQByteArray2(svgClose, len(svgClose)))

		// Reset style when leaving the icon.
		f.iconClose.SetStyle(nil)
		cursor := gui.NewQCursor()
		cursor.SetShape(core.Qt__ArrowCursor)
		f.SetCursor(cursor)
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

func getMinimizeSVG(color RGB) string {
	return fmt.Sprintf(`
	<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg">
	<g>
	<title>background</title>
	<rect fill="none" id="canvas_background" height="26" width="26" y="-1" x="-1"/>
	</g>
	<g>
	<title>Layer 1</title>
	<rect stroke="null" id="svg_3" height="3.352955" width="16.386974" y="20.712959" x="3.759288" stroke-opacity="null" fill="%s"/>
	</g>
	</svg>
	`, color.Hex())
}

func getCloseSVG(color RGB) string {
	return fmt.Sprintf(`
	<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg">
	<g>
	<title>background</title>
	<rect fill="none" id="canvas_background" height="26" width="26" y="-1" x="-1"/>
	</g>
	<g>
	<title>Layer 1</title>
	<path stroke="null" fill="%s" id="svg_1" d="m3.48053,9.684337l2.85088,-2.85088l5.701748,5.701748l5.701748,-5.701748l2.85088,2.85088l-5.701748,5.701748l5.701748,5.701748l-2.85088,2.85088l-5.701748,-5.701748l-5.701748,5.701748l-2.85088,-2.85088l5.70176,-5.701748l-5.70176,-5.701748z"/>
	</g>
	</svg>
	`, color.Hex())
}
