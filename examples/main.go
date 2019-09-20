package main

import (
	"os"

	"github.com/nokka/goqmlframeless"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// Enable high dpi scaling, useful for devices with high pixel density displays.
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create new frameless window.
	fw := goqmlframeless.NewWindow(goqmlframeless.Options{
		Width:       1024,
		Height:      600,
		Alpha:       1.0,
		Color:       goqmlframeless.RGB{R: 0, G: 0, B: 0},
		BorderColor: &goqmlframeless.RGB{R: 198, G: 154, B: 31},
		ShadowSize:  0.0,
	})

	// QML Widget that will be used to draw on.
	qmlWidget := quick.NewQQuickWidget(nil)
	qmlWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	// Add QML widget to layout.
	fw.Layout.AddWidget(qmlWidget, 0, 0)

	// Set the QML source.
	qmlWidget.SetSource(core.NewQUrl3("main.qml", 0))

	// Make sure the window is allowed to minimize.
	goqmlframeless.AllowMinimize(fw.WinId())

	fw.Show()
	app.Exec()
}
