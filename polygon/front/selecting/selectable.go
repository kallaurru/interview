package selecting

type OnSelect struct {
	sRow      int
	sRowTheme uint64
	count     int // количество строк всего
	countPP   int // количество строк на страницу
}

func New(count, countPP int, theme uint64) *OnSelect {
	if countPP < 1 {
		countPP = 1
	}
	return &OnSelect{
		sRow:      -1, // дефолтное состояние. Ничего не выделено
		sRowTheme: theme,
		count:     count,
		countPP:   countPP,
	}
}

func (os *OnSelect) Up() {
	if os.sRow-1 <= 0 {
		os.sRow = 0
		return
	}
	os.sRow -= 1
}

func (os *OnSelect) Down() {
	if os.sRow+1 == os.count {
		return // мы в самом конце
	}
	os.sRow += 1
}

func (os *OnSelect) PgUp() {
	page := os.sRow/os.countPP + 1
	if page == 1 {
		return // уже в начале
	}
	os.sRow = (page-1)*os.count - os.countPP
}

func (os *OnSelect) PgDown() {
	page := os.sRow/os.countPP + 1
	if page*os.count >= os.count {
		return // уже в конце
	}
	os.sRow = (page+1)*os.count - os.countPP
}

// Selected - возвращает номер строки которая сейчас выделена. Нумерация с 0
func (os *OnSelect) Selected() int {
	return 0
}
