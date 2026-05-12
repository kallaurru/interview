package front

import (
	"fmt"
	glg "github.com/bahlo/generic-list-go"
	t "github.com/kallaurru/interview/polygon/front/talk"
	w "github.com/kallaurru/interview/polygon/front/widgets"
	"sync/atomic"
)

type OnActionF func(ef *EFront) *w.Wdg

type EFront struct {
	widgets    []*w.Wdg // 0 nil от 10 начинаются модалки
	onDash     OnActionF
	focus      atomic.Uint32
	focusItems *glg.List
}

const (
	constKeyEnter uint32 = 0x01
	constKeyPgUp  uint32 = 0x02
	constKeyPgDw  uint32 = 0x04
	constKeyUp    uint32 = 0x10
	constKeyDw    uint32 = 0x20
	constKeyTab   uint32 = 0x0100
	constKeyEsc   uint32 = 0x010000

	constFocus int = 0
)

func New(wdgCount uint8) *EFront {
	return &EFront{
		widgets: make([]*w.Wdg, 0, wdgCount),
		onDash:  OnDash,
	}
}

func (ef *EFront) Init(chKb chan<- t.KbEvent) {

}

func (ef *EFront) AddWdg(wdg *w.Wdg, delta int) {
	ef.widgets[wdg.Id+delta] = wdg
}

func (ef *EFront) onFocus(wdgId int) {
	focus := int(ef.focus.Load())
	oldWdg := ef.widgets[focus]
	newWdg := ef.widgets[wdgId]
	if newWdg == nil {
		return
	}
	ef.focus.Swap(uint32(wdgId))

	ff := t.Focusable(newWdg)
	ff.OnActivate(true)
	ef.onDraw(newWdg)

	if oldWdg != nil {
		ff = t.Focusable(oldWdg)
		ff.OnActivate(false)
		ef.onDraw(oldWdg)
	}
}

func (ef *EFront) onClickEnter() {

}
func (ef *EFront) onPgUp(wdgId int) {
	wdg := ef.widgets[wdgId]
	if wdg == nil {
		return
	}
	if wdg.Type == constTypeTab || wdg.Type == constTypeList {
		wdg.Value -= 100
	}
}

func (ef *EFront) onPgDown(wdgId int) {
	wdg := ef.widgets[wdgId]
	if wdg == nil {
		return
	}
	if wdg.Type == constTypeTab || wdg.Type == constTypeList {
		wdg.Value += 100
	}
}

func (ef *EFront) onDraw(wdg *w.Wdg) {
	fmt.Println(wdg.String())
}

func (ef *EFront) CallbackHealthy() func(serviceId string, state uint8) {
	obj := ef
	return func(serviceId string, state uint8) {
		if state > 0 {
			state = 1
		}
		dash := ef.onDash(obj)
		switch serviceId {
		default:
			dash.Value = int(state)
			if state == 0 {
				dash.Active = false
			} else {
				dash.Active = true
			}
		}
	}
}
