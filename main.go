package main

import (
	"os"

	"github.com/nokka/widget-test/window"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func main() {

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create new frameless window.
	fw := window.NewFramelessWindow(1024, 600)

	// QML Widget that will be used to draw on.
	qmlWidget := newQmlWidget()

	fw.Layout.AddWidget(qmlWidget, 0, 0)

	qmlWidget.SetSource(core.NewQUrl3("qml/main.qml", 0))

	fw.Show()
	app.Exec()
}

/*func main() {

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create a window
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Hello Widgets Example")

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	widget := widgets.NewQWidget(nil, 0)
	qmlWidget := newQmlWidget()
	layout := widgets.NewQVBoxLayout()

	layout.AddWidget(qmlWidget, 0, 0)
	widget.SetLayout(layout)
	window.SetCentralWidget(widget)

	qmlWidget.SetSource(core.NewQUrl3("qml/main.qml", 0))

	// make the window visible
	window.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}*/

// newQmlWidget returns a configured QML widget.
func newQmlWidget() *quick.QQuickWidget {
	var qwidget = quick.NewQQuickWidget(nil)
	qwidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	return qwidget
}
