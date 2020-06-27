package goqmlframeless

import (
	"fmt"
)

// setupFrameColor will setup the color for the frame.
func (f *QFramelessWindow) setupFrameColor() {
	style := fmt.Sprintf("background-color: rgba(%d, %d, %d, %f);", f.frameColor.R, f.frameColor.G, f.frameColor.B, f.colorAlpha)
	f.Widget.SetStyleSheet(" * { background-color: rgba(0, 0, 0, 0.0); color: rgba(0, 0, 0, 0); }")

	var borderTop string
	if f.borderColor != nil {
		borderTop = fmt.Sprintf("border-top: 2px solid %s;", f.borderColor.Hex())
	}

	roundSizeString := fmt.Sprintf("%d", f.borderRadius) + "px"

	f.frame.SetStyleSheet(fmt.Sprintf(`
	#QFramelessWidget {
		padding-top: 2px; padding-right: 2px; padding-bottom: 2px; padding-left: 2px;
		border-top-left-radius: %s;
		border-top-right-radius: %s;
		%s;
		%s; 
	}`, roundSizeString, roundSizeString, borderTop, style))
}
