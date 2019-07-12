package window

import (
	"fmt"
	"runtime"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// setupFrameColor ...
func (f *QFramelessWindow) setupFrameColor() {
	style := fmt.Sprintf("background-color: rgba(%d, %d, %d, %f);", f.frameColor.R, f.frameColor.G, f.frameColor.B, f.colorAlpha)
	f.Widget.SetStyleSheet(" * { background-color: rgba(0, 0, 0, 0.0); color: rgba(0, 0, 0, 0); }")

	// TODO: Proper radius.
	borderSizeString := fmt.Sprintf("%d", f.borderSize*2) + "px"

	f.frame.SetStyleSheet(fmt.Sprintf(`
	#QFramelessWidget {
		border: 0px solid %s; 
		padding-top: 2px; padding-right: %s; padding-bottom: %s; padding-left: %s; 
		%s; 
	}`, f.frameColor.Hex(), borderSizeString, borderSizeString, borderSizeString, style))
}

// setupFrameShadow ...
func (f *QFramelessWindow) setupFrameShadow() {
	f.Layout.SetContentsMargins(f.shadowMargin, f.shadowMargin, f.shadowMargin, f.shadowMargin)

	if f.shadowMargin == 0 {
		return
	}

	shadow := widgets.NewQGraphicsDropShadowEffect(nil)
	var alpha int
	if runtime.GOOS == "darwin" {
		alpha = 220
		shadow.SetOffset3(0, 10)
	} else {
		alpha = 100
		shadow.SetOffset3(0, 0)
	}

	shadow.SetBlurRadius((float64)(f.shadowMargin))
	shadow.SetColor(gui.NewQColor3(0, 0, 0, alpha))
	f.frame.SetGraphicsEffect(shadow)
}
