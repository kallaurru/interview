package widgets

import (
	"fmt"
	"github.com/kallaurru/interview/polygon/front/selecting"
	"github.com/kallaurru/interview/polygon/front/talk"
)

const (
	ConstTypeTab   uint8 = 0x01
	ConstTypeList  uint8 = 0x02
	ConstTypeGauge uint8 = 0x04
	ConstTypeModal uint8 = 0x08
)

type Wdg struct {
	Active   bool
	Type     uint8
	Id       int
	Value    int
	onSelect talk.Selectable
}

func New(id int, t uint8) *Wdg {
	return &Wdg{
		Active:   false,
		Type:     t,
		Id:       id,
		Value:    0,
		onSelect: nil,
	}
}

func (wdg *Wdg) CanSelect() {
	rows := 115
	rowsPerPage := 10
	var theme uint64 = 8737732

	wdg.onSelect = selecting.New(rows, rowsPerPage, theme)
}

func (wdg *Wdg) OnActivate(active bool) {
	wdg.Active = active
}
func (wdg *Wdg) String() string {
	return fmt.Sprintf(
		"Active: %v, Type: %d, Id: %d, Value: %d\n",
		wdg.Active, wdg.Type, wdg.Id, wdg.Value)
}
