package MicroLogger

import (
	"github.com/fatih/color"
	l "log"
)

type LoggerLevel int

const (
	INFO LoggerLevel = iota
	OK
	WARN
	FAILED
	FATAL
	EXTENSION
	DEBUG
)

var FatalColor = color.New(color.BgHiRed, color.FgHiWhite)

func Log(level LoggerLevel, log string) {
	switch level {
	case INFO:
		l.Printf("[ INFO ] %s\n", log)
		break
	case OK:
		l.Printf("[  %s  ] %s\n", color.HiGreenString("OK"), log)
		break
	case WARN:
		l.Printf("[ %s ] %s\n", color.HiYellowString("WARN"), log)
		break
	case FAILED:
		l.Printf("[%s] %s\n", color.HiRedString("FAILED"), log)
		break
	case FATAL:
		l.Printf("[ %s ] %s\n", FatalColor.Sprint("FATL"), log)
		break
	case EXTENSION:
		l.Println("          ->", log)
		break
	case DEBUG:
		l.Printf("[ %s ] %s\n", color.BlueString("DBUG"), log)
		break
	default:
		l.Printf("[%s] Parameter level is not of type LoggerLevel!\n", color.HiRedString("FAILED"))
		break
	}
}

func main() {

}
