// +build darwin

package window

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

// addTitleBarButtons
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
	f.titleBarLayout.SetAlignment(f.TitleBarBtnWidget, core.Qt__AlignLeft)

	// Add buttons to the layout.
	f.titleBarLayout.AddWidget(f.btnClose, 0, 0)
	f.titleBarLayout.AddWidget(f.btnMinimize, 0, 0)
	f.titleBarLayout.AddWidget(f.titleLabel, 0, 0)
}

// colorizeTitlebarButtons ...
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
