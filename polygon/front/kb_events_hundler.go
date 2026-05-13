package front

import w "github.com/kallaurru/interview/polygon/front/widgets"

func OnDash(ef *EFront) *w.Wdg {
	id := 1
	return ef.widgets[id]
}

func PlugActionsKbd(ef *EFront, kbd uint32) {
	switch kbd {
	default:
		return
	case constKeyMenuA:
		wdg := ef.widgets[10]
		ef.swapFocus(wdg.Id)
		ef.onDraw(wdg)
	case constKeyMenuB:
		wdg := ef.widgets[11]
		ef.swapFocus(wdg.Id)
		ef.onDraw(wdg)
	}
}
