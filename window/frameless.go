package window

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

// QFramelessWindow is the main frameless window.
type QFramelessWindow struct {
	widgets.QMainWindow

	Widget *widgets.QWidget
	Layout *widgets.QVBoxLayout

	// Attributes.
	shadowMargin int
	borderSize   int
	colorAlpha   float64

	// Frame.
	frame       *widgets.QFrame
	frameLayout *widgets.QVBoxLayout
	frameColor  *RGB

	// Title bar.
	titleBar       *widgets.QWidget
	titleBarLayout *widgets.QHBoxLayout

	// Darwin title bar buttons.
	btnMinimize *widgets.QToolButton
	btnClose    *widgets.QToolButton

	// Windows and Linux title bar buttons.
	iconMinimize *SVGButton
	iconClose    *SVGButton

	// Mouse events.
	titleBarMousePos  *core.QPoint
	position          *core.QPoint
	isTitleBarPressed bool
}

// NewFramelessWindow ...
func NewFramelessWindow(width int, height int) *QFramelessWindow {
	f := NewQFramelessWindow(nil, 0)
	f.SetFixedSize2(width, height)

	f.frameColor = &RGB{R: 8, B: 8, G: 5}
	f.shadowMargin = 0
	f.borderSize = 1
	f.colorAlpha = 1

	// Central widget and layout.
	f.Widget = newWidget()
	f.Layout = newLayout()

	// Set layout and central widget.
	f.Widget.SetLayout(f.Layout)
	f.SetCentralWidget(f.Widget)

	// Create the frame for the window.
	f.createFrame()
	f.setupFrameColor()
	f.setupFrameShadow()

	// Setup window flags and attributes.
	f.setupWindowFlags()
	f.setupAttributes()

	// Setup OS specific title bar.
	addTitleBarButtons(f)
	styleTitlebarButtons(f)

	// Setup event handling for title bar.
	setupTitleBarActions(f)

	return f
}

func newWidget() *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetSizePolicy2(widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum, widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum)
	return widget
}

func newLayout() *widgets.QVBoxLayout {
	layout := widgets.NewQVBoxLayout()
	layout.SetContentsMargins(0, 0, 0, 0)
	layout.SetSpacing(0)
	return layout
}

// SetupUI ...
func (f *QFramelessWindow) createFrame() {
	f.InstallEventFilter(f)

	// Setup the frame, this is the base we'll add everything too.
	f.frame = widgets.NewQFrame(f.Widget, 0)
	f.frame.SetObjectName("QFramelessWidget")
	f.frame.SetSizePolicy2(widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum, widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum)

	// Setup frame layout which is add to the frame.
	f.frameLayout = widgets.NewQVBoxLayout2(f.frame)
	f.frameLayout.SetContentsMargins(f.borderSize, f.borderSize, f.borderSize, 0)
	f.frameLayout.SetContentsMargins(0, 0, 0, 0)
	f.frameLayout.SetSpacing(0)
	f.frame.SetLayout(f.frameLayout)

	// Title bar is used for dragging the window around and also has icons.
	f.titleBar = widgets.NewQWidget(f.frame, 0)
	f.titleBar.SetObjectName("TitleBar")
	f.titleBar.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed)

	// // Title bar layout is added to the title bar, to add widgets on.
	f.titleBarLayout = widgets.NewQHBoxLayout2(f.titleBar)
	f.titleBarLayout.SetContentsMargins(0, 0, 0, 0)

	// Add the title bar to the frame.
	f.frameLayout.AddWidget(f.titleBar, 0, 0)

	// Finally, add the frame to the main layout.
	f.Layout.AddWidget(f.frame, 0, 0)
}

// SetupWindowFlags ...
func (f *QFramelessWindow) setupWindowFlags() {
	f.SetWindowFlag(core.Qt__Window, true)
	f.SetWindowFlag(core.Qt__FramelessWindowHint, true)
	f.SetWindowFlag(core.Qt__NoDropShadowWindowHint, true)
	f.SetWindowFlag(core.Qt__MSWindowsFixedSizeDialogHint, true)
}

// SetupAttributes ...
func (f *QFramelessWindow) setupAttributes() {
	f.SetAttribute(core.Qt__WA_TranslucentBackground, true)
	f.SetAttribute(core.Qt__WA_NoSystemBackground, true)
	f.SetAttribute(core.Qt__WA_Hover, true)
	f.SetMouseTracking(true)
}
