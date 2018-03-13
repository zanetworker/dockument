package log

import (
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/zanetworker/go-kubesanity/pkg/kubesanityutils"
)

//Error provides error logging with a reference to the module name
func Error(message string) {
	_, file, line, _ := runtime.Caller(1)
	log.Errorf("[%s] %s : line: %d", getFileNameCapitalized(file), kubesanityutils.ColorString("red", message), line)
}

//ErrorS provides error logging with a reference to the module name
func ErrorS(message string, err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Errorf("[%s] %s : line: %d:%s", getFileNameCapitalized(file), kubesanityutils.ColorString("red", message), line, err.Error())
}

//Fatal provides fatal error logging with a reference to the module name
func Fatal(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Fatalf("[%s] : line: %d: %s", getFileNameCapitalized(file), line, kubesanityutils.ColorString("red", err.Error()))
}

//FatalS provides a fatal error with a message string
func FatalS(message string) {
	_, file, line, _ := runtime.Caller(1)
	log.Fatalf("[%s] : line: %d: %s", getFileNameCapitalized(file), line, kubesanityutils.ColorString("red", message))
}

//Debug provides debug level logging with a reference to the module name
func Debug(message string, err error) {
	_, file, _, _ := runtime.Caller(11)
	log.Debugf("[%s] %s : %s", getFileNameCapitalized(file), kubesanityutils.ColorString("green", message), err.Error())
}

//Info provides Info level logging with a reference to the module name
func Info(message string) {
	_, file, _, _ := runtime.Caller(1)
	log.Infof("[%s] %s", getFileNameCapitalized(file), kubesanityutils.ColorString("green", message))
}
