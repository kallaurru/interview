package front

import (
	"context"
	"fmt"
	glg "github.com/bahlo/generic-list-go"
	t "github.com/kallaurru/interview/polygon/front/talk"
	w "github.com/kallaurru/interview/polygon/front/widgets"
	"sync/atomic"
)

/*
	что бы не потерять ссылку на лист с generics
	glg "github.com/bahlo/generic-list-go"
*/

type OnActionF func(ef *EFront) *w.Wdg
type OnActionKbd func(ef *EFront, kb uint32)

type EFront struct {
	grid      *w.Wdg
	widgets   []*w.Wdg // 0 - nil, 1 - 9 виджеты или nil если их меньше, > 9 модалки
	queue     *glg.List[int]
	onDash    OnActionF
	onKbClick OnActionKbd // Вариативность действий кнопок
	focus     atomic.Uint32
	editable  atomic.Bool
}

const (
	constKeyEnter uint32 = 0x01
	constKeyPgUp  uint32 = 0x02
	constKeyPgDw  uint32 = 0x04
	constKeyUp    uint32 = 0x10
	constKeyDw    uint32 = 0x20
	constKeyTab   uint32 = 0x0100
	constKeyEsc   uint32 = 0x010000
	constKeyExit  uint32 = 0x100000
	constKeyMenuB uint32 = 0x100001
	constKeyMenuA uint32 = 0x100002

	constFocus int = 0
	constDelta int = 10
)

func New(wdg []*w.Wdg) *EFront {
	f := &EFront{
		queue:     glg.New[int](),
		widgets:   wdg,
		onDash:    OnDash,
		onKbClick: PlugActionsKbd,
	}
	f.grid = wdg[0]

	return f
}

func (ef *EFront) Init(ctx context.Context, chKb <-chan t.KbEvent) {
begin:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled")
			break begin
		case e := <-chKb:
			if e.Key == constKeyExit {
				fmt.Println("Std exit")
				break begin
			}
			switch e.Key {
			case constKeyTab:
				ef.onFocus()
			case constKeyEsc:
				ef.onEscape()
			case constKeyEnter:
				ef.onClickEnter()
			default:
				ef.onKbClick(ef, e.Key)
			}
		}
	}
}

func (ef *EFront) onFocus() {
	focus := int(ef.focus.Load())
	if focus > constDelta {
		return
	}
	nextId := focus + 1
	if nextId > constDelta-1 {
		nextId = 0
	}
	ef.swapFocus(nextId)
}

func (ef *EFront) swapFocus(wdgId int) {
	focus := int(ef.focus.Load())
	if focus > constDelta {
		return
	}
	if wdgId >= len(ef.widgets) {
		wdgId = 0
	}
	oldWdg := ef.widgets[focus]
	newWdg := ef.widgets[wdgId]
	if newWdg == nil {
		return
	}
	ef.focus.Swap(uint32(wdgId))

	ff := t.Focusable(newWdg)
	ff.OnActivate(true)
	//ef.onDraw(newWdg) // todo возможно отрисовка будет не здесь

	if oldWdg != nil {
		ff = t.Focusable(oldWdg)
		ff.OnActivate(false)
		// ef.onDraw(oldWdg) // todo возможно отрисовка будет не здесь
	}
}

func (ef *EFront) onClickEnter() {
	wdg := ef.focusedWdg()
	if wdg == nil {
		return
	}
	if wdg.Type == w.ConstTypeTab {
		fmt.Println("Table value:", wdg.Value)
	}

	ef.onDraw(wdg)
}

func (ef *EFront) onPgUp() {
}

func (ef *EFront) onPgDown() {
}

func (ef *EFront) onDraw(wdg *w.Wdg) {
	fmt.Println(wdg.String())
}

func (ef *EFront) onEscape() {
	wdg := ef.focusedWdg()
	if wdg == nil {
		return
	}
	if wdg.WindowType == w.ConstWindowTypeModal {
		ef.swapFocus(0)
	}
}

func (ef *EFront) focusedWdg() *w.Wdg {
	focus := ef.focus.Load()
	if focus == 0 {
		return ef.grid
	}

	if int(focus) > len(ef.widgets) {
		return nil
	}

	return ef.widgets[int(focus)]
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
