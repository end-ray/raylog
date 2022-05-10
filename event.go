package raylog

import (
	"fmt"
	"os"
	"time"
)

type Configuration struct {
	Home  string
	Level Level
}

var config Configuration

type Level uint8 // uint8: представляет целое число от 0 до 255 и занимает 1 байт

type Event struct {
	t     time.Time
	code  string
	level Level
	msg   string
}

const (
	PanicLevel Level = iota // 0   iota: auto increment
	FatalLevel              // 1
	ErrorLevel              // 2
	WarnLevel               // 3
	InfoLevel               // 4
	DebugLevel              // 5
	TraceLevel              // 6
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "PANIC"
	}
	return ""
}

func newEvent(code string, level Level) *Event {
	var event Event

	event.t = time.Now()
	event.code = code
	event.level = level

	return &event
}

func (e *Event) Err(err error) {
	e.msg = err.Error()

	//fmt.Printf("%T\n", e.t)
	//fmt.Println(e.level.String(), e.t.Format("02-01-2006 15:04:05\n"), e.code, e.msg)

	log := fmt.Sprintf("%s %s %s: %s\n", e.level.String(), e.t.Format("02-01-2006 15:04:05\n"), e.code, e.msg)

	if config.Level <= e.level {
		writeLog(log)
	}

	if e.level == PanicLevel {
		panic(log)
	}
	fmt.Println(log)

}

func (e *Event) Msg(message string) {
	e.msg = message

	//fmt.Printf("%T\n", e.t)
	//fmt.Println(e.level.String(), e.t.Format("02-01-2006 15:04:05\n"), e.code, e.msg)

	log := fmt.Sprintf("%s %s %s: %s\n", e.level.String(), e.t.Format("02-01-2006 15:04:05\n"), e.code, e.msg)

	if config.Level >= e.level {
		writeLog(log)
	}

	if e.level == PanicLevel {
		panic(log)
	}
	fmt.Println(log)

}

func SetConfig(path string) {
	config.Home = path //назначаем переменной значение
}

func SetLevel(level Level) {
	config.Level = level
}

func writeLog(message string) {

	f, err := os.OpenFile(config.Home, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(message); err != nil {
		panic(err)
	}
}
