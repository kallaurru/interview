package carrier

import (
	"bytes"
	"encoding/binary"
)

const (
	constDefDelim    byte = ';'
	constDefChanSize int  = 8
	constDefBodySize int  = 128

	constStartCountFields = 0
	constStartKeyFieldIdx = 4
	constStartCountLines  = 8
	constStartFlgHeaders  = 12
)

/*
header - разметка байт:
0 - 3  - количество полей в строке сообщения
4 - 7  - индекс поля в котором находится ключ
8 - 11 - количество строк данных
12 - 15 - если там 1, значит первая строка данных будут заголовки

Строки в теле разделяем переносом строки
*/

type Carrier struct {
	header []byte // храним заголовок
	body   []byte
	delim  byte
}

func NewCarr(keyIdx int, headers []string, d ...byte) Carrier {
	body := new(bytes.Buffer)
	header := make([]byte, 16)
	delim := constDefDelim
	if len(d) > 0 {
		delim = d[0]
	}
	carri := Carrier{
		delim: delim,
		body:  make([]byte, 0, constDefBodySize),
	}
	if headers != nil {
		binary.BigEndian.PutUint32(header[constStartFlgHeaders:constStartFlgHeaders+4], uint32(1))
		binary.BigEndian.PutUint32(header[constStartCountFields:constStartCountFields+4], uint32(len(headers)))
		res := packLine(headers, delim)
		body.Write(res)
		carri.body = append(carri.body, body.Bytes()...)
	}
	if keyIdx < 0 {
		keyIdx = 0 // по умолчанию 0
	}
	binary.BigEndian.PutUint32(header[constStartKeyFieldIdx:constStartKeyFieldIdx+4], uint32(keyIdx))
	carri.header = header

	return carri
}

func (c *Carrier) AddLine(val []string) {
	data := packLine(val, c.delim)
	res := append(c.body, data...)
	c.body = res

	exists := binary.BigEndian.Uint32(c.header[constStartCountLines : constStartCountLines+4])
	exists++
	binary.BigEndian.PutUint32(c.header[constStartCountLines:constStartCountLines+4], exists)
}

func (c Carrier) Lines() int {
	lines := int(binary.BigEndian.Uint32(c.header[constStartCountLines : constStartCountLines+4]))
	if c.HasHeaders() {
		lines++
	}
	return lines
}

func (c Carrier) Fields() int {
	return int(binary.BigEndian.Uint32(c.header[constStartCountFields : constStartCountFields+4]))
}

func (c Carrier) HasHeaders() bool {
	val := binary.BigEndian.Uint32(c.header[constStartFlgHeaders : constStartFlgHeaders+4])

	return val > 0
}

func (c Carrier) KeyF() int {
	val := binary.BigEndian.Uint32(c.header[constStartKeyFieldIdx : constStartKeyFieldIdx+4])
	return int(val)
}

func (c Carrier) Len() int {
	return len(c.body) + len(c.header)
}

func Unmarshall(c Carrier) <-chan [][]byte {
	out := make(chan [][]byte, constDefChanSize)
	go func(out chan<- [][]byte, body []byte, count int) {
		defer close(out)
		start := 0
		for start < len(body) {
			line, pos := parseLine(body, start)
			if pos == -1 {
				return
			}
			fields := parseFields(line, count, c.delim)
			out <- fields
			start += pos
		}
	}(out, c.body, c.Lines())

	return out
}

func (c Carrier) Parse() map[int][][]byte {
	out := make(map[int][][]byte, c.Lines())
	start := 0
	lineIdx := 0
	for start < len(c.body) {
		line, pos := parseLine(c.body, start)
		fields := parseFields(line, c.Fields(), c.delim)
		start += pos
		out[lineIdx] = fields
		lineIdx++
	}
	return out
}

func packLine(fields []string, delim byte) []byte {
	buf := new(bytes.Buffer)

	for i := 0; i < len(fields)-1; i++ {
		buf.WriteString(fields[i])
		buf.Write([]byte{delim})
	}
	// пишем последнюю строку
	buf.WriteString(fields[len(fields)-1])
	buf.Write(defaultLineDelim())

	return buf.Bytes()
}

func parseLine(part []byte, start int) ([]byte, int) {
	// part[start:] будет возвращать отсчет с 0. Pos не будет учитывать сквозную нумерацию
	pos := bytes.Index(part[start:], defaultLineDelim())

	if pos == -1 {
		pos = len(part[start:])
	}
	out := make([]byte, pos)
	copy(out, part[start:pos+start]) // сквозная нумерация

	return out, pos + 2
}

func parseFields(data []byte, count int, delim byte) [][]byte {
	out := make([][]byte, 0, count)
	start := 0
	for start < len(data) {
		// data[start:] всегда будет начинаться с 0 и pos может быть меньше чем сквозная нумерация
		pos := bytes.IndexByte(data[start:], delim)
		if pos == -1 {
			// значит это последнее поле
			pos = len(data[start:])
		}
		item := make([]byte, pos)
		copy(item, data[start:pos+start]) // данные вырезаем из общего слайса со сквозной нумерацией
		out = append(out, item)
		start += pos + 1
	}
	return out
}

func defaultLineDelim() []byte {
	return []byte{0x0d, 0x0a}
}
