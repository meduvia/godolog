package writer_test

import (
	"testing"

	"github.com/meduvia/godolog/log"
	"github.com/meduvia/godolog/writer"
)

var debug string
var info string
var warn string
var errorval string

type LogBackendTest struct {
}

func (LogBackend *LogBackendTest) Debug(text string) {
	debug = text
}
func (LogBackend *LogBackendTest) Info(text string) {
	info = text
}
func (LogBackend *LogBackendTest) Warn(text string) {
	warn = text
}
func (LogBackend *LogBackendTest) Error(text string) {
	errorval = text
}
func TestConsoleWriter(t *testing.T) {
	var logbackend writer.ConsoleWriterInterface = &LogBackendTest{}
	consolewriter := writer.NewConsoleWriter(&logbackend)
	product := "App"
	service := "Service"
	location := "EU01-D01-R01-N01"
	code := "OP-R403"
	message := "Rest API Error 403"

	t.Run("Test verbose write", func(t *testing.T) {
		logobj := log.CreateLog(product, service, location, log.FlagVerbose, code, message)
		consolewriter.Write(logobj)
		val, err := logobj.ToJSONString()
		if debug != val || err != nil {
			t.Fail()
		}
	})

	t.Run("Test debug write", func(t *testing.T) {
		logobj := log.CreateLog(product, service, location, log.FlagDebug, code, message)
		consolewriter.Write(logobj)
		val, err := logobj.ToJSONString()
		if debug != val || err != nil {
			t.Fail()
		}
	})
	t.Run("Test info write", func(t *testing.T) {
		logobj := log.CreateLog(product, service, location, log.FlagInfo, code, message)
		consolewriter.Write(logobj)
		val, err := logobj.ToJSONString()
		if info != val || err != nil {
			t.Fail()
		}
	})
	t.Run("Test warn write", func(t *testing.T) {
		logobj := log.CreateLog(product, service, location, log.FlagWarn, code, message)
		consolewriter.Write(logobj)
		val, err := logobj.ToJSONString()
		if warn != val || err != nil {
			t.Fail()
		}
	})
	t.Run("Test error write", func(t *testing.T) {
		logobj := log.CreateLog(product, service, location, log.FlagError, code, message)
		consolewriter.Write(logobj)
		val, err := logobj.ToJSONString()
		if errorval != val || err != nil {
			t.Fail()
		}
	})
	t.Run("Test fatal write", func(t *testing.T) {
		logobj := log.CreateLog(product, service, location, log.FlagFatal, code, message)
		consolewriter.Write(logobj)
		val, err := logobj.ToJSONString()
		if errorval != val || err != nil {
			t.Fail()
		}
	})
	t.Run("Test default write", func(t *testing.T) {
		logobj := log.CreateLog(product, service, location, "default", code, message)
		consolewriter.Write(logobj)
		val, err := logobj.ToJSONString()
		if debug != val || err != nil {
			t.Fail()
		}
	})

	// t.Run("Test tojsonstring error handle", func(t *testing.T) {

	// 	consolewriter.Write(logobj)
	// 	val, err := logobj.ToJSONString()
	// 	if errorval != writer.ErrorJSONString || err != nil {
	// 		t.Fail()
	// 	}
	// })

	
}

