package debug

import (
	"fmt"

	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/color"
)

func DebugHere(msg ...any) {
	if global.EnableSetting.Debug {
		s := fmt.Sprintf("---> %s", msg)
		fmt.Println(color.Cyan(s))
	}
}
