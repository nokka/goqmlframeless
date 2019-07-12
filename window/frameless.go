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

	// Attributes
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

	// QUESTIONABLE
	TitleBarBtnWidget *widgets.QWidget
	titleLabel        *widgets.QLabel

	// ICONS, REFACTOR?

	// Darwin title bar buttons.
	btnMinimize *widgets.QToolButton
	btnClose    *widgets.QToolButton

	// Windows and Linux title bar buttons.
	iconMinimize *QToolButtonForNotDarwin
	iconRestore  *QToolButtonForNotDarwin
	iconClose    *QToolButtonForNotDarwin
}

// NewFramelessWindow ...
func NewFramelessWindow(width int, height int) *QFramelessWindow {
	f := NewQFramelessWindow(nil, 0)
	f.SetFixedSize2(width, height)

	// TODO: Parameterize
	f.frameColor = &RGB{R: 0, B: 0, G: 0}
	f.shadowMargin = 1
	f.borderSize = 1
	f.colorAlpha = 1

	// Central widget and layout.
	f.Widget = newWidget()

	// TODO: Implement real margin.
	f.Layout = newLayout(1)

	// Set layout and central widget.
	f.Widget.SetLayout(f.Layout)
	f.SetCentralWidget(f.Widget)

	// Configuration.
	f.createFrame()
	f.setupWindowFlags()
	f.SetupAttributes()

	// MISSING:
	// - Title bar actions

	f.SetupFrameEvents()

	f.SetupFrameColor()
	f.SetupFrameShadow()

	return f
}

func newWidget() *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetSizePolicy2(widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum, widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum)
	return widget
}

func newLayout(margin int) *widgets.QVBoxLayout {
	layout := widgets.NewQVBoxLayout()
	layout.SetContentsMargins(margin, margin, margin, margin)
	layout.SetSpacing(0)
	return layout
}

// SetupUI ...
func (f *QFramelessWindow) createFrame() {
	f.InstallEventFilter(f)

	// FRAME (f.WindowWidget)
	f.frame = widgets.NewQFrame(f.Widget, 0)
	f.frame.SetObjectName("QFramelessWidget")
	f.frame.SetSizePolicy2(widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum, widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum)

	// FRAME LAYOUT (f.WindowVLayout)
	f.frameLayout = widgets.NewQVBoxLayout2(f.frame)
	f.frameLayout.SetContentsMargins(f.borderSize, f.borderSize, f.borderSize, 0)
	f.frameLayout.SetContentsMargins(0, 0, 0, 0)
	f.frameLayout.SetSpacing(0)
	f.frame.SetLayout(f.frameLayout)

	// TITLE BAR
	f.titleBar = widgets.NewQWidget(f.frame, 0)
	f.titleBar.SetObjectName("TitleBar")
	f.titleBar.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed)

	// TITLE BAR LAYOUT
	f.titleBarLayout = widgets.NewQHBoxLayout2(f.titleBar)
	f.titleBarLayout.SetContentsMargins(0, 0, 0, 0)

	// TITLE LABEL
	f.titleLabel = widgets.NewQLabel(nil, 0)
	f.titleLabel.SetObjectName("TitleLabel")
	f.titleLabel.SetAlignment(core.Qt__AlignCenter)

	// BUTTONS
	addTitleBarButtons(f)

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
func (f *QFramelessWindow) SetupAttributes() {
	f.SetAttribute(core.Qt__WA_TranslucentBackground, true)
	f.SetAttribute(core.Qt__WA_NoSystemBackground, true)
	f.SetAttribute(core.Qt__WA_Hover, true)
	f.SetMouseTracking(true)
}
