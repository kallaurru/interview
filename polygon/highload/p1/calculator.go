package p1

// CalculatorBofE - калькулятор back-of-envelope
// input dau - пользователей в сутки
//
//	r - количество запросов одного пользователя
//	k - пиковый коэффициент
//	b - сколько байт записи порождает один запрос
func CalculatorBofE(dau, r, k, b int) (avgRps int64, peakRps int64, storBytesDay int64) {
	const secPerDay int = 86400
	avg := (dau*r + (secPerDay - 1)) / secPerDay
	avgRps = int64(avg)
	peakRps = int64(avg * k)
	storBytesDay = int64(dau * r * b)

	return avgRps, peakRps, storBytesDay
}
