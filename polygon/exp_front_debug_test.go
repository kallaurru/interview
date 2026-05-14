package polygon

import (
	"context"
	f "github.com/kallaurru/interview/polygon/front"
	tt "github.com/kallaurru/interview/polygon/front/talk"
	w "github.com/kallaurru/interview/polygon/front/widgets"
	"testing"
)

const (
	constKeyEnter uint32 = 0x01
	constKeyPgUp  uint32 = 0x02
	constKeyPgDw  uint32 = 0x04
	constKeyUp    uint32 = 0x10
	constKeyDw    uint32 = 0x20
	constKeyTab   uint32 = 0x0100
	constKeyEsc   uint32 = 0x010000
	constKeyExit  uint32 = 0x100000
)

func TestExperimentalFront(t *testing.T) {
	// создаем виджеты
	front := f.New(fixturesWdg())
	ctx, cancel := context.WithCancel(context.Background())
	chE := make(chan tt.KbEvent, 1)
	defer cancel()
	// имитируем нажатия клавиш
	go front.Init(ctx, chE)
}

func fixturesWdg() []*w.Wdg {
	out := make([]*w.Wdg, 12)

	grid := w.New(0, w.ConstTypeGrid)
	out[0] = grid
	// col 0 row 0 dashboard
	dash := w.New(1, w.ConstTypeTab)
	out[1] = dash
	// col 0 row 1 objects tab
	objects := w.New(2, w.ConstTypeTab)
	out[2] = objects
	// col 1 row 0 rules
	rules := w.New(3, w.ConstTypeTab)
	out[3] = rules
	// modal menu 1
	menuBasicAction := w.New(10, w.ConstTypeTab)
	menuBasicAction.WindowType = w.ConstWindowTypeModal
	out[10] = menuBasicAction
	// modal menu 2
	advancedActions := w.New(11, w.ConstTypeTab)
	advancedActions.WindowType = w.ConstWindowTypeModal
	out[11] = advancedActions

	// modal confirm
	confirm := w.New(12, w.ConstTypeList)
	confirm.WindowType = w.ConstWindowTypeModal
	out[12] = confirm

	return out
}

func fixtureKeys() []uint32 {
	return []uint32{
		constKeyTab,
		constKeyTab,
		constKeyTab,
		constKeyEsc,
		constKeyExit,
	}
}
