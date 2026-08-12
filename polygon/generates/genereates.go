package generates

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"
	"time"
)

const (
	constReqIDSize = 16
)

func GenerateReqID(uuid uint16) ([]byte, bool) {
	// reqID = PPPPHHRRMMMMMMMM Итого 16 байт
	// PPPP - prefix req- ; HH - тек час всегда 2 цифры; MMMMMMMM - кол-во милисекунд с начала дня
	// RR - уникальный константный идентификатор
	// пакуем сначала символы потом числа что бы было удобнее разбирать
	digits := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	elapsedMsOfStartDay := func() uint64 {
		now := time.Now()
		startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		elapsedMs := now.Sub(startOfDay).Milliseconds()

		return uint64(elapsedMs)
	}
	buf := new(bytes.Buffer)
	buf.Write([]byte{'r', 'e', 'q', '-'})

	allMili := 60 * 60 * 24 * 1000
	partSize := allMili / 24
	elapsed := elapsedMsOfStartDay()
	part := elapsed/uint64(partSize) + 1 // в 00:30 будет показывать что это 1 первый час
	// HH = 00  будет указывать на проблему
	if part < 0 && part > 24 {
		part = 0
	}

	if part < 10 {
		buf.Write([]byte{digits[0], digits[part]})
	} else if part < 20 {
		buf.Write([]byte{digits[1], digits[part-10]})
	} else {
		buf.Write([]byte{digits[2], digits[part-20]})
	}
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp, uuid)
	buf.Write(tmp)

	tmp = make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, elapsed)
	buf.Write(tmp)

	return buf.Bytes(), len(buf.Bytes()) == constReqIDSize
}

func ParseNiceReqID(reqID []byte) (string, bool) {
	if len(reqID) < constReqIDSize {
		return "", false
	}
	b := strings.Builder{}
	pUuid := reqID[6:8]
	pElaps := reqID[8:16]

	for _, val := range reqID[:4] {
		b.WriteByte(val)
	}
	uint16Val := binary.BigEndian.Uint16(pUuid)
	uintStr := strconv.Itoa(int(uint16Val))
	b.WriteString(uintStr)
	b.WriteByte(45)
	// part of day
	for _, val := range reqID[4:6] {
		b.WriteByte(val)
	}
	uint64Val := binary.BigEndian.Uint64(pElaps)
	uintStr = strconv.FormatUint(uint64Val, 10)

	b.WriteByte(45)
	b.WriteString(uintStr)

	return b.String(), true
}
