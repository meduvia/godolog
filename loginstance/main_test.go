package loginstance_test

import (
	"testing"

	"github.com/meduvia/godolog/log"
	"github.com/meduvia/godolog/loginstance"
	"github.com/meduvia/godolog/writer"
)

type ConsoleWriterBack struct {
	logfunc func (string)
}

func (console *ConsoleWriterBack) Debug(message string) {
	console.logfunc(message)
}
func (console *ConsoleWriterBack) Info(message string) {
	console.logfunc(message)
}
func (console *ConsoleWriterBack) Warn(message string) {
	console.logfunc(message)
}
func (console *ConsoleWriterBack) Error(message string) {
	console.logfunc(message)
}

func TestLogInstance(t *testing.T) {
	var last_log string = "null"
	 logfunc := func (message string) {
		last_log = message
	}
	var logback writer.ConsoleWriterInterface = &ConsoleWriterBack{
		logfunc: logfunc,
	}
	logwriter := writer.NewConsoleWriter(&logback)
	logmgr := loginstance.NewLogInstance(&logwriter, log.FlagVerbose)

	product := "App"
	service := "Service"
	location := "EU01-D01-R01-N01"
	level := log.FlagDebug
	code := "OP-R403"
	message := "Rest API Error 403"
	logmgr.Log(product, service, location, level, code, message)
	if (last_log == "null") {
		t.Fail() // Log was not sended
	}
}


func TestLogInstanceLevel(t *testing.T) {
	var last_log string = "null"
	logfunc := func (message string) {
		last_log = message
 }
 var logback writer.ConsoleWriterInterface = &ConsoleWriterBack{
	 logfunc: logfunc,
 }
 logwriter := writer.NewConsoleWriter(&logback)
 logmgr := loginstance.NewLogInstance(&logwriter, log.FlagDebug)

 product := "App"
 service := "Service"
 location := "EU01-D01-R01-N01"
 code := "OP-R403"
 message := "Rest API Error 403"
 level_not_pass := log.FlagVerbose
 level_pass := log.FlagError
 logmgr.Log(product, service, location, level_not_pass, code, message)
 if (last_log != "null") {
	 t.Fail() // log was sended but should be filtered
 }
 logmgr.Log(product, service, location, level_pass, code, message)
 if (last_log == "null") {
	 t.Fail() // log should be displayed
 }
}