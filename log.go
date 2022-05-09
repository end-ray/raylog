package raylog

func CheckMsg(level uint8, code string, message string) {
	switch Level(level) {
	case InfoLevel:
		Info(code).Msg(message)
	}
}

func CheckErr(level uint8, code string, err error) {
	switch Level(level) {
	case WarnLevel:
		Warn(code).Err(err)
	case PanicLevel:
		Panic(code).Err(err)
	}
}

func Output(level uint8, code string, err error) {

	switch Level(level) {
	case InfoLevel:
		Info(code).Err(err)
	case WarnLevel:
		Warn(code).Err(err)
	case PanicLevel:
		Panic(code).Err(err)
	}
}

func Info(code string) *Event {
	return newEvent(code, InfoLevel)
}

// Неожиданные параметры вызова, странный формат запроса, использование дефолтных значений в замен не корректных.
// Вообще все, что может свидетельствовать о не штатном использовании.
func Warn(code string) *Event {
	return newEvent(code, WarnLevel)
}

// Исключения не совместимые с работой приложения (defer)
func Panic(code string) *Event {
	return newEvent(code, PanicLevel)
}
