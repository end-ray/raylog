package raylog

import (
	"fmt"
	"os"
	"time"
)

type Configuration struct {
	Home string
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
	TraceLevel Level = iota // 0   iota: auto increment
	DebugLevel              // 1
	InfoLevel               // 2
	WarnLevel               // 3
	ErrorLevel              // 4
	FatalLevel              // 5
	PanicLevel              // 6
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
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

//func (e *Event) Msg(msg string) *Event {
//	return nil
//}

func (e *Event) Err(err error) {
	e.msg = err.Error()

	//fmt.Printf("%T\n", e.t)
	//fmt.Println(e.level.String(), e.t.Format("02-01-2006 15:04:05\n"), e.code, e.msg)

	log := fmt.Sprintf("%s %s %s: %s\n", e.level.String(), e.t.Format("02-01-2006 15:04:05\n"), e.code, e.msg)

	writeLog(log)

	if e.level == PanicLevel {
		panic(log)
	}
	fmt.Println(log)

}

func GetConfig(path string) {
	config.Home = path //назначаем переменной значение
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
