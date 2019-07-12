package window

import (
	"fmt"

	"github.com/therecipe/qt/core"
)

// SetupFrameEvents ...
func (f *QFramelessWindow) SetupFrameEvents() {
	f.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
		fmt.Println("EVENT FILTER RUNNING")
		//e := gui.NewQMouseEventFromPointer(core.PointerFromQEvent(event))
		switch event.Type() {
		case core.QEvent__ActivationChange:
			fmt.Println("ACTIVIATION CHANGE RUNNING")
			styleTitlebarButtons(f)

		/*case core.QEvent__HoverMove:
			f.updateCursorShape(e.GlobalPos())

		case core.QEvent__Leave:
			cursor := gui.NewQCursor()
			cursor.SetShape(core.Qt__ArrowCursor)
			f.SetCursor(cursor)

		case core.QEvent__MouseMove:
			f.mouseMove(e)

		case core.QEvent__MouseButtonPress:
			f.mouseButtonPressed(e)

		case core.QEvent__MouseButtonRelease:
			f.isDragStart = false
			f.isLeftButtonPressed = false
			f.hoverEdge = None
		*/
		default:
		}

		return f.Widget.EventFilter(watched, event)
	})
}
