// +build darwin

package goqmlframeless

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// addTitleBarButtons will create all the buttons on the title bar.
func addTitleBarButtons(f *QFramelessWindow) {
	// Sizing policy.
	btnSizePolicy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed, widgets.QSizePolicy__ToolButton)

	// Minimize button.
	f.btnMinimize = widgets.NewQToolButton(f.titleBar)
	f.btnMinimize.SetObjectName("BtnMinimize")
	f.btnMinimize.SetSizePolicy(btnSizePolicy)

	// Close button.
	f.btnClose = widgets.NewQToolButton(f.titleBar)
	f.btnClose.SetObjectName("BtnClose")
	f.btnClose.SetSizePolicy(btnSizePolicy)

	// Title bar layout.
	f.titleBarLayout.SetSpacing(0)
	f.titleBarLayout.SetAlignment(nil, core.Qt__AlignLeft)

	// Add buttons to the layout.
	f.titleBarLayout.AddWidget(f.btnClose, 0, 0)
	f.titleBarLayout.AddWidget(f.btnMinimize, 0, 0)
	f.titleBarLayout.AddStretch(0)
}

// styleTitlebarButtons will style the OS specific buttons.
func styleTitlebarButtons(f *QFramelessWindow) {
	var baseStyle, minimizeColor, closeColor string
	baseStyle = ` #BtnMinimize, #BtnClose {
		min-width: 10px;
		min-height: 10px;
		max-width: 10px;
		max-height: 10px;
		border-radius: 6px;
		border-width: 1px;
		border-style: solid;
		margin: 4px;
	}`
	minimizeColor = `
		#BtnMinimize {
			background-color: rgba(128, 128, 128, 0.3);
			border-color: rgb(128, 128, 128, 0.2);
		}
	`
	closeColor = `
		#BtnClose {
			background-color: rgba(128, 128, 128, 0.3);
			border-color: rgb(128, 128, 128, 0.2);
		}
	`

	minimizeColorHover := `
		#BtnMinimize:hover {
			background-color: rgb(253, 190, 65);
			border-color: rgb(239, 170, 47);
			background-repeat: no-repeat;
			background-position: center center; 
		}
	`
	closeColorHover := `
		#BtnClose:hover {
			background-color: rgb(252, 98, 93);
			border-color: rgb(239, 75, 71);
			background-repeat: no-repeat;
			background-position: center center; 
		}
	`

	f.btnMinimize.SetStyleSheet(baseStyle + minimizeColor + minimizeColorHover)
	f.btnClose.SetStyleSheet(baseStyle + closeColor + closeColorHover)
}

// setupTitleBarEvents will setup all the events for the title bar.
func setupTitleBarEvents(f *QFramelessWindow) {
	// Setup title bar actions.
	f.titleBar.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.Widget.Raise()
		f.isTitleBarPressed = true
		f.titleBarMousePos = e.GlobalPos()
		f.position = f.Pos()
	})

	f.titleBar.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		f.isTitleBarPressed = false
	})

	// Setup movable window.
	f.titleBar.ConnectMouseMoveEvent(func(e *gui.QMouseEvent) {
		if !f.isTitleBarPressed {
			return
		}
		x := f.position.X() + e.GlobalPos().X() - f.titleBarMousePos.X()
		y := f.position.Y() + e.GlobalPos().Y() - f.titleBarMousePos.Y()
		newPos := core.NewQPoint2(x, y)
		f.Move(newPos)
	})

	// Minimize button events.
	f.btnMinimize.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.isTitleBarPressed = false
	})

	f.btnMinimize.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		f.SetWindowState(core.Qt__WindowMinimized)
		f.Widget.Hide()
		f.Widget.Show()
	})

	// Close button events.
	f.btnClose.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.isTitleBarPressed = false
	})

	f.btnClose.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		f.Close()
	})
}
